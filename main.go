package main

import "github.com/HowellHA/fun-blog/router"

func main() {
	e := router.Init()
	e.Run(":8080")
}
