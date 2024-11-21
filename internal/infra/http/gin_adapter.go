package httpserver

import (
	"github.com/gin-gonic/gin"
)

type GinHTTPServerAdapter struct {
	Engine   *gin.Engine
	basePath string
}

func NewGinHTTPServerAdapter() *GinHTTPServerAdapter {
	return &GinHTTPServerAdapter{
		Engine: gin.Default(),
	}
}

func (g *GinHTTPServerAdapter) SetTrustedProxies(trustedProxies []string) error {
	return g.Engine.SetTrustedProxies(trustedProxies)
}

func (g *GinHTTPServerAdapter) SetBasePath(basePath string) {
	g.basePath = basePath
}

func (g *GinHTTPServerAdapter) Run(adds ...string) error {
	return g.Engine.Run(adds...)
}

func (g *GinHTTPServerAdapter) SetSwagger(path string, callback gin.HandlerFunc) {
	g.Engine.GET(g.basePath+path, callback)
}

func (g *GinHTTPServerAdapter) GET(path string, callback func(HTTPContext)) {
	g.Engine.GET(
		g.basePath+path,
		func(c *gin.Context) {
			callback(c)
		},
	)
}

func (g *GinHTTPServerAdapter) POST(path string, callback func(HTTPContext)) {
	g.Engine.POST(
		g.basePath+path,
		func(c *gin.Context) {
			callback(c)
		},
	)
}

func (g *GinHTTPServerAdapter) PUT(path string, callback func(HTTPContext)) {
	g.Engine.PUT(
		g.basePath+path,
		func(c *gin.Context) {
			callback(c)
		},
	)
}

func (g *GinHTTPServerAdapter) PATCH(path string, callback func(HTTPContext)) {
	g.Engine.PATCH(
		g.basePath+path,
		func(c *gin.Context) {
			callback(c)
		},
	)
}

func (g *GinHTTPServerAdapter) DELETE(path string, callback func(HTTPContext)) {
	g.Engine.DELETE(
		g.basePath+path,
		func(c *gin.Context) {
			callback(c)
		},
	)
}
