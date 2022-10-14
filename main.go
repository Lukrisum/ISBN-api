package main

import (
	// "github.com/gin-gonic/gin"
	// "ISBN/controllers"
	"ISBN/global"
	"fmt"
)

func main() {
	// r := gin.Default()

	for _,apiKey := range global.ApiKeys{
		fmt.Println(apiKey.ApiKey)
	}

	// r.GET("/api/isbn/:ISBN", controllers.GetBookInfo)

	// r.Run("localhost:23333")

}
