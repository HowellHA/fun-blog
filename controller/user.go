package controller

import (
	"errors"
	"fmt"

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
	if CheckErrWithResult(ctx, err) {
		return
	}

	err = errors.New("123")
	CheckErr(ctx, err)

	err = fmt.Errorf("你的密码%s有错误", param.Password)
	CheckErr(ctx, err)
}
