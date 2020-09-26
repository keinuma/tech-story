package gateway

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"github.com/keinuma/go-graphql/api/domain/model"
	"github.com/keinuma/go-graphql/api/infra/database/dao"
)

type User struct {
	ctx context.Context
	tx  *gorm.DB
}

func NewUser(ctx context.Context, tx *gorm.DB) *User {
	return &User{
		ctx: ctx,
		tx:  tx,
	}
}

func (u *User) CreateUser(user model.User) (*model.User, error) {
	var daoUser dao.User
	daoUser = daoUser.ToDAO(user)
	if err := u.tx.Create(&daoUser).Error; err != nil {
		return nil, errors.New("[gateway.CreateUser] failed create user")
	}
	entityStory, err := daoUser.ToEntity()
	if err != nil {
		return nil, err
	}
	return entityStory, nil
}
