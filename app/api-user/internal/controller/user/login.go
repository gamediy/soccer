package user

import (
	"context"
	"fmt"
	"star_net/app/api-user/api/user"
	"star_net/app/api-user/internal/service/depositsvc"
)

func (c cUser) Login(ctx context.Context, _ *user.LoginReq) (_ *user.LoginRes, err error) {

	s := depositsvc.GetDepositItem
	s.Name = "11"
	fmt.Println(&s.Name)
	return
}
