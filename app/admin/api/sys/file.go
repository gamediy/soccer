package sys

import (
	"star_net/db/model/entity"
	"star_net/model"

	"github.com/gogf/gf/v2/frame/g"
)

type UploadFileReq struct {
	g.Meta `tags:"/sys/file" dc:"上传图片" method:"post" path:"/file" summary:"上传图片"`
	Group  int `json:"group" dc:"分组:1头像,2图片,3动图,4音频,5文件" d:"2" in:"query"`
}

type UploadFilesRes struct {
	NewFileName string `json:"new_file_name" dc:"新文件名"`
}

// ListApiReq 查询列表
type ListFileReq struct {
	g.Meta `tags:"/sys/file" method:"get" path:"/file/list" dc:"文件查询列表"`
	model.CommonPageReq
}
type ListFileRes struct {
	Total int            `json:"total"`
	List  []*entity.File `json:"list"`
}

// DelApiReq 删除
type DelFileReq struct {
	g.Meta `tags:"/sys/file" method:"delete" path:"/file" dc:"文件删除单个"`
	Id     int `json:"id"`
}

type DelFileRes struct {
}

type GetFileReq struct {
	g.Meta `tags:"/sys/file" method:"get" path:"/file" dc:"文件查询单个"`
	Id     uint64 `json:"id"`
}

type GetFileRes struct {
	Data *entity.File
}

// 修改文件
type UpdateFileReq struct {
	g.Meta `tags:"/sys/file" method:"put" path:"/file" dc:"文件修改"`
	Id     uint64 `json:"id"`
	Group  int    `json:"group" dc:"分组:1头像,2图片,3动图,4音频,5文件" d:"2" in:"query"`
	Status int    `json:"status" dc:"状态:1启用,2禁用" d:"1" in:"query"`
}
