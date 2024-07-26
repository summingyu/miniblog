// Copyright 2024 summingyu(余苏明) <summingbest@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/summingyu/miniblog.

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
