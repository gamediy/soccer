package ordersvc

import (
	"context"
	"fmt"
)

var (
	BetBatch = betBatch{}
)

type BetBatchInput struct {
	List []BetInput
}
type BetBatchOutput struct {
	Errors []string
}
type betBatch struct {
	BetBatchInput
}

func (input *betBatch) Exec(ctx context.Context) (BetBatchOutput, error) {

	erros := BetBatchOutput{}
	for _, item := range input.List {
		b := bet{}
		b.OddsId = item.OddsId
		b.Amount = item.Amount
		err := b.Exec(ctx)
		if err != nil {
			erros.Errors = append(erros.Errors, fmt.Sprintf("%s %s", b.OddsId, err.Error()))
		}
	}
	return erros, nil

}
