package play

import (
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/utility/utils/xplay"
	"strings"
)

type P1000 struct {
}

type P1001 struct {
}

// 100通用的结算
type P100Utils struct {
}

func (P100Utils) WonCheck(openResult OpenResult, calcInfo CalcInfo) float64 {

	return getP100Profit(openResult, calcInfo)
}

// Over 1
func (P1000) WonCheck(openResult OpenResult, calcInfo CalcInfo) float64 {

	return getP100Profit(openResult, calcInfo)
}

// Over 1/1.5
func (P1001) WonCheck(openResult OpenResult, calcInfo CalcInfo) float64 {

	return getP100Profit(openResult, calcInfo)
}

func getP100Profit(openResult OpenResult, calcInfo CalcInfo) float64 {
	split := strings.Split(calcInfo.Rule, " ")
	if len(split) < 1 {
		return 0
	}
	total, err := xplay.OpenResutToTotal(openResult.Result)
	if err != nil {
		return 0
	}
	ts := split[0]
	nu := strings.Split(split[1], "/")

	if ts == "Over" { //大
		n1 := gconv.Float64(nu[0])
		if total >= n1 {
			if len(nu) == 1 {
				return calcInfo.Odds * calcInfo.BetAmount
			} else {
				n2 := gconv.Float64(nu[1])
				if total > n2 {
					return calcInfo.Odds * calcInfo.BetAmount
				} else {
					return calcInfo.Odds * calcInfo.BetAmount * 0.5
				}
			}

		}

	} else if ts == "Under" { //小

	}

	return 0
}
