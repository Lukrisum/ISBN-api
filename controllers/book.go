package controllers

import (
	"ISBN/controllers/utils"
	"github.com/gin-gonic/gin"
)

func GetBookInfo(ctx *gin.Context) {

	ISBN := ctx.Param("ISBN")

	data, code, err := utils.GetBookInfoFromJiKe(ISBN)

	if err != nil {
		ErrorResponse(ctx, code)
		return
	}

	SuccessResponse(ctx, data, code)
}
