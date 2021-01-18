package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/igorbelousov/go-web-core/foundation/web"
	"github.com/igorbelousov/go-web-core/internal/auth"
	"github.com/igorbelousov/go-web-core/internal/mid"
	"github.com/jmoiron/sqlx"
)

//API function for define routers
func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth, db *sqlx.DB) *web.App {

	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))

	cg := checkGroup{
		build: build,
		db:    db,
	}

	app.Handle(http.MethodGet, "/readiness", cg.readiness)
	app.Handle(http.MethodGet, "/liveiness", cg.liveness)

	return app
}
