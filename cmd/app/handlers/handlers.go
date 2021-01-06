package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/igorbelousov/go-web-core/foundation/web"
	"github.com/igorbelousov/go-web-core/internal/mid"
)

//API function for define routers
func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {

	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Panics(log))

	check := check{
		log: log,
	}

	app.Handle(http.MethodGet, "/readiness", check.readiness)

	return app
}
