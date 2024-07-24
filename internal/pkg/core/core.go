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
