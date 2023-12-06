package main

import (
	"github.com/HowellHA/fun-blog/router"
)

func main() {
	e := router.Init()
	err := e.Run(":8080")
	if err != nil {
		panic(err)
	}
}
