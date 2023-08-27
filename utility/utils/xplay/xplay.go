package xplay

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

// 开奖结果转换成比分总数
func OpenResutToTotal(openResult string) (float64, error) {
	two, i, err := OpenResutToTwo(openResult)
	if err != nil {
		return 0, err
	}
	return two + i, nil
}

// 开奖结果转换成两个比分
func OpenResutToTwo(openResult string) (float64, float64, error) {
	split := strings.Split(openResult, "-")
	if len(split) < 2 {
		return 0, 0, fmt.Errorf("长度错误")
	}
	return gconv.Float64(split[0]), gconv.Float64(split[1]), nil
}
