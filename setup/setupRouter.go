package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/riqueGo/url_shortner/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/shorten", controller.ShortenURL)
	r.GET("/url/:code", controller.GetUrl)

	return r
}
