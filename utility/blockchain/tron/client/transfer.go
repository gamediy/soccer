package client

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/craftto/go-tron/pkg/common"
	"github.com/craftto/go-tron/pkg/keystore"
	"github.com/craftto/go-tron/pkg/proto/api"
	"github.com/craftto/go-tron/pkg/proto/core"
	"github.com/craftto/go-tron/pkg/transaction"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/protobuf/proto"
	"math"
	"math/big"
)

// Transfer from to base58 address
func (g *GrpcClient) Transfer(from, toAddress string, amount int64) (*api.TransactionExtention, error) {
	var err error

	contract := &core.TransferContract{}
	if contract.OwnerAddress, err = common.DecodeBase58(from); err != nil {
		return nil, err
	}
	if contract.ToAddress, err = common.DecodeBase58(toAddress); err != nil {
		return nil, err
	}
	contract.Amount = amount

	ctx, cancel := g.GetContext()
	defer cancel()

	tx, err := g.Client.CreateTransaction2(ctx, contract)
	if err != nil {
		return nil, err
	}

	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("bad transaction")
	}

	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}
	return tx, nil
}
func (g *GrpcClient) TransferTrx(to_addr string, from_addr string, from_privatekey string, amount float64) (*api.Return, error) {

	av := math.Pow10(6) * amount
	transfer, err := TronGrpcClient.Transfer(from_addr, to_addr, int64(av))
	if err != nil {
		return nil, err
	}
	h256h := sha256.New()
	marshal, err := proto.Marshal(transfer.GetTransaction().RawData)
	if err != nil {
		return nil, err
	}
	h256h.Write(marshal)
	hash := h256h.Sum(nil)
	privkey, err := crypto.HexToECDSA(from_privatekey)
	if err != nil {
		return nil, err
	}
	signature, err := crypto.Sign(hash, privkey)
	if err != nil {
		return nil, err
	}
	transfer.Transaction.Signature = append(transfer.Transaction.Signature, signature)
	transaction, err := TronGrpcClient.Client.BroadcastTransaction(context.Background(), transfer.Transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
func (g *GrpcClient) TransferUsdt(to_addr string, from_privatekey string, amount float64) (*transaction.Transaction, error) {
	trc20, err := TronGrpcClient.NewTrc20("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")
	key, err := keystore.ImportFromPrivateKey(from_privatekey)
	if err != nil {
		return nil, err
	}
	av := math.Pow10(6) * amount
	transfer, err := trc20.Transfer(key, to_addr, big.NewInt(int64(av)))
	return transfer, err

}
