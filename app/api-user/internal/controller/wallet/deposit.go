package wallet

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/app/api-user/api/wallet"
	"star_net/app/api-user/internal/service/depositsvc"
)

func (c cWallet) CreateDeposit(ctx context.Context, req *wallet.CreateDepositReq) (res *wallet.CreateDepositRes, _ error) {
	x := depositsvc.Submit{
		PayId:           req.PayId,
		Amount:          req.Amount,
		TranserImg:      req.TranserImg,
		TransferOrderNo: req.TransferOrderNo,
	}
	exec, err2 := x.Exec(ctx)
	res = &wallet.CreateDepositRes{}
	if err2 != nil {
		res.OrderNo = ""
		return res, err2
	}
	res.OrderNo = exec
	return res, nil
}
func (c cWallet) ListPlatformDeposit(ctx context.Context, _ *wallet.ListPlatformDepositReq) (res *wallet.ListPlatformDepositRes, err error) {

	data, err := depositsvc.GetDepositItem.Exec(ctx)
	if err != nil {
		return nil, err
	}
	res = &wallet.ListPlatformDepositRes{}
	gconv.Struct(data, res)
	return res, nil
}
func (c cWallet) DepositRecord(ctx context.Context, req *wallet.DepositRecordReq) (res *wallet.DepositRecordRes, err error) {
	depositsvc.Record.Page = req.Page
	depositsvc.Record.Size = req.Size
	depositsvc.Record.OrderNo = req.OrderNo
	total, data, err := depositsvc.Record.Exec(ctx)
	if err != nil {
		return nil, err
	}
	res = &wallet.DepositRecordRes{}
	res.List = data
	res.Total = total
	return res, nil
}
