package presenter

import (
	"github.com/keinuma/tech-story/domain/model"
	"github.com/keinuma/tech-story/domain/service"
	"github.com/keinuma/tech-story/presenter/request"
)

type UserHandler interface {
	CreateUser(input request.NewUser) (*model.User, error)
}

type User struct {
	userService service.User
}

func NewUser(userService service.User) *User {
	return &User{
		userService: userService,
	}
}

func (u *User) CreateUser(input request.NewUser) (*model.User, error) {
	entityUser := model.User{
		UID:         input.UID,
		Name:        input.Name,
		Description: input.Description,
	}
	user, err := u.userService.CreateUser(entityUser)
	if err != nil {
		return nil, err
	}
	return user, err
}
