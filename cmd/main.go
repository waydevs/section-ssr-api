package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waydevs/section-ssr-api/cmd/api/routes"
	"github.com/waydevs/section-ssr-api/internal/platform/configs"
)

func main() {
	var config = &configs.AppConfig{
		Port: ":8080",
		Env:  "development",
	}

	routes.NewRouter(gin.Default()).Build(config)

}
