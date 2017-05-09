package utils

import (
	"errors"
	"fmt"
)

func NewError(msg string, cause error) error {
	return errors.New(fmt.Sprintf("%s << %s", msg, cause.Error()))
}
