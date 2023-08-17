package report

import (
	"context"
	"star_net/app/admin/api/report"
	"star_net/app/admin/internal/service/reportsvc"
)

var (
	Report = cRport{}
)

type cRport struct {
}

func (cRport) WalletLog(ctx context.Context, req *report.WalletLogReq) (res *report.WalletLogRes, err error) {
	log := reportsvc.WalletLog
	log.Begin = req.Begin
	log.End = req.End
	log.Account = req.Account
	log.Page = req.Page
	log.Size = req.Size
	logRes, err := log.Report(ctx)
	res = &report.WalletLogRes{}
	res.List = logRes.List
	res.Total = logRes.Total
	res.Sum = logRes.Sum
	return res, err
}
