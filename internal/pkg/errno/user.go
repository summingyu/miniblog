// Copyright 2024 summingyu(余苏明) <summingbest@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/summingyu/miniblog.

package errno

var (
	ErrUserAlreayExist   = &Errno{HTTP: 400, Code: "FailedOperation.UserAlreadyExist", Message: "user already exist."}
	ErrPasswordIncorrect = &Errno{HTTP: 401, Code: "InvalidParameter.PasswordIncorrect", Message: "password incorrect."}
	ErrUserNotFound      = &Errno{HTTP: 404, Code: "ResourceNotFound.UserNotFound", Message: "user not found."}
)
