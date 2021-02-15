package requests

import "github.com/wecanooo/kora/core/validator"

type CreateUser struct {
	Email    string `valid:"email"`
	Username string `valid:"username"`
	Password string `valid:"password"`
}

func (u *CreateUser) Options() validator.Options {
	return validator.Options{
		Rules: validator.MapData{
			"username": {"required", "max:50"},
			"email":    {"required", "max:255", "email", "not_exists:users"},
			"password": {"required", "min:6"},
		},
		Messages: validator.MapData{
			"username": {
				"required:이름은 반드시 필요합니다",
				"max:이름은 50자 이내로 입력해야 합니다",
			},
			"email": {
				"required:이메일은 반드시 필요합니다",
				"max:이메일은 255자 이내로 입력해야 합니다",
				"email:이메일 형식오류",
				"not_exists:이메일이 이미 등록되었습니다",
			},
			"password": {
				"required:비밀번호는 반드시 필요합니다",
				"min:비밀번호는 최소 6자 이상이어야 합니다",
			},
		},
	}
}
