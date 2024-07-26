// Copyright 2024 summingyu(余苏明) <summingbest@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/summingyu/miniblog.

package core

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/summingyu/miniblog/internal/pkg/errno"
)

type ErrResponse struct {
	// Code 指定了业务错误码.
	Code string `json:"code"`
	// Message 指定了可以直接对外展示的错误信息.
	Message string `json:"message"`
}

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		hcode, code, message := errno.Decode(err)
		c.JSON(hcode, ErrResponse{
			Code:    code,
			Message: message,
		})

		return
	}

	c.JSON(http.StatusOK, data)
}
