package pprof_init

import (
	"log"
	"net/http"
	"net/http/pprof"
)

func init() {
	go func() {
		r := http.NewServeMux()

		r.HandleFunc("/debug/pprof/", pprof.Index)
		r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		r.HandleFunc("/debug/pprof/profile", pprof.Profile)
		r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		r.HandleFunc("/debug/pprof/trace", pprof.Trace)

		if err := http.ListenAndServe(":6060", r); err != nil {
			log.Printf("error serving pprof: %v", err)
		}
	}()
}
