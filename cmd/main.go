package main

import (
	"github.com/k8-proxy/k8-go-api/controllers"
	"github.com/k8-proxy/k8-go-api/router"
	"github.com/subosito/gotenv"
)

var httpRouter router.Router = router.NewMuxRouter()

func init() {
	gotenv.Load()
}

func main() {
	const port string = ":8100"

	httpRouter.POST("/api/rebuild/file", controllers.RebuildFile)
	httpRouter.POST("/api/rebuild/zip", controllers.Rebuildzip)
	httpRouter.POST("/api/rebuild/base64", controllers.RebuildBase64)

	httpRouter.SERVE(port)
}
