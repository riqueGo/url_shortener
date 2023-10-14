package domain

import (
	"math/rand"
	"strings"
	"time"
)

type UrlDomain struct {
	Code string
	URL  string
}

func (u *UrlDomain) GenerateUniqueCode() {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const codeLenght = 6

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	var code strings.Builder
	for i := 0; i < codeLenght; i++ {
		code.WriteByte(charset[generator.Intn(len(charset))])
	}

	u.Code = code.String()
}
