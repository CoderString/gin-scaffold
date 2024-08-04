package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app, f, err := initApp()
	if err != nil {
		panic(err)
	}
	defer f()

	g := gin.Default()
	g.POST("/add", app.UserSrv.Add())
	g.Run(":8090")
}
