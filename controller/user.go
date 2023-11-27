package controller

import (
	"github.com/gin-gonic/gin"
)

type userCtrl struct{}

func (userCtrl) Info(ctx *gin.Context) {

}

type UserLoginParam struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

func (userCtrl) Login(ctx *gin.Context) {
	param := UserLoginParam{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

}
