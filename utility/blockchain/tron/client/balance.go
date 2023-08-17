package client

import (
	"github.com/gogf/gf/v2/util/gconv"
)

func (g *GrpcClient) GetBalanceTrx(addr string) (float64, error) {
	account, err := TronGrpcClient.GetAccount(addr)
	if err != nil {
		return 0, err
	}
	f := gconv.Float64(gconv.Float64(account.GetBalance()) / 1000000)
	return f, nil
}

func (g *GrpcClient) GetBalanceUsdt(addr string) (float64, error) {
	trc20, err := TronGrpcClient.NewTrc20("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")
	if err != nil {
		return 0, err
	}
	balance, err := trc20.GetBalance(addr)
	return gconv.Float64(balance) / 1000000, err
}
