package deposit

import (
	"star_net/app/api-user/internal/service/deposit"
)

var (
	Ctrl    = &cDeposit{}
	Service = deposit.Service
)

type cDeposit struct {
}
