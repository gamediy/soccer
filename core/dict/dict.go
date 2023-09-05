package dict

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"star_net/db/dao"
	"time"
)

const (
	ImageDomain = "image_domain"
)

func GetVAsString(ctx context.Context, key string) string {
	one, err := dao.Dict.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     "dict:" + key,
		Force:    false,
	}).Where("k", key).One()
	if err != nil || one == nil {
		return ""
	}
	return one.Map()["v"].(string)
}
