package sys

import (
	"context"
	"star_net/app/admin/api/sys"
	"star_net/app/admin/internal/service/syssvc"
	"star_net/model"
)

var (
	Ws = cWs{}
)

type cWs struct {
}

func (c *cWs) AdminWs(ctx context.Context, _ *sys.AdminWsReq) (res *sys.AdminWsRes, err error) {
	if err = syssvc.GetAdminWs(ctx); err != nil {
		return nil, err
	}
	return
}

func (c *cWs) SendMsg(ctx context.Context, req *sys.SendMsgReq) (res *sys.SendMsgRes, err error) {
	var (
		d = model.AdminMsg{}
	)
	d.Position = req.Position
	d.ToUname = req.ToUname
	d.ToUid = req.ToUid
	d.Msg = req.Msg
	d.Type = req.Type
	if err = syssvc.SendAdminMsg(ctx, &d); err != nil {
		return nil, err
	}
	return
}
