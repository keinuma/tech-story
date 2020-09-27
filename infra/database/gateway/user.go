package gateway

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"github.com/keinuma/tech-story/domain/model"
	"github.com/keinuma/tech-story/infra/database/dao"
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

func (u *User) GetUsersByIDs(userIDs []int) (model.Users, error) {
	var daoUsers dao.Users
	if err := u.tx.Where("id IN (?)", userIDs).Find(&daoUsers).Error; err != nil {
		return nil, errors.New("[gateway.GetUsersByIDs] failed get users by ID")
	}
	entityUsers, err := daoUsers.ToEntity()
	if err != nil {
		return nil, err
	}
	return entityUsers, nil
}

func (u *User) GetUserByUID(uid string) (*model.User, error) {
	var daoUser dao.User
	if err := u.tx.Where("uid = ?", uid).First(&daoUser).Error; err != nil {
		return nil, errors.New("[gateway.GetUserByUID] failed get user by UID")
	}
	entityUser, err := daoUser.ToEntity()
	if err != nil {
		return nil, err
	}
	return entityUser, nil
}
