package push

import (
	"star_net/utility/utils/xpusher"
)

const (
	ChannelAdmin = "admins"
)
const (
	EventDeposit  = "ordersvc"
	EventWithdraw = "withdraw"
)

type MessageItem struct {
	Title string
	Body  string
}
type Message struct {
	EN MessageItem `json:"en"`
	CH MessageItem `json:"ch"`
}

func (m *Message) Trigger(channel string, event string) error {
	err := xpusher.Trigger(channel, event, m)
	return err

}
