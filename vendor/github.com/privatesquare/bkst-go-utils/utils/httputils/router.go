package httputils

import (
	"github.com/gin-gonic/gin"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"net/http"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logger.GinZap())
	r.Use(gin.Recovery())
	r.NoRoute(NoRoute)
	r.HandleMethodNotAllowed = true
	r.NoMethod(MethodNotAllowed)
	return r
}

func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, RestMsg{Message: http.StatusText(http.StatusOK)})
}

// NoRoute no route controller handles request on endpoints that are not configured
func NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, RestMsg{Message: "Path Not Found"})
}

// MethodNotAllowed method not allowed controller handles request on known endpoints but on methods that are not configured
func MethodNotAllowed(ctx *gin.Context) {
	ctx.JSON(http.StatusMethodNotAllowed, RestMsg{Message: http.StatusText(http.StatusMethodNotAllowed)})
}
