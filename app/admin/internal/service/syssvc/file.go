package syssvc

import (
	"context"
	"fmt"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xfile"
)

type DelFile struct {
	Id uint64
}

func (s DelFile) Exec(ctx context.Context) error {
	var file entity.File
	_ = dao.File.Ctx(ctx).Scan(&file, "id", s.Id)
	if file.Id == 0 {
		return fmt.Errorf("数据不存在")
	}
	x := xfile.NewCloudFlareFromCtx(ctx)
	if err := x.Del(ctx, file.Url); err != nil {
		return err
	}
	if _, err := dao.File.Ctx(ctx).Delete("id", s.Id); err != nil {
		return err
	}
	return nil
}
