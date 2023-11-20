package router

import (
	"github.com/HowellHA/fun-blog/controller"
	"github.com/gin-gonic/gin"
)

type userRoute struct{}

var _ iRoute = userRoute{}
var user userRoute

func (userRoute) Register(engine *gin.Engine) {
	ctrl := controller.User
	g := engine.Group(
		"/user",
	)

	g.POST(
		"/login",
		ctrl.Login,
	)

	g.GET(
		"/info",
		ctrl.Info,
	)

}
