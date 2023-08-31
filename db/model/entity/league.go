// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// League is the golang structure for table league.
type League struct {
	Id      int64  `json:"id"      description:"编号"`
	Country string `json:"country" description:"国家"`
	Status  int    `json:"status"  description:"状态"`
	ZhName  string `json:"zhName"  description:"中文名称"`
	EnName  string `json:"enName"  description:"英文名称"`
	Icon    int    `json:"icon"    description:""`
	VdName  string `json:"vdName"  description:""`
}
