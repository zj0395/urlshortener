package errors

import (
	"github.com/zj0395/golib/liberr"
)

var (
	CodeNotExist = &liberr.Error{200, 1002, "not exist"}
)
