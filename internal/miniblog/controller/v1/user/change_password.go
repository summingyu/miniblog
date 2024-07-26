// Copyright 2024 summingyu(余苏明) <summingbest@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/summingyu/miniblog.

package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/summingyu/miniblog/internal/pkg/core"
	"github.com/summingyu/miniblog/internal/pkg/errno"
	"github.com/summingyu/miniblog/internal/pkg/log"
	v1 "github.com/summingyu/miniblog/pkg/api/miniblog/v1"
)

func (ctrl *UserController) ChangePassword(c *gin.Context) {
	log.C(c).Infow("Change password function called")

	var r v1.ChangePasswordRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.C(c).Errorw("Failed to bind request", "error", err)
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		log.C(c).Errorw("Failed to validate request", "error", err)
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := ctrl.b.Users().ChangePassword(c, c.Param("name"), &r); err != nil {
		log.C(c).Errorw("Failed to change password", "error", err)
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}
