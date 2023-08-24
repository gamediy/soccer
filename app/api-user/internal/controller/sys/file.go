package sys

import (
	"context"
	"star_net/app/api-user/api/sys"
	"star_net/utility/utils/xfile"
)

func (c sSys) UploadFile(ctx context.Context, req *sys.UploadFileReq) (res *sys.UploadFileRes, err error) {
	handler := xfile.NewCloudFlareFromCtx(ctx)
	name, err := handler.Upload(ctx, req.Group)
	if err != nil {
		return nil, err
	}
	return &sys.UploadFileRes{
		NewFileName: name,
	}, nil
}
