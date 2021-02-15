package errno

import (
	"github.com/wecanooo/kora/core/constants"
	"net/http"
)

var (
	UnknownErr = &Errno{HTTPCode: http.StatusOK, Code: constants.UnknownError, Message: "unknown error"}
	ReqErr = &Errno{HTTPCode: http.StatusOK, Code: constants.RequestError, Message: "request error"}
	ResourceErr = &Errno{HTTPCode: http.StatusOK, Code: constants.ResourceError, Message: "resource error"}
	DatabaseErr = &Errno{HTTPCode: http.StatusOK, Code: constants.DatabaseError, Message: "database error"}
	TokenErr = &Errno{HTTPCode: http.StatusOK, Code: constants.TokenError, Message: "token error"}
	NotFoundErr = &Errno{HTTPCode: http.StatusOK, Code: constants.NotFoundError, Message: "route not found"}
	AuthErr = &Errno{HTTPCode: http.StatusOK, Code: constants.AuthError, Message: "auth error"}
)
