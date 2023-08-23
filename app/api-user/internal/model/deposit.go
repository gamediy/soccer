package model

type AmountItem struct {
	PayId    int    `json:"payId" dc:"ID"`
	Title    string `json:"title"            description:"标题"`
	Protocol string `json:"protocol"         description:"协议"`
	Logo     string `json:"logo"             description:"Logo"`
	Currency string `json:"currency"         description:"货币单位"`
	Address  string `json:"address"          description:"地址或卡号"`
	Detail   string `json:"detail"           description:"详情"`
}
