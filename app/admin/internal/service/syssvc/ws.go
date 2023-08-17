package syssvc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"star_net/model"
	"star_net/pure/get"
)

var (
	users  = gmap.New(true)
	admins = gmap.New(true)
)

func SendAdminMsg(ctx context.Context, d *model.AdminMsg) error {
	toAdmin, err := get.AdminByUname(ctx, d.ToUname)
	if err != nil {
		return err
	}
	d.ToUid = uint64(toAdmin.Id)
	d.FromUname = get.AdminUnameFromCtx(ctx)
	return NoticeAdmin(ctx, d, uint64(toAdmin.Id))
}

func NoticeAdmin(ctx context.Context, msg interface{}, uid uint64) error {
	to := admins.Get(uid)
	if to != nil {
		marshal, _ := json.Marshal(msg)
		if err := to.(*ghttp.WebSocket).WriteMessage(1, marshal); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
	}
	return nil
}
func GetAdminWs(ctx context.Context) error {
	var (
		r = ghttp.RequestFromCtx(ctx)
	)
	ws, err := r.WebSocket()
	if err != nil {
		return err
	}
	uid := get.AdminIdFromCtx(ctx)
	if uid == 0 {
		return errors.New("链接失败，获取UID为空")
	}
	value := admins.Get(uid)
	if value != nil {
		if err = value.(*ghttp.WebSocket).Close(); err != nil {
			g.Log().Error(r.Context(), err)
		}
		admins.Remove(uid)
	}
	admins.Set(uid, ws)
	printAdminWs()
	for {
		messageType, msg, err := ws.ReadMessage()
		if err != nil {
			if err = ws.Close(); err != nil {
				g.Log().Error(ctx, err)
			}
			admins.Remove(uid)
			printAdminWs()
			return nil
		}
		g.Log().Info(gctx.New(), "ws:lSys msg ", messageType, msg)
	}
}

func printAdminWs() {
	g.Log().Infof(gctx.New(), "admin连接个数%v %v", len(admins.Map()), admins.Keys())
}
