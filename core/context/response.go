package context

import (
	"github.com/wecanooo/kora/core/constants"
	"github.com/wecanooo/kora/core/errno"
	"net/http"
)

// CommonResponse defines a common response data struct
type CommonResponse struct {
	Code    constants.LogicCode `json:"code"`
	Message string              `json:"msg"`
	Data    interface{}         `json:"data,omitempty"`
}

// NewCommonResponse creates a common response instance
func NewCommonResponse(code constants.LogicCode, message string, data interface{}) *CommonResponse {
	return &CommonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewSuccessResponse creates a common response instance with success message
func NewSuccessResponse(message string, data interface{}) *CommonResponse {
	return NewCommonResponse(constants.Success, message, data)
}

// NewErrorResponse creates a common response instance with error message
func NewErrorResponse(e *errno.Errno) *CommonResponse {
	return NewCommonResponse(e.Code, e.Message, e.Data)
}

// SuccessJSON returns a success json data
func (c *AppContext) SuccessJSON(data interface{}) error {
	return c.JSON(http.StatusOK, NewSuccessResponse("success", data))
}

// ErrorJSON returns a error json data
func (c *AppContext) ErrorJSON(e *errno.Errno) error {
	return c.JSON(e.HTTPCode, NewErrorResponse(e))
}
