// Copyright 2024 summingyu(余苏明) <summingbest@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/summingyu/miniblog.

package user

import (
	"github.com/gin-gonic/gin"

	"github.com/summingyu/miniblog/internal/pkg/core"
	"github.com/summingyu/miniblog/internal/pkg/errno"
	"github.com/summingyu/miniblog/internal/pkg/log"
	v1 "github.com/summingyu/miniblog/pkg/api/miniblog/v1"
)

func (ctrl *UserController) Login(c *gin.Context) {
	log.C(c).Infow("Login function called")

	var r v1.LoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.C(c).Errorw("Failed to bind request", "error", err)
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	resp, err := ctrl.b.Users().Login(c, &r)
	if err != nil {
		log.C(c).Errorw("Failed to login", "error", err)
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}
