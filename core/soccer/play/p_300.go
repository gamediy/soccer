package play

// 波胆
type P300Utils struct {
}

func (P300Utils) WonCheck(openResult OpenResult, calcInfo CalcInfo) float64 {
	if openResult.Result == calcInfo.Rule {
		return calcInfo.BetAmount * calcInfo.Odds
	}
	return 0
}
