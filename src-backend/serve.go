package main

import (
	"fmt"

	"spa-web-app-template/src-backend/api"
	"spa-web-app-template/src-backend/libs"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	libs.DBInit()
	r := gin.Default()
	config := cors.DefaultConfig()
	r.Use(cors.New(config))
	r.POST("/test", test)
	apiGroup := r.Group("/api")
	apiGroup.POST("/insertUser", api.InsertUser)
	r.Run(":8080")
}

func root(ctx *gin.Context) {
	fmt.Println("hello")
	ctx.Writer.WriteString("root")
}

func test(ctx *gin.Context) {
	fmt.Println("hello3")
	ctx.Writer.WriteString("hello3")
}
