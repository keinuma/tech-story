package presenter

import (
	"encoding/json"
	"github.com/keinuma/go-graphql/api/domain/entity"
	"github.com/keinuma/go-graphql/api/domain/service"
	"github.com/keinuma/go-graphql/api/graph/model"
)

type UserHandler interface {
	CreateUser(input model.NewUser) (*model.User, error)
}

type User struct {
	userService service.User
}

func NewUser(userService service.User) *User {
	return &User{
		userService: userService,
	}
}

func (u *User) CreateUser(input model.NewUser) (*model.User, error) {
	entityUser := entity.User{
		UUID:        input.UUID,
		Name:        input.Name,
		Description: input.Description,
	}
	user, err := u.userService.CreateUser(entityUser)
	if err != nil {
		return nil, err
	}
	var userModel model.User
	bytes, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &userModel)
	if err != nil {
		return nil, err
	}
	return &userModel, err
}
