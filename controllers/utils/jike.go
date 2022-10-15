package utils

import (
	"ISBN/global"
	"errors"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func GetBookInfoFromJiKe(ISBN string) (interface{}, int, error) {

	params := url.Values{}

	for i := 0; i < len(global.ApiKeys); i++ {

		if global.ApiKeys[i].LastTime > 0 && time.Now().Unix()-global.ApiKeys[i].LastTime <= 1 {
			if i == 4 {
				break
			}
			continue
		}

		params.Set("apikey", global.ApiKeys[i].ApiKey)

		Url, urlErr := url.Parse(global.JikeApi + "/situ/book/isbn/" + ISBN)

		if urlErr != nil {
			return nil, http.StatusBadRequest, errors.New("url parse err")
		}

		Url.RawQuery = params.Encode()
		urlPath := Url.String()

		global.ApiKeys[i].LastTime = time.Now().Unix()

		getRes, getErr := http.Get(urlPath)

		defer func(Body io.ReadCloser) {
			closeErr := Body.Close()
			if closeErr != nil {
				log.Println(closeErr)
			}
		}(getRes.Body)

		if getErr != nil || getRes.StatusCode != http.StatusOK {
			// todo: log
			return nil, getRes.StatusCode, getErr
		}

		body, _ := ioutil.ReadAll(getRes.Body)

		jsonRes := gjson.Get(string(body), "data").Value()

		if jsonRes == nil {
			return jsonRes, http.StatusNotFound, errors.New("ISBN not found")
		}

		return jsonRes, http.StatusOK, nil
	}

	return nil, http.StatusTooManyRequests, errors.New("please retry later")
}
