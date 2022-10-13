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
	global.ClientError: "请求数据错误",
}

func SuccessResponse(ctx *gin.Context, data interface{}, extraDescription string) {

	message := "success"
	if extraDescription != "" {
		message = extraDescription
	}

	ctx.JSON(200, map[string]interface{}{
		"code": 0,
		"msg":  message,
		"data": data,
	})

	ctx.Abort()
}

func ErrorResponse(ctx *gin.Context, code uint, extraDescription string) {

	httpStatus, isMessageExits := HttpStatus[code]
	if !isMessageExits {
		httpStatus = 403
	}

	message, isMessageExits := Message[code]
	if !isMessageExits {
		message = "未描述的错误" + extraDescription
	}

	ctx.JSON(httpStatus, map[string]interface{}{
		"code": code,
		"msg":  message,
		"data": nil,
	})

	ctx.Abort()
}
