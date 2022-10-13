package main

import (
	"github.com/gin-gonic/gin"
	"ISBN/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/api/isbn/:ISBN", controllers.GetBookInfo)

	r.Run("localhost:23333")
}
