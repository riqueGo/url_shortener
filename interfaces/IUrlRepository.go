package interfaces

import "github.com/riqueGo/url_shortener/domain"

type IUrlRepository interface {
	SaveUrl(urlDomain *domain.UrlDomain) error
	GetUrl(code string) (string, error)
}
