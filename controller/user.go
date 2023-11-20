package controller

import (
	"github.com/gin-gonic/gin"
)

type userCtrl struct{}

func (userCtrl) Info(ctx *gin.Context)  {}
func (userCtrl) Login(ctx *gin.Context) {}
