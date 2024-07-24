package errno

import "fmt"

type Errno struct {
	HTTP    int
	Code    string
	Message string
}

// Errno 实现 error 接口
func (err *Errno) Error() string {
	return err.Message
}

func (err *Errno) SetMessage(format string, args ...interface{}) *Errno {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

func Decode(err error) (hcode int, code string, message string) {
	if err == nil {
		return OK.HTTP, OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Errno:
		return typed.HTTP, typed.Code, typed.Message
	default:
	}
	return InternalServerError.HTTP, InternalServerError.Code, InternalServerError.Message
}
