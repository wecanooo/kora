package errno

import (
	"fmt"
	"github.com/wecanooo/kora/core/constants"
)

type Errno struct {
	HTTPCode int
	Message  string
	Code     constants.LogicCode
	Err      error
	Data     interface{}
}

// Error returns a error message
func (e Errno) Error() string {
	return e.Message
}

// WithErr returns a errno instance from error
func (e Errno) WithErr(err error) error {
	if err == nil {
		return nil
	}

	return &Errno{
		HTTPCode: e.HTTPCode,
		Code: e.Code,
		Message: err.Error(),
		Err: err,
	}
}

// WithData returns a errno instance from data
func (e Errno) WithData(d interface{}) error {
	return &Errno{
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Data:     d,
	}
}

// WithMessage returns a errno instance from message string
func (e Errno) WithMessage(msg string) error {
	return &Errno{
		HTTPCode: e.HTTPCode,
		Message:  msg,
		Code:     e.Code,
		Err:      e.Err,
	}
}

// WithMessagef returns a errno instance from formatted arguments
func (e Errno) WithMessagef(format string, args ...interface{}) error {
	return &Errno{
		HTTPCode: e.HTTPCode,
		Message:  fmt.Sprintf(format, args...),
		Err:      e.Err,
	}
}

// WithErrMessage returns a errno instance from error and message
func (e Errno) WithErrMessage(err error, msg string) error {
	if err == nil {
		return nil
	}

	return &Errno{
		HTTPCode: e.HTTPCode,
		Message:  msg,
		Code:     e.Code,
		Err:      err,
	}
}

// WithErrMessagef returns a errno instance from err and formatted arguments
func (e Errno) WithErrMessagef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return &Errno{
		HTTPCode: e.HTTPCode,
		Message:  fmt.Sprintf(format, args...),
		Code:     e.Code,
		Err:      err,
	}
}