package controllers

import (
	"ISBN/controllers/utils"
	"ISBN/global"
	"github.com/gin-gonic/gin"
)

func GetBookInfo(ctx *gin.Context) {

	ISBN := ctx.Param("ISBN")

	data, code, err := utils.GetBookInfoFromJiKe(ISBN)

	if err != nil {

		resCode := global.ClientError
		if code >= 500 {
			resCode = global.ServerError
		}
		ctx.JSON(code, map[string]interface{}{
			"code": resCode,
			"msg":  err.Error(),
			"data": nil,
		})

		ctx.Abort()
		return
	}

	ctx.JSON(code, map[string]interface{}{
		"code": global.Success,
		"msg":  "success",
		"data": data,
	})

	ctx.Abort()
}
