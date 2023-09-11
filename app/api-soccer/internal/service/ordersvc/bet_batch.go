package ordersvc

import (
	"context"
	"fmt"
)

type BetBatchInput struct {
	List []BetInput
}
type BetBatchOutput struct {
	Errors []string
}
type BetBatch struct {
	BetBatchInput
}

func (input *BetBatch) Exec(ctx context.Context) (BetBatchOutput, error) {

	erros := BetBatchOutput{}
	for _, item := range input.List {
		b := Bet{}
		b.OddsId = item.OddsId
		b.Amount = item.Amount
		err := b.Exec(ctx)
		if err != nil {
			erros.Errors = append(erros.Errors, fmt.Sprintf("%d %s", b.OddsId, err.Error()))
		}
	}
	return erros, nil

}
