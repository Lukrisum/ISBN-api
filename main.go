package main

import (
	"ISBN/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/isbn/:ISBN", controllers.GetBookInfo)

	r.Run("localhost:23333")

}
