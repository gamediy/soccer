package tron

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/utility/blockchain/tron/client"
	"testing"
	"time"
)

func Test_CreateAccount(t *testing.T) {
	g.Dump(client.TronGrpcClient.GenerateKey())
}
func Test_GetTrxBalance(t *testing.T) {
	account, err2 := client.TronGrpcClient.GetBalanceTrx("TSk4erYrpiAywiBRvwUMPPYKmjWj55N4Qc")
	g.Dump(account, err2)

}

func Test_TransferTrx(t *testing.T) {
	g.Dump(client.TronGrpcClient.TransferTrx("TSk4erYrpiAywiBRvwUMPPYKmjWj55N4Qc", "TPof1kGiU1fKK7pJLvyaEeXMXJs8WQD9Ad", "", 0.001))
}

func Test_Abi(t *testing.T) {
	for {
		g.Dump(client.TronGrpcClient.GenerateKey())
		g.Dump(client.TronGrpcClient.GetBalanceUsdt("TPof1kGiU1fKK7pJLvyaEeXMXJs8WQD9Ad"))
		time.Sleep(3 * time.Second)
	}

}

func Test_TransferUSDT(t *testing.T) {
	g.Dump(client.TronGrpcClient.TransferUsdt("TLRzQLaVXgYkSvgym5CSgpvxx6VTpJ15Uh", "", 0.001))
}
