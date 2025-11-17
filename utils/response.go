package utils

import "github.com/gin-gonic/gin"

//  struktur untuk response
type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// response success
func SuccessResponse(ctx *gin.Context, code int, message string, data interface{}, meta ...interface{}) {
	res := Response{
		Code:    code,
		Status:  "success",
		Message: message,
		Data:    data,
	}

	if len(meta) > 0 {
		res.Meta = meta[0]
	}
	ctx.JSON(code, res)
}

// response error
func ErrorResponse(ctx *gin.Context, code int, message string) {
	res := Response{
		Code:    code,
		Status:  "error",
		Message: message,
	}
	ctx.JSON(code, res)
}
