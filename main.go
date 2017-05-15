package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"runtime"
	"runtime/debug"
	rtpprof "runtime/pprof"
	"strconv"
)

import _ "net/http/pprof"

//import gomhttp "github.com/rakyll/gom/http"

func AttachProfiler(router *mux.Router) {
	/*
		router.PathPrefix("/debug/pprof/profile").HandlerFunc(pprof.Profile)
		//	router.PathPrefix("/debug/pprof/heap").HandlerFunc(pprof.Heap)
		router.PathPrefix("/debug/pprof/profile").HandlerFunc(pprof.Profile)

		router.HandleFunc("/debug/pprof/", pprof.Index)
		router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		router.HandleFunc("/debug/pprof/profile", pprof.Profile)
		router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	*/
	router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	router.HandleFunc("/debug/goroutines", getGoroutinesCountHandler)
	router.HandleFunc("/debug/stacktrace", getStackTraceHandler)
}

// get the count of number of go routines in the system.
func countGoRoutines() int {
	return runtime.NumGoroutine()
}

func getGoroutinesCountHandler(w http.ResponseWriter, r *http.Request) {
	// Get the count of number of go routines running.
	count := countGoRoutines()
	w.Write([]byte(strconv.Itoa(count)))
}
func getStackTraceHandler(w http.ResponseWriter, r *http.Request) {

	stack := debug.Stack()

	w.Write(stack)

	rtpprof.Lookup("goroutine").WriteTo(w, 2)

}

var block chan int

func SayHello(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 5; i++ {
		go func() {
			block <- 1
		}()
	}

	for i := 0; i < 10000000000; i++ {
		math.Pow(36, 89)
	}

	fmt.Fprint(w, "Hello!")
}

func main() {
	r := mux.NewRouter()
	AttachProfiler(r)
	r.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":6060", r)
}
