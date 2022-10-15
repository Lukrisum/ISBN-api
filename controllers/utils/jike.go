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

type BookInfo struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Subname       string `json:"subname"`
	Author        string `json:"author"`
	Translator    string `json:"translator"`
	Publishing    string `json:"publishing"`
	Published     string `json:"published"`
	Designed      string `json:"designed"`
	Code          string `json:"code"`
	Douban        int64  `json:"douban"`
	DoubanScore   int    `json:"doubanScore"`
	Brand         string `json:"brand"`
	Weight        string `json:"weight"`
	Size          string `json:"size"`
	Pages         string `json:"pages"`
	PhotoUrl      string `json:"photoUrl"`
	LocalPhotoUrl string `json:"localPhotoUrl"`
	Price         string `json:"price"`
	Froms         string `json:"froms"`
	Num           int32  `json:"num"`
	CreateTime    string `json:"createTime"`
	Uptime        string `json:"uptime"`
	AuthorIntro   string `json:"authorIntro"`
	Description   string `json:"description"`
}

func GetBookInfoFromJiKe(ISBN string) (BookInfo, uint, error) {

	params := url.Values{}

	for i, apiKey := range global.ApiKeys {

		if apiKey.LastTime == 0 {
			apiKey.LastTime = time.Now().Unix()
		} else if apiKey.LastTime > 0 {
			if time.Now().Unix()-apiKey.LastTime <= 1 {
				continue
			}
		}

		params.Set("apikey", apiKey.ApiKey)

		Url, urlErr := url.Parse(global.JikeApi + "/situ/book/isbn" + ISBN)

		if urlErr != nil {
			return BookInfo{}, global.ClientError, errors.New("url parse err")
		}

		Url.RawQuery = params.Encode()
		urlPath := Url.String()

		apiKey.LastTime = time.Now().Unix()
		getRes, getErr := http.Get(urlPath)

		defer func(Body io.ReadCloser) {
			closeErr := Body.Close()
			if closeErr != nil {
				log.Println(closeErr)
			}
		}(getRes.Body)

		if getErr != nil || getRes.StatusCode != http.StatusOK {
			if i == 4 {
				// todo: log
				return BookInfo{}, global.ServerError, getErr
			}
			continue
		}

		body, _ := ioutil.ReadAll(getRes.Body)

		jsonRes, ok := gjson.Get(string(body), "data").Value().(BookInfo)

		if !ok {
			return BookInfo{}, global.ServerError, errors.New("Jike api dead")
		}

		return jsonRes, global.Success, nil
	}

	return BookInfo{}, global.Success, nil
}
