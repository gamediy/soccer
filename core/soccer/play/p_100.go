package play

import "star_net/core/soccer"

type P1000 struct {
}

type P1001 struct {
}

// 小
func (P1000) WonCheck(openResult soccer.OpenResult, calcRule string) (err error) {

	return nil
}

// 大
func (P1001) WonCheck(openResult soccer.OpenResult, calcRule string) (err error) {

	return nil
}
