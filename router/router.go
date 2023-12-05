package router

import (
	"github.com/HowellHA/fun-blog/middleware"
	"github.com/gin-gonic/gin"
)

type iRoute interface {
	Register(*gin.Engine)
}

var iRoutes = []iRoute{
	user,
}

func Init() *gin.Engine {
	engine := gin.New()
	engine.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.GlobalErrorHandle,
	)

	for _, r := range iRoutes {
		r.Register(engine)
	}

	return engine
}
