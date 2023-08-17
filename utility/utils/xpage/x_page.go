package xpage

import (
	"github.com/gogf/gf/v2/util/gconv"
	"math"
)

func MakePageInfo(pageIn, sizeIn, totalIn interface{}) (totalPage, total, page, size int64) {
	if sizeIn == 0 {
		sizeIn = 10
	}
	totalPage = int64(math.Ceil(gconv.Float64(totalIn) / gconv.Float64(sizeIn))) //这里计算总页数时，要向上取整
	if totalPage <= 0 {
		totalPage = 1
	}
	if gconv.Int64(pageIn) == 0 {
		pageIn = 1
	}
	page = gconv.Int64(pageIn)
	total = gconv.Int64(totalIn)
	size = gconv.Int64(sizeIn)
	return
}
