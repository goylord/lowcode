package main

import (
	"github.com/gin-gonic/gin"
	routes "github.com/goylold/lowcode/routes"
)

func main() {
	router := gin.Default()
	routes.RouterRegister(router)
	router.Run()
}
