package sys

import (
	"context"
	"star_net/app/admin/api/sys"
	"star_net/app/admin/internal/service/syssvc"
)

var (
	Role = cRole{}
)

type cRole struct {
}

func (c cRole) GetById(ctx context.Context, req *sys.GetRoleByIdReq) (res sys.GetRoleByIdRes, _ error) {
	res = sys.GetRoleByIdRes{}
	id, menus, err := syssvc.GetRoleById(ctx, req.Id)
	//res.Menu = menus
	//res.Name = id.Name
	//res.Status = id.Status
	//res.OrderNo = id.OrderNo
	//res.CreatedAt = id.CreatedAt
	//res.UpdatedAt = id.UpdatedAt
	res.Role = id
	res.Menu = menus
	return res, err
}

func (c *cRole) Add(ctx context.Context, req *sys.AddRoleReq) (res *sys.AddRoleRes, err error) {
	err = syssvc.Role.Add(ctx, req.Role)
	return nil, err
}

func (c *cRole) Edit(ctx context.Context, req *sys.EditRoleReq) (res *sys.EditRoleRes, err error) {
	err = syssvc.Role.Edit(ctx, req.Role)
	return nil, err
}

func (c *cRole) Del(ctx context.Context, req *sys.DelRoleReq) (res *sys.DelRoleRes, err error) {
	err = syssvc.Role.Del(ctx, req.Id)
	return nil, err
}

func (c *cRole) List(ctx context.Context, req *sys.ListRoleReq) (res *sys.ListRoleRes, err error) {
	list, total, err := syssvc.Role.List(ctx)
	if err != nil {
		return nil, err
	}
	return &sys.ListRoleRes{
		Total: total,
		List:  list,
	}, nil
}
