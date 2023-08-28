package play

import (
	"github.com/gogf/gf/v2/util/gconv"
	"math"
	"star_net/utility/utils/xplay"
	"strings"
)

// 让球
type P400Utils struct {
}

func (P400Utils) WonCheck(openResult OpenResult, calcInfo CalcInfo) float64 {

	return getP400Profit(openResult, calcInfo)
}

func getP400Profit(result OpenResult, info CalcInfo) float64 {

	n1, n2, err := xplay.OpenResutToTwo(result.Result)
	if err != nil {
		return 0
	}

	split := strings.Split(info.Rule, " ")
	team := split[0]
	info.Rule = split[1]
	nu := strings.Split(info.Rule, "/")
	hc := math.Abs(gconv.Float64(nu[0]))

	if len(nu) > 1 {
		nn2 := gconv.Float64(nu[1])
		hc = nn2 - 0.25
	}

	var score float64
	addCon := strings.Contains(info.Rule, "+")
	if team == "主" {
		if addCon {
			score = n1 + hc - n2
		} else {
			score = n1 - hc - n2
		}

	} else {
		if addCon {
			score = n2 + hc - n1
		} else {
			score = n2 - hc - n1
		}
	}
	if score == 0 {
		return info.BetAmount
	} else if score == 0.25 {
		return info.Odds * info.BetAmount * 0.5
	} else if score > 0.25 {
		return info.Odds * info.BetAmount
	} else if score == -0.25 {
		return info.BetAmount * 0.5
	}

	return 0

}
