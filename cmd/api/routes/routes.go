package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/waydevs/section-ssr-api/internal/platform/configs"
)

type Router struct {
	eng *gin.Engine
}

func NewRouter(eng *gin.Engine) *Router {
	return &Router{eng: eng}
}

type Middlewares map[string]gin.HandlerFunc

// Build router
func (r *Router) Build(config *configs.AppConfig) {
	r.buildRoutes(Middlewares{})
	r.eng.Run(config.Port)
}

// buildRoutes define routes with middlewares
func (r *Router) buildRoutes(middlewares Middlewares) {
	r.eng.GET("/home", func(c *gin.Context) { c.String(200, "Nothing here") })

}
