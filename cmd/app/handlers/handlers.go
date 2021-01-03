package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/dimfeld/httptreemux/v5"
)

//API function for define routers
func API(build string, shutdown chan os.Signal, log *log.Logger) *httptreemux.ContextMux {

	tm := httptreemux.NewContextMux()

	check := check{
		log: log,
	}

	tm.Handle(http.MethodGet, "/test", check.readiness)

	return tm
}
