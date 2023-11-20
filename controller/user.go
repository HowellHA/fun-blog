package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userCtrl struct{}

func (userCtrl) Info(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"name":  "fun",
		"email": "438676280@qq.com",
	})
}
func (userCtrl) Login(ctx *gin.Context) {}
