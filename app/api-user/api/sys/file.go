package sys

import "github.com/gogf/gf/v2/frame/g"

type UploadFileReq struct {
	g.Meta `tags:"系统" sm:"上传文件" method:"post" path:"/file/upload"`
	File   string `json:"file" v:"required"`
	Group  int    `json:"group" dc:"分组 1 系统 2 文件 3 凭证" v:"required#文件分组不能为空"`
}
type UploadFileRes struct {
	NewFileName string `json:"new_file_name" dc:"新文件名"`
}
