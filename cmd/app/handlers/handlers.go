package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/igorbelousov/go-web-core/foundation/web"
)

//API function for define routers
func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {

	app := web.NewApp()

	check := check{
		log: log,
	}

	app.Handle(http.MethodGet, "/test", check.readiness)

	return app
}
