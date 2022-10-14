package service

import (
	"ISBN/global"
	"ISBN/models/dao"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func handleApiKeys() {
	for _, apiKey := range global.ApiKeys {
		fmt.Print(apiKey)
	}
}

func GetBookInfoFromJike(ISBN string) {

	var (
		serverErr     error
		errCode uint
		data    dao.BookInfo
	)

	params := url.Values{}

	for i, apiKey := range global.ApiKeys {
		params.Set("apikey", apiKey.ApiKey)

		Url, err := url.Parse(global.JikeApi + "/situ/book/isbn" + ISBN)

		if err != nil {

		}

		Url.RawQuery = params.Encode()
		urlPath := Url.String()

		res, serverErr := http.Get(urlPath)
		body, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		if serverErr != nil || res.StatusCode != http.StatusOK {
			if i == 4 {
				errCode = global.ServerError
				return , errCode, serverErr
			}
		}

	}

	// body, _ := ioutil.ReadAll(res.Body)
}
