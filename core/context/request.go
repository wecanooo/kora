package context

import (
	"github.com/wecanooo/kora/core/errno"
	"github.com/wecanooo/kora/core/validator"
)

func (c *AppContext) BindValidatorStruct(v validator.Validator) error {
	if err := c.Bind(v); err != nil {
		return err
	}

	if errs, ok := validator.ValidateStruct(v); !ok {
		return errno.ReqErr.WithData(errs)
	}

	return nil
}
