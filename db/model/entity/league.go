// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// League is the golang structure for table league.
type League struct {
	Id      int64  `json:"id"      description:"编号"`
	Title   string `json:"title"   description:"标题"`
	Country string `json:"country" description:"国家"`
	Status  int    `json:"status"  description:"状态"`
	Name    string `json:"name"    description:"中文名称"`
	EnName  string `json:"enName"  description:"英文名称"`
}
