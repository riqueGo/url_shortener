package main

import (
	"github.com/riqueGo/url_shortner/setup"
)

func main() {
	r := setup.SetupRouter()
	r.Run(":8080")
}
