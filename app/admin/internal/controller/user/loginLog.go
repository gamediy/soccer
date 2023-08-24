package user

import (
	"context"
	"star_net/app/admin/api/user"
	"star_net/utility/utils/xcrud"

	userModel "star_net/app/admin/internal/model"
)

var (
	UserLoginLog = cUserLoginLog{}
)

type cUserLoginLog struct{}

func (c cUserLoginLog) Read(ctx context.Context, req *user.ReadUserLoginLogReq) (_ *user.ReadUserLoginLogRes, _ error) {
	var d user.ReadUserLoginLogRes
	x := xcrud.Read{Ctx: ctx, Table: "u_user_login_log", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cUserLoginLog) ReadList(ctx context.Context, req *user.ReadListUserLoginLogReq) (_ *user.ReadListUserLoginLogRes, _ error) {
	var (
		d     = make([]*userModel.UserLoginLog, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx,
		Table: "u_user_login_log ",
		Page:  req.Page, Size: req.Size}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &user.ReadListUserLoginLogRes{List: d, Total: total}, nil
}
func (c cUserLoginLog) Create(ctx context.Context, req *user.CreateUserLoginLogReq) (_ *user.CreateUserLoginLogRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "u_user_login_log", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cUserLoginLog) Del(ctx context.Context, req *user.DelUserLoginLogReq) (_ *user.DelUserLoginLogRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "u_user_login_log", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cUserLoginLog) Update(ctx context.Context, req *user.UpdateUserLoginLogReq) (_ *user.UpdateUserLoginLogRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "u_user_login_log", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
