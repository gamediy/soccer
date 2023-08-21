package sys

import (
	"context"
	"star_net/app/admin/api/sys"
	"star_net/utility/utils/xpusher"
)

var (
	Pusher = cPusher{}
)

type cPusher struct{}

func (c cPusher) SendAdmins(ctx context.Context, req *sys.SendAdminsReq) (_ *sys.SendAdminsRes, _ error) {
	if err := xpusher.SendAdmins(req.Data); err != nil {
		return nil, err
	}
	return
}
