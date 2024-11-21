package httpserver

type HTTPServer interface {
	SetTrustedProxies(trustedProxies []string) error
	Run(addr ...string) error
}

type HTTPRoutes interface {
	GET(string, HTTPHandlerFunc) HTTPRoutes
	POST(string, HTTPHandlerFunc) HTTPRoutes
	DELETE(string, HTTPHandlerFunc) HTTPRoutes
	PATCH(string, HTTPHandlerFunc) HTTPRoutes
	PUT(string, HTTPHandlerFunc) HTTPRoutes
	SetBasePath(basePath string) HTTPRoutes
}

type HTTPContext interface {
	Header(key, value string)
	JSON(code int, obj any)
	BindJSON(obj any) error
	Param(key string) string
	DefaultQuery(key, defaultValue string) string
}

type HTTPHandlerFunc func(HTTPContext)

type Payload map[string]any
