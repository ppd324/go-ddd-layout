package router

import "github.com/gin-gonic/gin"

type Router interface {
	Mount(gin *gin.Engine)
	Name() string
}
