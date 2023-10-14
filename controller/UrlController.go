package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/riqueGo/url_shortener/contants"
	"github.com/riqueGo/url_shortener/domain"
	"github.com/riqueGo/url_shortener/interfaces"
	"log"
	"net/http"
)

var urlMap = make(map[string]string)

type UrlController struct {
	repository interfaces.IUrlRepository
}

func NewUrlController(repository interfaces.IUrlRepository) *UrlController {
	return &UrlController{repository: repository}
}

func (ctrl UrlController) ShortenURL(c *gin.Context) {
	var urlDomain domain.UrlDomain

	if err := c.BindJSON(&urlDomain); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.repository.SaveUrl(&urlDomain)
	if err != nil {
		log.Println("Error: ", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": urlDomain.Code})
}

func (ctrl UrlController) GetUrl(c *gin.Context) {
	code := c.Param("code")

	url, err := ctrl.repository.GetUrl(code)

	if err != nil {
		if errors.Is(err, contants.ErrCodeNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, url)
}

func (ctrl UrlController) GetAllUrls(c *gin.Context) {
	urls, err := ctrl.repository.GetAllUrls()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"urls": urls})
}
