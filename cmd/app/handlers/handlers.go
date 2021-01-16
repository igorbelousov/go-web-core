package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/igorbelousov/go-web-core/foundation/web"
	"github.com/igorbelousov/go-web-core/internal/auth"
	"github.com/igorbelousov/go-web-core/internal/mid"
)

//API function for define routers
func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth) *web.App {

	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))

	check := check{
		build: build,
		log:   log,
	}

	app.Handle(http.MethodGet, "/readiness", check.readiness)
	app.Handle(http.MethodGet, "/liveiness", check.liveness)

	return app
}
