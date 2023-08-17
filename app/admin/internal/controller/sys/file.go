package sys

import (
	"context"
	"star_net/app/admin/api/sys"
	"star_net/app/admin/internal/service/syssvc"
	"star_net/db/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type cFile struct {
}

var (
	File = cFile{}
)

// UploadFile 上传文件
func (c *cFile) UploadFile(ctx context.Context, req *sys.UploadFileReq) (res *sys.UploadFilesRes, err error) {
	storerType := "cloud" // 或者 "cloud"
	rootFilePath, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		return nil, err
	}
	file := ghttp.RequestFromCtx(ctx).GetUploadFile("file")
	storer := syssvc.NewFileStorer(storerType, rootFilePath.String())
	var fileName string
	fileName, err = storer.Upload(ctx, req.Group, file)
	if err != nil {
		return nil, err
	}
	return &sys.UploadFilesRes{
		NewFileName: fileName,
	}, nil
}

// ListFile 查询列表
func (c *cFile) ListFile(ctx context.Context, req *sys.ListFileReq) (res *sys.ListFileRes, err error) {
	storerType := "local" // 或者 "cloud"
	rootFilePath, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		return nil, err
	}
	storer := syssvc.NewFileStorer(storerType, rootFilePath.String())
	var data []*entity.File
	var count int
	data, count, err = storer.ListFile(ctx, req)
	if err != nil {
		return nil, err
	}
	return &sys.ListFileRes{
		Total: count,
		List:  data,
	}, nil
}

// DeleteFile 删除文件
func (c *cFile) DeleteFile(ctx context.Context, req *sys.DelFileReq) (res *sys.DelFileRes, err error) {
	storerType := "cloud" // 或者 "cloud"
	rootFilePath, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		return nil, err
	}
	storer := syssvc.NewFileStorer(storerType, rootFilePath.String())
	err = storer.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &sys.DelFileRes{}, nil
}

// 详情查询 id
func (c *cFile) GetFile(ctx context.Context, req *sys.GetFileReq) (res *sys.GetFileRes, err error) {
	storerType := "local" // 或者 "cloud"
	rootFilePath, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		return nil, err
	}
	storer := syssvc.NewFileStorer(storerType, rootFilePath.String())
	var data *entity.File
	data, err = storer.GetFile(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &sys.GetFileRes{
		Data: data,
	}, nil
}
