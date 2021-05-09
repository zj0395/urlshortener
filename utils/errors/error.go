package errors

type Error struct {
	HttpStatus int    `json:"-"`
	Errno      int    `json:"errno"`
	Msg        string `json:"msg"`
}

func (t *Error) Error() string {
	return t.Msg
}

func (t *Error) String() string {
	return t.Msg
}

func FormatError(err error) *Error {
	if v, ok := err.(*Error); ok {
		return v
	}
	res := *UnknownError
	res.Msg = err.Error()
	return &res
}

var (
	NotFound   = &Error{404, 404, "unsupport api"}
	PanicError = &Error{500, 500, "server error"}

	ParamError   = &Error{200, 1001, "param error"}
	CodeNotExist = &Error{200, 1002, "not exist"}

	UnknownError      = &Error{200, 1100, "unknown error"}
	ErrorFlagValError = &Error{200, 1101, "unknown error"}

	DbError = &Error{200, 1201, "unknown error"}
)
