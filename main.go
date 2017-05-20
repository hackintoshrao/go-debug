package debug

import (
	"github.com/gorilla/mux"
	"net/http"
)

import _ "net/http/pprof"

func RegisterDebugInfo(router *mux.Router) {
	router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
}
