package walletsvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"star_net/app/admin/internal/model"
	"star_net/consts"
	"star_net/core/wallet"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xtrans"
)

type Withdraw struct {
	OrderNo      int64
	Status       int
	StatusRemark string
	SysRemark    string
}

func (w *Withdraw) Update(ctx context.Context) error {
	userInfo := ctx.Value(consts.UserInfo).(model.UserInfo)
	order := entity.Withdraw{}
	dao.Withdraw.Ctx(ctx).Scan(&order, dao.Withdraw.Columns().OrderNo, w.OrderNo)
	if order.OrderNo == 0 {
		return fmt.Errorf("没有此订单")
	}
	if w.Status == order.Status {
		return fmt.Errorf("不需要修改")
	}
	if order.Status != 0 {
		return fmt.Errorf("不能修改")
	}
	order.StatusRemark = w.StatusRemark
	order.SysRemark = w.SysRemark
	order.AdminOperate = fmt.Sprintf("%s_%v_%s", userInfo.Account, userInfo.Uid, userInfo.ClientIp)
	order.FinishAt = gtime.Now()
	update := wallet.BalanceUpdate{}
	update.Amount = order.Amount
	update.OrderNoRelation = w.OrderNo
	update.Uid = order.Uid
	update.Note = fmt.Sprintf("%s_%s", order.Protocol, order.Address)
	if w.Status == consts.WithdrawStatusSuccess {
		user := entity.User{}
		dao.User.Ctx(ctx).Scan(&user, order.Uid)
		order.Status = consts.WithdrawStatusSuccess
		order.StatusRemark = xtrans.New(user.Lang).T(ctx, "提现成功")
		_, err := dao.Withdraw.Ctx(ctx).Update(&order, dao.Withdraw.Columns().OrderNo, w.OrderNo)
		return err
	} else if w.Status == consts.DepositStatusFail {
		order.Status = consts.WithdrawStatusFail
		update.BalanceCode = wallet.WithdrawFail
		update.Note = w.StatusRemark
		update.Title = "Fail " + w.StatusRemark
		return update.Update(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err := tx.Model(dao.Withdraw.Table()).Ctx(ctx).Update(&order, dao.Withdraw.Columns().OrderNo, w.OrderNo)
			return err
		})
	} else {
		_, err := dao.Withdraw.Ctx(ctx).Update(&order)
		return err
	}

	return nil

}
