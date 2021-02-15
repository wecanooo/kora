package requests

import (
	"context"
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"github.com/wecanooo/kora/core"
	"strconv"
	"strings"
	"unicode/utf8"
)

func init() {
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		tableName := rng[0]
		dbField := field
		if len(rng) == 2 && rng[1] != "" {
			dbField = rng[1]
		}
		val := value.(string)

		if core.GetStore().Exists(context.Background(), tableName, dbField, val) {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v is already exists", val)
		}

		return nil
	})

	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
		if valLength > l {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 의 길이는 %d 를 초과할 수 없습니다", field, l)
		}
		return nil
	})

	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
		if valLength < l {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 의 길이는 %d 보다 커야 합니다", field, l)
		}
		return nil
	})
}
