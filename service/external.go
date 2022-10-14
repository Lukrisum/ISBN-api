package service

import (
	"ISBN/global"
	"ISBN/models/dao"
	"net/http"
	"net/url"
)

func GetBookInfoFromJike(ISBN string) {

	var (
		err     error
		errCode uint
		data    interface{}
	)

	params := url.Values{}
	params.Set("apikey", apiKey)

	Url, err := url.Parse(global.JikeApi + "/situ/book/isbn" + ISBN)
	if err != nil {
		errCode = global.ServerError
		return
	}

	Url.RawQuery = params.Encode()
	urlPath := Url.String()

	res, err := http.Get(urlPath)
	defer res.Body.Close()

	if err != nil || res.StatusCode != http.StatusOK {
		
		return data, errCode, err
	}

	// body, _ := ioutil.ReadAll(res.Body)
}
