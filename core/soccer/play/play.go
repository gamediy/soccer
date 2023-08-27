package play

const (
	PlayTypeHandicap = "Handicap"
)

var (
	List = make(map[int]Play)
)

func init() {
	//大球
	List[1000] = P100Utils{}
	List[1001] = P100Utils{}
	List[1002] = P100Utils{}
	List[1003] = P100Utils{}
	List[1004] = P100Utils{}
	List[1005] = P100Utils{}
	List[1006] = P100Utils{}
	List[1007] = P100Utils{}
	List[1008] = P100Utils{}
	List[1009] = P100Utils{}

	//小球
	List[1100] = P100Utils{}
	List[1101] = P100Utils{}
	List[1102] = P100Utils{}
	List[1103] = P100Utils{}
	List[1104] = P100Utils{}
	List[1105] = P100Utils{}
	List[1106] = P100Utils{}
	List[1107] = P100Utils{}
	List[1108] = P100Utils{}
	List[1109] = P100Utils{}

	//独赢
	List[2000] = P2000{}
	List[2001] = P2001{}
	List[2002] = P2002{}
	//波胆
	List[3000] = P300Utils{}
	List[3001] = P300Utils{}
	List[3002] = P300Utils{}
	List[3003] = P300Utils{}
	List[3004] = P300Utils{}
	List[3005] = P300Utils{}
	List[3006] = P300Utils{}
	List[3007] = P300Utils{}
	List[3008] = P300Utils{}
	List[3009] = P300Utils{}
	List[3010] = P300Utils{}
	List[3011] = P300Utils{}
	List[3012] = P300Utils{}
	//让球
	List[4000] = P400Utils{}
	List[4001] = P400Utils{}
	List[4002] = P400Utils{}
	List[4003] = P400Utils{}
	List[4004] = P400Utils{}
	List[4005] = P400Utils{}
	List[4006] = P400Utils{}
	List[4007] = P400Utils{}
	List[4008] = P400Utils{}
	List[4009] = P400Utils{}
	List[4010] = P400Utils{}
	List[4011] = P400Utils{}
	List[4012] = P400Utils{}
	List[4013] = P400Utils{}
	List[4014] = P400Utils{}
	List[4015] = P400Utils{}
	List[4016] = P400Utils{}
	List[4017] = P400Utils{}
	List[4018] = P400Utils{}

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
	Handicap  float64 //让球
}
