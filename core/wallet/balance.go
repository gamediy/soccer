package wallet

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"math"
	"star_net/core/report"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"time"
)

const (
	precision = 1000 //余额精度
)

const (
	Bet          = -700
	Withdraw     = -100
	WithdrawFail = 100
	Deposit      = 500
	DepositBank  = 501
	DepositGcash = 502
	Profit       = 201
	Gift         = 400
	Unfreeze     = 800
	Freeze       = -800
)

type BalanceUpdate struct {
	Uid             int64
	Amount          float64
	OrderNoRelation int64
	BalanceCode     int

	Note string
}

func (this *BalanceUpdate) Update(ctx context.Context, fc func(ctx context.Context, tx gdb.TX) error) error {

	timeout, cancelFunc := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancelFunc()
	var err error

	for {
		err = this.updateExec(ctx, fc)
		if err != nil {
			fmt.Println(err)
		}
		if err == nil {
			return nil
		}
		select {
		case <-timeout.Done():
			return err
		case <-time.After(1 * time.Second):

		}
	}
}
func (this *BalanceUpdate) updateExec(ctx context.Context, fc func(ctx context.Context, tx gdb.TX) error) error {
	this.Amount = math.Abs(float64(this.Amount))
	realAmouont := int64(this.Amount * precision)
	if this.Amount == 0 || realAmouont == 0 {
		return fmt.Errorf("error amount")
	}
	user := entity.User{}
	wallet := entity.Wallet{}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if err := tx.Model(dao.User.Table()).Where("id", this.Uid).LockUpdate().Scan(&user); err != nil {
			return err
		}
		if err := tx.Model(dao.Wallet.Table()).Where("uid", this.Uid).Scan(&wallet); err != nil {
			return err
		}
		if user.Id == 0 || wallet.Uid == 0 {
			return fmt.Errorf("NO such user")
		}
		if this.BalanceCode == Unfreeze {
			if realAmouont > wallet.Freeze {
				return gerror.NewCode(gcode.New(-1000, "余额不足", ""))
			}
		}
		bc := entity.BalanceCode{}
		tx.Model(dao.BalanceCode.Table()).Where("code", this.BalanceCode).Scan(&bc)
		node, err := snowflake.NewNode(1)
		if err != nil {
			return err
		}
		orderNo := node.Generate()
		walletLog := &entity.WalletLog{}
		walletLog.OrderNo = orderNo.Int64()
		walletLog.OrderNoRelation = this.OrderNoRelation
		walletLog.Uid = user.Id
		walletLog.Pid = user.Pid
		walletLog.Account = user.Account
		walletLog.ParentPath = user.ParentPath
		walletLog.BalanceBefore = float64(wallet.Balance / 1000)
		walletLog.Amount = this.Amount
		walletLog.Title = bc.Title
		walletLog.BalanceAfter = walletLog.BalanceBefore + this.Amount
		walletLog.Pid = user.Pid
		walletLog.BalanceCode = this.BalanceCode
		walletLog.Note = this.Note
		walletLog.ParentPath = user.ParentPath
		if this.BalanceCode <= 0 {
			if wallet.Balance < realAmouont {
				return gerror.NewCode(gcode.New(-1000, "余额不足", ""))
			}
			walletLog.Amount = -this.Amount
			walletLog.BalanceAfter = walletLog.BalanceBefore - this.Amount
		}
		if walletLog.BalanceAfter < 0 {
			return gerror.NewCode(gcode.New(-1001, "余额错误", ""))
		}

		switch this.BalanceCode {
		case Bet:
			wallet.TotalBet += this.Amount
			wallet.TotalProfit -= this.Amount
		case Deposit:
			wallet.TotalDeposit += this.Amount
		case DepositBank:
			wallet.TotalDeposit += this.Amount
		case DepositGcash:
			wallet.TotalDeposit += this.Amount
		case Withdraw:
			wallet.TotalWithdraw += this.Amount
		case Profit:
			wallet.TotalProfit += this.Amount
		case Gift:
			wallet.TotalGift += this.Amount
		case Unfreeze:
			wallet.Freeze -= realAmouont
		case Freeze:
			wallet.Freeze += realAmouont
			//wallet.TotalFreeze += this.Amount
		}

		last, err := tx.Model(dao.Wallet.Table()).Data(g.Map{
			"balance":        walletLog.BalanceAfter * precision,
			"freeze":         wallet.Freeze,
			"total_bet":      wallet.TotalBet,
			"total_deposit":  wallet.TotalDeposit,
			"total_profit":   wallet.TotalProfit,
			"total_withdraw": wallet.TotalWithdraw,
			"total_gift":     wallet.TotalGift,
			//"total_freeze":   wallet.TotalFreeze,
		}).Where("uid=? and balance!=?", user.Id, walletLog.BalanceAfter*precision).Update()
		if err != nil {
			return err
		}
		affected, err := last.RowsAffected()
		if err != nil || affected == 0 {
			return err
		}
		code := entity.BalanceCode{}
		tx.Model(dao.BalanceCode.Table()).Scan(&code, "code", this.BalanceCode)
		if walletLog.Title == "" {
			walletLog.Title = code.Title
		}
		_, err = tx.Model(dao.WalletLog.Table()).Insert(&walletLog)
		if err != nil {
			return err
		}
		if fc != nil {
			err := fc(ctx, tx)
			if err != nil {
				return err
			}
		}

		r := report.Report{}
		r.User = user
		r.Amount = this.Amount
		r.BalanceCodeEntity = code
		go r.PutReport()
		return nil
	})

}
