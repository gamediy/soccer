```text
// Get$Dict$Req 查询单个
type Get$Dict$Req struct {
	g.Meta `tags:"/$sys$/$dict$" method:"get" path:"/$dict$" dc:"查询单个"`
	Id     uint64 `json:"id"`
}

type Get$Dict$Res entity.$Dict$
// List$Dict$Req 查询列表
type List$Dict$Req struct {
	g.Meta `tags:"/$sys$/$dict$" method:"get" path:"/$dict$/list" dc:"查询列表"`
	model.CommonPageReq
}
type List$Dict$Res struct {
	Total int            `json:"total"`
	List  []*entity.$Dict$ `json:"list"`
}

// Add$Dict$Req 添加
type Add$Dict$Req struct {
	g.Meta `tags:"/$sys$/$dict$" method:"post" path:"/$dict$" dc:"添加"`
	*entity.$Dict$
}
type Add$Dict$Res model.CommonRes


// Update$Dict$Req 修改
type Update$Dict$Req struct {
	g.Meta `tags:"/$sys$/$dict$" method:"put" path:"/$dict$" dc:"修改单个数据"`
	*entity.$Dict$
}
type Update$Dict$Res model.CommonRes

// Del$Dict$Req 删除
type Del$Dict$Req struct {
	g.Meta `tags:"/$sys$/$dict$" method:"delete" path:"/$dict$" dc:"删除单个数据"`
	Id     uint64 `json:"id"`
}

type Del$Dict$Res model.CommonRes
```