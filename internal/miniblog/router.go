// Copyright 2024 summingyu(余苏明) <summingbest@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/summingyu/miniblog.

package miniblog

import (
	"github.com/gin-gonic/gin"

	"github.com/summingyu/miniblog/internal/miniblog/controller/v1/user"
	"github.com/summingyu/miniblog/internal/miniblog/store"
	"github.com/summingyu/miniblog/internal/pkg/core"
	"github.com/summingyu/miniblog/internal/pkg/errno"
	"github.com/summingyu/miniblog/internal/pkg/log"
	mw "github.com/summingyu/miniblog/internal/pkg/middleware"
)

// installRouters 安装miniblog接口路由
func installRouters(g *gin.Engine) error {
	// 注册404 Handler
	g.NoRoute(func(c *gin.Context) {
		log.C(c).Debugw("404 page not found", "path", c.Request.URL)
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	uc := user.New(store.S)

	g.POST("/login", uc.Login)

	// 创建v1 路由分组
	v1 := g.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
			userv1.Use(mw.Authn())
		}
	}

	return nil
}
