package services

import (
	"context"
	db "github.com/wecanooo/kora/app/models"
	"github.com/wecanooo/kora/core"
)

type IUserServices interface {
	List() ([]db.User, error)
}

type UserServices struct {}

// NewUserServices creates user services
func NewUserServices() IUserServices {
	return &UserServices{}
}

// List 사용자 목록 처리
func (*UserServices) List() ([]db.User, error) {
	arg := db.ListUsersParams{
		Limit:  5,
		Offset: 5,
	}
	return core.GetStore().ListUsers(context.Background(), arg)
}



