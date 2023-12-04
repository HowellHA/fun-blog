package main

import (
	"github.com/HowellHA/fun-blog/router"
	"github.com/HowellHA/fun-blog/service"
)

func main() {
	service.MySQL.Init()
	defer service.MySQL.Close()
	e := router.Init()
	err := e.Run(":8080")
	if err != nil {
		panic(err)
	}
}
