package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/isbn/:ISBN", handle)

	r.Run("localhost:2222")
}

func handle(c *gin.Context) {

	ISBN := c.Param("ISBN")
	apiKey, ok := c.GetQuery("apikey")

	fmt.Print(ISBN)
	fmt.Print(apiKey)

	// var person Person
	// if err := c.ShouldBindJSON(&person); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	
	if !ok {
		c.String(http.StatusBadRequest, "no apiKey")
		return
	}

	params := url.Values{}
	params.Set("apikey", apiKey)

	Url, _ := url.Parse("https://api.jike.xyz/situ/book/isbn/" + ISBN)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	res, err := http.Get(urlPath)
	defer res.Body.Close()
	if err != nil || res.StatusCode != http.StatusOK  {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
