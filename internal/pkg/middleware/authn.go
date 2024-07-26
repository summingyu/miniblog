// Copyright 2024 summingyu(余苏明) <summingbest@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/summingyu/miniblog.

package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/summingyu/miniblog/internal/pkg/core"
	"github.com/summingyu/miniblog/internal/pkg/errno"
	"github.com/summingyu/miniblog/internal/pkg/known"
	"github.com/summingyu/miniblog/pkg/token"
)

// RequestID 是一个Gin中间件, 用来在每一个HTTP请求的context, response 中注入`X-Request-ID` 键值对
func Authn() func(c *gin.Context) {
	return func(c *gin.Context) {
		username, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set(known.XUsernameKey, username)
		// 调用下一个中间件
		c.Next()
	}
}
