package controllers

import (
	"ISBN/global"
	"ISBN/models/dto"
	"ISBN/service"
	"github.com/gin-gonic/gin"
)

func GetBookInfo(ctx *gin.Context) {

	var getBookInfo dto.GetBookInfo

	ISBN := ctx.Param("ISBN")
	apiKey, ok := ctx.GetQuery("apikey")

	if !ok {
		ErrorResponse(ctx, global.ClientError)
	}

	getBookInfo.ISBN = ISBN
	getBookInfo.ApiKey = apiKey

	var (
		err     error
		data    interface{}
		errCode 		uint
	)

	data, errCode, err = service.GetBookInfoFromJike(getBookInfo.ISBN)

	if err != nil {
		ErrorResponse(ctx,errCode,)
		return
	}

	SuccessResponse(ctx,data)
}