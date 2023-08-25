package play

const (
	PlayTypeHandicap = "Handicap"
)

var (
	List = make(map[int]Play)
)

func init() {
	List[1000] = P1000{}
	List[1001] = P1001{}
}

type Play interface {
	WonCheck(openResult OpenResult, calcInfo CalcInfo) (profit float64) //结算

}

type OpenResult struct {
	Result     string
	BoutStatus int
}
type CalcInfo struct {
	Rule      string
	Odds      float64
	BetAmount float64
}
