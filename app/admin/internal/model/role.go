package model

import "star_net/db/model/entity"

type Role struct {
	entity.Role
	Menu []*entity.RoleMenu `json:"menu"`
}
