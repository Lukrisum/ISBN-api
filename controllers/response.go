package controllers

import (
	"ISBN/global"
	"github.com/gin-gonic/gin"
)

var HttpStatus = map[uint]int{
	global.ServerError: 500,
	global.ClientError: 400,
}

var Message = map[uint]string{
	global.ServerError: "服务端错误",
	global.ClientError: "请求错误",
}

func SuccessResponse(ctx *gin.Context, data interface{}, code uint) {

	ctx.JSON(200, map[string]interface{}{
		"code": code,
		"msg":  "success",
		"data": data,
	})

	ctx.Abort()
}

func ErrorResponse(ctx *gin.Context, code uint) {

	httpStatus, isMessageExits := HttpStatus[code]
	if !isMessageExits {
		httpStatus = 403
	}

	message, isMessageExits := Message[code]
	if !isMessageExits {
		message = "未描述的错误"
	}

	ctx.JSON(httpStatus, map[string]interface{}{
		"code": code,
		"msg":  message,
		"data": nil,
	})

	ctx.Abort()
}
