package main

import (
	"github.com/riqueGo/url_shortener/controller"
	"github.com/riqueGo/url_shortener/repository"
	"github.com/riqueGo/url_shortener/setup"
)

func main() {
	db := setup.SetupPostgresDb()
	repo := repository.NewUrlRepository(db)
	ctrl := controller.NewUrlController(repo)

	r := setup.SetupRouter(ctrl)
	r.Run(":8080")
}
