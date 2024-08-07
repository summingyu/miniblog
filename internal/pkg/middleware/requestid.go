// Copyright 2024 summingyu(余苏明) <summingbest@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/summingyu/miniblog.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/summingyu/miniblog/internal/pkg/known"
)

// RequestID 是一个Gin中间件, 用来在每一个HTTP请求的context, response 中注入`X-Request-ID` 键值对
func RequestID() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取请求的ID
		requestID := c.Request.Header.Get(known.XRequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		// 将请求的ID注入到上下文中
		c.Set(known.XRequestIDKey, requestID)
		// 将RequestID写入到响应头中
		c.Writer.Header().Set(known.XRequestIDKey, requestID)
		// 调用下一个中间件
		c.Next()
	}
}
