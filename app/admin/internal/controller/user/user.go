package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/admin/api/user"
	"star_net/app/admin/internal/service/usersvc"
	"star_net/model"
	"star_net/utility/utils/xcrud"
	"star_net/utility/utils/xpwd"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c cUser) Read(ctx context.Context, req *user.ReadUserReq) (_ *user.ReadUserRes, _ error) {
	var d user.ReadUserRes
	x := xcrud.Read{Ctx: ctx, Table: "u_user", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cUser) ReadList(ctx context.Context, req *user.ReadListUserReq) (_ *user.ReadListUserRes, _ error) {
	x := usersvc.ReadList{
		Ctx: ctx,
		Req: req,
	}
	d, total, err := x.Exec()
	if err != nil {
		return nil, err
	}
	return &user.ReadListUserRes{List: d, Total: total}, nil
}
func (c cUser) Create(ctx context.Context, req *user.CreateUserReq) (_ *user.CreateUserRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "u_user", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cUser) Del(ctx context.Context, req *user.DelUserReq) (_ *user.DelUserRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "u_user", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cUser) Update(ctx context.Context, req *user.UpdateUserReq) (_ *user.UpdateUserRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "u_user", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cUser) PutPass(ctx context.Context, req *user.PutPassReq) (_ *model.CommonRes, _ error) {
	x := xcrud.Update{
		Ctx:   ctx,
		Table: "u_user",
		Data:  g.Map{"password": xpwd.GenPwd(req.Pass)},
		V:     req.Id,
	}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
