package walletsvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/app/admin/internal/model"
	"star_net/consts"
	"star_net/core/wallet"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xtrans"
)

type Deposit struct {
	OrderNo      int64
	Status       int
	StatusRemark string
	SysRemark    string
}

func (s *Deposit) Update(ctx context.Context) error {
	userInfo := ctx.Value("userInfo").(model.UserInfo)
	order := entity.Deposit{}
	dao.Deposit.Ctx(ctx).Scan(&order, dao.Deposit.Columns().OrderNo, s.OrderNo)
	if order.OrderNo == 0 {
		return fmt.Errorf("没有此订单")
	}
	if s.Status == order.Status {
		return fmt.Errorf("不需要修改")
	}
	if order.Status == consts.DepositStatusSuccess {
		return fmt.Errorf("已经成功，不能修改")
	}

	order.StatusRemark = s.StatusRemark
	order.SysRemark = s.SysRemark
	order.AdminOperate = fmt.Sprintf("%s_%v_%s", userInfo.Account, userInfo.AdminId, userInfo.ClientIp)
	order.Status = s.Status
	order.FinishAt = gtime.Now()
	if s.Status == consts.DepositStatusSuccess {
		user := entity.User{}
		dao.User.Ctx(ctx).Scan(&user, order.Uid)
		update := wallet.BalanceUpdate{}
		update.Amount = gconv.Float64(order.Amount)
		update.Title = xtrans.New(user.Lang).T(ctx, "充值成功")
		order.StatusRemark = xtrans.New(user.Lang).T(ctx, "充值成功")
		update.Uid = order.Uid
		update.OrderNoRelation = s.OrderNo
		update.Note = fmt.Sprintf("%s_%s", order.Protocol, order.Address)
		update.BalanceCode = wallet.Deposit
		if order.Net == "BANK" {
			update.BalanceCode = wallet.DepositBank
		} else if order.Net == "GCASH" {
			update.BalanceCode = wallet.DepositGcash
		}
		return update.Update(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err := tx.Model(dao.Deposit.Table()).Ctx(ctx).Update(&order, dao.Deposit.Columns().OrderNo, order.OrderNo)
			return err
		})

	} else {
		_, err := dao.Deposit.Ctx(ctx).Update(&order, dao.Deposit.Columns().OrderNo, order.OrderNo)
		return err
	}

}
