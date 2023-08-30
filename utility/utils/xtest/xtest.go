package xtest

import (
	"context"
	"star_net/consts"
	"star_net/model"
)

func GetContextByUser(uid int64, account string) context.Context {

	return context.WithValue(context.Background(), consts.UserInfo, model.UserInfo{UidInt64: uid, Account: account})
}
func GetContext() context.Context {

	return context.WithValue(context.Background(), consts.UserInfo, model.UserInfo{UidInt64: 121, Account: "join"})
}
