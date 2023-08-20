package play

import "star_net/core/soccer"

type p1000 struct {
}

type p1001 struct {
}

// 小
func (p1000) WonCheck(openResult soccer.OpenResult, calcRule string) (err error) {

	return nil
}

// 大
func (p1001) WonCheck(openResult soccer.OpenResult, calcRule string) (err error) {

	return nil
}
