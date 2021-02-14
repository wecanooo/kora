package serializers

import db "github.com/wecanooo/kora/app/models"

type (
	User struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}
	UserList []User
)

// NewUserSerializer 사용자 정보 serializer
func NewUserSerializer(m db.User) User {
	return User{
		Email:    m.Email,
		Username: m.Username,
	}
}

// NewUserListSerializer 사용자 목록 serializer
func NewUserListSerializer(ms []db.User) UserList {
	us := make(UserList, 0)
	for _, m := range ms {
		us = append(us, NewUserSerializer(m))
	}
	return us
}
