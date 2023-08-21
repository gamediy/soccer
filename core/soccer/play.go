package soccer

import "star_net/core/soccer/play"

const (
	PlayTypeHandicap = "Handicap"
)

var (
	PlayList = make(map[int]Play)
)

func init() {
	PlayList[1000] = play.P1000{}
	PlayList[1001] = play.P1001{}
}

type Play interface {
	WonCheck(openResult OpenResult, calcRule string) (err error) //结算

}

type OpenResult struct {
	Result     string
	BoutStatus int
}
