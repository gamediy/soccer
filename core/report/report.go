package report

import (
	"context"
	"github.com/avast/retry-go/v4"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xuuid"
)

type Report struct {
	User              entity.User
	Amount            float64
	BalanceCodeEntity entity.BalanceCode
}

func (this *Report) PutReport() error {
	err := retry.Do(func() error {
		date := gtime.Now().Format("Y-m-d")
		ctx := context.TODO()
		entity := entity.ReportWalletDay{}
		dao.ReportWalletDay.Ctx(ctx).Where("date=? and uid=? and balance_code=?", date, this.User.Id, this.BalanceCodeEntity.Code).Scan(&entity)
		if entity.Id == 0 {
			entity.Amount = this.Amount
			entity.Id = xuuid.GetsnowflakeUUID().Int64()
			entity.BalanceCode = this.BalanceCodeEntity.Code
			entity.ParentPath = this.User.ParentPath
			entity.Pid = this.User.Pid
			entity.Account = this.User.Account
			entity.Uid = this.User.Id
			entity.Count = 1
			entity.Title = this.BalanceCodeEntity.Title
			entity.Date = gtime.NewFromStr(date)
			_, err := dao.ReportWalletDay.Ctx(ctx).Insert(&entity)
			if err != nil {
				g.Log().Error(ctx, err)
				return err
			}
		} else {
			entity.Amount += this.Amount
			entity.Count += 1
			_, err := dao.ReportWalletDay.Ctx(ctx).Update(&entity, dao.ReportWalletDay.Columns().Id, entity.Id)
			if err != nil {
				g.Log().Error(ctx, err)
				return err
			}
		}
		return nil
	}, retry.Attempts(3))
	return err

}
