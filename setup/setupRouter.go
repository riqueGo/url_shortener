package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/riqueGo/url_shortener/controller"
)

func SetupRouter(ctrl *controller.UrlController) *gin.Engine {
	r := gin.Default()

	r.POST("/", ctrl.ShortenURL)
	r.GET("/:code", ctrl.GetUrl)
	r.GET("/", ctrl.GetAllUrls)

	return r
}
