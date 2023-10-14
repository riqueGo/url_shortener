package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/riqueGo/url_shortener/controller"
)

func SetupRouter(ctrl *controller.UrlController) *gin.Engine {
	r := gin.Default()

	r.POST("/shorten", ctrl.ShortenURL)
	r.GET("/url/:code", ctrl.GetUrl)

	return r
}
