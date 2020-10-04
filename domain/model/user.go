package model

import (
	"fmt"
	"io"
	"strconv"
)

type Users []*User

type User struct {
	ID          int      `json:"id"`
	UID         string   `json:"uid"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Role        UserRole `json:"-"`
}

type UserRole struct {
	Role Role `json:"-"`
}

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser:
		return true
	}
	return false
}

func (e Role) String() string {
	switch e {
	case RoleAdmin:
		return "管理者"
	case RoleUser:
		return "ユーザー"
	}
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
