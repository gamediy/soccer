package reportsvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"strings"
)

var (
	WalletLog = &walletLog{}
)

type walletLog struct {
	Begin   string `json:"begin"`   //开始时间
	End     string `json:"end"`     //结束时间
	Account string `json:"account"` //账号
	Page    int
	Size    int
}

type walletLogRes struct {
	List  gdb.List
	Total int64
	Sum   gdb.List
}

func (s *walletLog) Report(ctx context.Context) (res *walletLogRes, err error) {
	defer func() {
		if res.List == nil {
			res.List = make(gdb.List, 0)
			res.Total = 0
		}
	}()
	if s.Begin == "" {
		s.Begin = gtime.Now().AddDate(0, -6, 0).String()
	}
	if s.End == "" {
		s.End = gtime.Now().String()
	}

	bList := []*entity.BalanceCode{}
	dao.BalanceCode.Ctx(ctx).Where("status", 1).Scan(&bList)
	var builder strings.Builder
	var total strings.Builder
	var sum strings.Builder
	total.WriteString(`SELECT
         count(distinct uid) as count
		 from r_report_wallet_day

`)
	sum.WriteString(`
	SELECT
      sum(amount) as amount,
      title
		 from		r_report_wallet_day
`)
	builder.WriteString(`SELECT
							uid,
							date,
							account,`)
	builder.WriteString(fmt.Sprintf("\n"))
	for index, code := range bList {
		builder.WriteString(fmt.Sprintf("SUM(CASE WHEN balance_code = %d THEN count ELSE 0 END) AS c_%s,\n", code.Code, code.Title))
		builder.WriteString(fmt.Sprintf("SUM(CASE WHEN balance_code = %d THEN amount ELSE 0 END) AS a_%s", code.Code, code.Title))
		if index+1 < len(bList) {
			builder.WriteString(",\n")
		}
	}
	builder.WriteString(` FROM
							r_report_wallet_day `)
	var where strings.Builder
	where.WriteString(fmt.Sprintf("\n where date>='%s' and date<='%s'", s.Begin, s.End))
	if s.Account != "" {
		where.WriteString(fmt.Sprintf(" and account='%s'", s.Account))
	}
	builder.WriteString(where.String())
	total.WriteString(where.String())
	sum.WriteString(where.String())
	builder.WriteString(" group by  uid,account,date")
	sum.WriteString(" group by balance_code")
	builder.WriteString(fmt.Sprintf(" order by id desc limit %d offset %d ", s.Size, s.Page*s.Size))
	one, err := g.DB().Ctx(ctx).GetOne(ctx, total.String())
	res = &walletLogRes{}
	if one == nil {
		return res, nil
	}
	res.Total = one.Map()["count"].(int64)
	if res.Total != 0 {
		all, _ := g.DB().Ctx(ctx).GetAll(ctx, builder.String())
		sum, _ := g.DB().Ctx(ctx).GetAll(ctx, sum.String())
		res.List = all.List()
		res.Sum = sum.List()
	}
	return res, nil
}
