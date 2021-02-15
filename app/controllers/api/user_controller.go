package api

import (
	db "github.com/wecanooo/kora/app/models"
	"github.com/wecanooo/kora/app/requests"
	"github.com/wecanooo/kora/app/serializers"
	"github.com/wecanooo/kora/app/services"
	"github.com/wecanooo/kora/core/context"
	"github.com/wecanooo/kora/core/errno"
	"github.com/wecanooo/kora/core/pkg/password"
)

// IUserController defines a user controller interface
type IUserController interface {
	Index(*context.AppContext) error
	Show(*context.AppContext, *db.User) error
	Create(appContext *context.AppContext) error
}

// UserController defines a user controller struct
type UserController struct {
	UserServices services.IUserServices
}

// NewUserController creates a user controller instance
func NewUserController(s services.IUserServices) IUserController {
	return &UserController{
		UserServices: s,
	}
}

// Index returns user list
func (u *UserController) Index(ctx *context.AppContext) error {
	users, err := u.UserServices.List()
	if err != nil {
		return errno.DatabaseErr.WithErr(err)
	}
	return ctx.SuccessJSON(serializers.NewUserListSerializer(users))
}

// Show returns a user
func (u *UserController) Show(ctx *context.AppContext, user *db.User) error {
	return nil
}

// Create creates a user
func (u *UserController) Create(ctx *context.AppContext) error {
	req := new(requests.CreateUser)
	if err := ctx.BindValidatorStruct(req); err != nil {
		return err
	}

	hashedPassword, err := password.Encrypt(req.Password)
	if err != nil {
		return err
	}

	user, err := u.UserServices.Create(db.CreateUserParams{
		Email:             req.Email,
		Username:          req.Username,
		EncryptedPassword: hashedPassword,
	})

	if err != nil {
		errno.DatabaseErr.WithErr(err)
	}

	return ctx.SuccessJSON(serializers.NewUserSerializer(user))
}