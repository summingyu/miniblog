package miniblog

import (
	"github.com/gin-gonic/gin"

	"github.com/summingyu/miniblog/internal/miniblog/controller/v1/user"
	"github.com/summingyu/miniblog/internal/miniblog/store"
	"github.com/summingyu/miniblog/internal/pkg/core"
	"github.com/summingyu/miniblog/internal/pkg/errno"
	"github.com/summingyu/miniblog/internal/pkg/log"
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

	// 创建v1 路由分组
	v1 := g.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
		}
	}

	return nil
}
