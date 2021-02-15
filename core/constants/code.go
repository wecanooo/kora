package constants

type LogicCode int

const (
	Success       LogicCode = 0
	UnknownError  LogicCode = -1
	RequestError  LogicCode = 100
	ResourceError LogicCode = 101
	DatabaseError LogicCode = 102
	TokenError    LogicCode = 103
	NotFoundError LogicCode = 104
	AuthError     LogicCode = 105
)
