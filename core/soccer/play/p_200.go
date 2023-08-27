package play

import (
	"star_net/utility/utils/xplay"
)

type P2000 struct {
}

type P2001 struct {
}
type P2002 struct {
}

// home
func (P2000) WonCheck(openResult OpenResult, calcInfo CalcInfo) float64 {
	a1, a2, err := xplay.OpenResutToTwo(openResult.Result)
	if err != nil {
		return 0
	}
	if a1 > a2 {
		return calcInfo.BetAmount * calcInfo.Odds
	}
	return 0
}

// away
func (P2001) WonCheck(openResult OpenResult, calcInfo CalcInfo) float64 {

	a1, a2, err := xplay.OpenResutToTwo(openResult.Result)
	if err != nil {
		return 0
	}
	if a1 < a2 {
		return calcInfo.BetAmount * calcInfo.Odds
	}
	return 0
}

// tie
func (P2002) WonCheck(openResult OpenResult, calcInfo CalcInfo) float64 {

	a1, a2, err := xplay.OpenResutToTwo(openResult.Result)
	if err != nil {
		return 0
	}
	if a1 == a2 {
		return calcInfo.BetAmount * calcInfo.Odds
	}
	return 0
}
