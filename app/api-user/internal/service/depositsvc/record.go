package depositsvc

import (
	"context"
	"star_net/app/api-user/internal/model"
	"star_net/app/api-user/internal/service"
	"star_net/db/dao"
)

func (s *deposit) Record(ctx context.Context, input model.DepositRecordInput) ([]*model.DepositRecordOutput, error) {

	userInfo := service.GetUserInfo(ctx)
	list := []*model.DepositRecordOutput{}
	dao.Deposit.Ctx(ctx).Where("uid", userInfo.Uid).Offset(input.Size * input.Page).Limit(input.Page).Scan(&list)
	return list, nil
}
