package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func response(ctx *gin.Context, httpCode int, message string, data any) {
	ctx.JSON(httpCode, gin.H{
		"code":    httpCode,
		"message": message,
		"data":    data,
	})
	return
}

func BadRequest(ctx *gin.Context, message string, data any) {
	response(ctx, http.StatusBadRequest, message, data)
}

func Unauthorized(ctx *gin.Context, message string, data any) {
	response(ctx, http.StatusUnauthorized, message, data)
}

func Forbidden(ctx *gin.Context, message string, data any) {
	response(ctx, http.StatusForbidden, message, data)
}

func NotFound(ctx *gin.Context, message string, data any) {
	response(ctx, http.StatusNotFound, message, data)
}

func MethodNotAllowed(ctx *gin.Context, message string, data any) {
	response(ctx, http.StatusMethodNotAllowed, message, data)
}

func InternalServerError(ctx *gin.Context, message string, data any) {
	response(ctx, http.StatusInternalServerError, message, data)
}

func Ok(ctx *gin.Context, data any) {
	response(ctx, http.StatusOK, "success", data)
}

func ParameterError(ctx *gin.Context, message string, data any) {
	response(ctx, http.StatusBadRequest, message, data)
}
