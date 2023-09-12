package withdrawsvc

import (
	"context"
	"star_net/app/api-user/internal/service"
	"star_net/db/dao"
)

type WithdrawAccountList struct {
	Page, Size int
}

type WithdrawAccount struct {
	Id       string `json:"id"       description:""`
	Net      string `json:"net"      description:""`
	Protocol string `json:"protocol" description:""`
	Address  string `json:"address"  description:""`
	Currency string `json:"currency" description:"currency"`
	Status   int    `json:"status"   description:""`
	Default  int    `json:"default"  description:"是否默认的"`
	Title    string `json:"title"    description:""`
}

func (s WithdrawAccountList) Exec(ctx context.Context) (int, []*WithdrawAccount, error) {
	var (
		u     = service.GetUserInfo(ctx)
		d     = make([]*WithdrawAccount, 0)
		total int
	)
	if s.Size > 100 {
		s.Size = 10
	}
	if err := dao.WithdrawAccount.Ctx(ctx).Where("uid", u.Uid).OrderDesc("id").Page(s.Page, s.Size).ScanAndCount(&d, &total, false); err != nil {
		return 0, nil, err
	}
	return total, d, nil
}
