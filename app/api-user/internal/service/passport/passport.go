package spassport

import (
	"context"
	"star_net/app/api-user/internal/model"
)

var (
	Service = &passport{}
)

type passport struct {
}

type Passport interface {
	Login(ctx context.Context, input model.LoginInput) error
	Register(ctx context.Context, input model.RegisterInput) error
}
