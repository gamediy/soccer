package model

type CommonPageReq struct {
	Page int `json:"page" v:"required"`
	Size int `json:"size" v:"required"`
}

type CommonPageRes struct {
	Page      int `json:"page" d:"1" dc:"当前页数"`
	TotalPage int `json:"total_page,omitempty" dc:"总页数"`
	Size      int `json:"size,omitempty" dc:"页面大小"`
	Total     int `json:"total" dc:"总条数"`
}

type CommonRes struct{}
type AdminMsg struct {
	FromUname string `json:"from_uname"`
	Position  string `json:"position"`
	ToUname   string `json:"to_uname"`
	ToUid     uint64 `json:"to_uid"`
	Msg       string `v:"required" json:"msg"`
	Type      string `v:"required" d:"info" json:"type"` // 用于通知类型
}
