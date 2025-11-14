package main

import (
	"github.com/AdeleyeShina/go-crud/initializer"
	"github.com/AdeleyeShina/go-crud/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv_variables()
	initializer.ConnectDB()
}

func main() {
	r := gin.Default()
	routes.UserRoute(r)
	r.Run()
}
