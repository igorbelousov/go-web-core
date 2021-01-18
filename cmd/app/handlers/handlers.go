package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/igorbelousov/go-web-core/foundation/web"
	"github.com/igorbelousov/go-web-core/internal/auth"
	"github.com/igorbelousov/go-web-core/internal/data/user"
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

	ug := userGroup{
		user: user.New(log, db),
	}
	app.Handle(http.MethodGet, "/v1/users/:page/:rows", ug.query, mid.Authenticate(a), mid.Authorize(auth.RoleAdmin))
	app.Handle(http.MethodGet, "/v1/users/token/:kid", ug.token)
	app.Handle(http.MethodGet, "/v1/users/:id", ug.queryByID, mid.Authenticate(a))
	app.Handle(http.MethodPost, "/v1/users", ug.create, mid.Authenticate(a), mid.Authorize(auth.RoleAdmin))
	app.Handle(http.MethodPut, "/v1/users/:id", ug.update, mid.Authenticate(a), mid.Authorize(auth.RoleAdmin))
	app.Handle(http.MethodDelete, "/v1/users/:id", ug.delete, mid.Authenticate(a), mid.Authorize(auth.RoleAdmin))

	return app
}
