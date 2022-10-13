package service

import (
	"ISBN/global"
	"net/http"
	"net/url"
)

func GetBookInfoFromJike(ISBN, apiKey string) {

	var (
		err     error
		errCode uint
		// data interface{}
		message string
	)

	params := url.Values{}
	params.Set("apikey", apiKey)

	Url, err := url.Parse(global.JikeApi + "/situ/book/isbn" + ISBN)
	if err != nil {
		errCode = global.ServerError
		message = "url 解析错误，请检查传参"
		return
	}

	Url.RawQuery = params.Encode()
	urlPath := Url.String()

	res, err := http.Get(urlPath)
	defer res.Body.Close()

	if err != nil || res.StatusCode != http.StatusOK {

		return
	}

	// body, _ := ioutil.ReadAll(res.Body)
}
