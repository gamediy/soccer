package sys

import "github.com/gogf/gf/v2/frame/g"

type GetDictReq struct {
	g.Meta `tags:"系统" sm:"获取字典信息" path:"/dict" method:"get"`
	K      string `json:"k" required:"k不能为空"`
}
type GetDictRes struct {
	V string `json:"v" dc:"值"`
}
type GetDictAllReq struct {
	g.Meta `tags:"系统" sm:"查询所有字典" path:"/dict/all" method:"get" dc:"返回的字段说明请具体访问一下接口"`
}
type GetDictAllRes g.Map
