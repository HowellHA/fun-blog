package controller

import "github.com/gin-gonic/gin"

var User = userCtrl{}

func CheckErr(ctx *gin.Context, err error) {
	if err != nil {
		_ = ctx.Error(err)
	}
}

func CheckErrWithResult(ctx *gin.Context, err error) bool {
	if err != nil {
		_ = ctx.Error(err)
		return true
	}

	return false
}
