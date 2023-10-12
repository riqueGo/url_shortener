package controller

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strings"
)

var urlMap = make(map[string]string)

func ShortenURL(c *gin.Context) {
	var requestData struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code := generateUniqueCode()

	urlMap[code] = requestData.URL

	c.JSON(http.StatusOK, gin.H{"code": code})
}

func GetUrl(c *gin.Context) {
	code := c.Param("code")

	url, exists := urlMap[code]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}

func generateUniqueCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const codeLenght = 6

	generator := rand.New(rand.NewSource(1))

	for {
		var code strings.Builder
		for i := 0; i < codeLenght; i++ {
			code.WriteByte(charset[generator.Intn(len(charset))])
		}
		if _, exists := urlMap[code.String()]; !exists {
			return code.String()
		}
	}
}
