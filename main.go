package main

import (
	"github.com/HowellHA/fun-blog/router"
	"github.com/HowellHA/fun-blog/service"
)

func main() {
	db := service.MySQLInit()
	db.Ping()
	e := router.Init()
	e.Run(":8080")
}
