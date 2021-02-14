package api

import (
	db "github.com/wecanooo/kora/app/models"
	"github.com/wecanooo/kora/app/serializers"
	"github.com/wecanooo/kora/app/services"
	"github.com/wecanooo/kora/core/context"
	"net/http"
)

type IUserController interface {
	Index(*context.AppContext) error
	Show(*context.AppContext, *db.User) error
}

type UserController struct {
	UserServices services.IUserServices
}

// NewUserController user controller constructor
func NewUserController(s services.IUserServices) IUserController {
	return &UserController{
		UserServices: s,
	}
}

// Index 사용자 목록 반환
func (u *UserController) Index(ctx *context.AppContext) error {
	users, err := u.UserServices.List()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, serializers.NewUserListSerializer(users))
}

// Show 사용자 정보 반환
func (u *UserController) Show(ctx *context.AppContext, user *db.User) error {
	return nil
}