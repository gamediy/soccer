package xpusher

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pusher/pusher-http-go"
)

var client pusher.Client

func SendAdmins(data interface{}) error {
	return Trigger("admins", "share", data)
}
func Trigger(channel, event string, data interface{}) error {
	err := client.Trigger(channel, event, data)
	if err != nil {
		return err
	}

	return nil
}
func InitFromCfg(ctx context.Context) {
	appId := g.Cfg().MustGet(ctx, "pusher.appId").String()
	key := g.Cfg().MustGet(ctx, "pusher.key").String()
	secret := g.Cfg().MustGet(ctx, "pusher.secret").String()
	cluster := g.Cfg().MustGet(ctx, "pusher.cluster").String()
	Init(appId, key, secret, cluster)
}
func Init(appID, key, secret, cluster string) {
	client = pusher.Client{
		AppID:   appID,
		Key:     key,
		Secret:  secret,
		Cluster: cluster,
	}
}
