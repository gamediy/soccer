package play

import (
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/utility/utils/xplay"
)

type P1000 struct {
}

type P1001 struct {
}

// 小球
func (P1000) WonCheck(openResult OpenResult, calcInfo CalcInfo) float64 {

	total, err := xplay.OpenResutToTotal(openResult.Result)
	if err != nil {
		return 0
	}
	if total < gconv.Int(calcInfo.Rule) {
		return calcInfo.Odds * calcInfo.BetAmount
	}
	return 0
}

// 大球
func (P1001) WonCheck(openResult OpenResult, calcInfo CalcInfo) float64 {
	total, err := xplay.OpenResutToTotal(openResult.Result)
	if err != nil {
		return 0
	}
	if total > gconv.Int(calcInfo.Rule) {
		return calcInfo.Odds * calcInfo.BetAmount
	}
	return 0
}
