package syssvc

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"star_net/consts"
	"star_net/db/dao"
	"star_net/utility/utils/xcasbin"
	"star_net/utility/utils/xelastic"
	"star_net/utility/xpusher"
)

type InitD struct{}

func (s *InitD) Exec(ctx context.Context) {
	// set log
	g.Log().SetFlags(glog.F_FILE_LONG | glog.F_TIME_DATE | glog.F_TIME_MILLI)
	// set imgPrefix
	if err := s.SetImgPrefix(ctx); err != nil {
		panic(err)
	}

	// set White Ips
	if err := s.SetWhiteIps(ctx); err != nil {
		panic(err)
	}

	xpusher.InitFromCfg(ctx) // init pusher
	xcasbin.InitFromCfg(ctx) // init casbin
	xelastic.InitCfg(ctx)    // init es
	xelastic.SaveEsLog("admin_log_")
}

func (s *InitD) SetImgPrefix(ctx context.Context) error {
	value, err := dao.Dict.Ctx(ctx).Value("v", "k= 'cloudflare_pub'")
	if err != nil {
		return err
	}
	consts.ImgPrefix = value.String()
	return nil
}

func (s *InitD) SetWhiteIps(ctx context.Context) error {
	d, err := dao.Dict.Ctx(ctx).Value("v", "k='white_ips'")
	if err != nil {
		return err
	}
	consts.WhiteIps = d.String()
	return nil
}
