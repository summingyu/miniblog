package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/summingyu/miniblog/internal/pkg/core"
	"github.com/summingyu/miniblog/internal/pkg/errno"
	"github.com/summingyu/miniblog/internal/pkg/log"
	v1 "github.com/summingyu/miniblog/pkg/api/miniblog/v1"
)

func (ctrl *UserController) Create(c *gin.Context) {
	log.C(c).Infow("Create user function called")

	var r v1.CreateUserRequest
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

	if err := ctrl.b.Users().Create(c, &r); err != nil {
		log.C(c).Errorw("Failed to create user", "error", err)
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}
