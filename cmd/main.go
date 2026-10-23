package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

type FileHandler struct {
	staticFilePath string
	filename string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).

func (fh FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(fh.staticFilePath, r.URL.Path)

	fp, err := os.Stat(path)

	if(os.IsNotExist(err) || fp.IsDir()){
		http.ServeFile(w, r, filepath.Join(fh.staticFilePath, fh.filename))
		return
	}

	if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static file
	http.FileServer(http.Dir(fh.staticFilePath)).ServeHTTP(w, r)
}

func main() {
	var wait time.Duration
    flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.Parse()

	//create routes and register this router
	
	routes := mux.NewRouter()

	routes.HandleFunc("/api/health", func (w http.ResponseWriter, r *http.Request)  {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	spa := FileHandler{staticFilePath: "public", filename: "index.html"}
	routes.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Addr: ":8080",
		Handler: routes,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
		IdleTimeout: 24 * time.Minute,
	}

	go func()  {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Something weird happened: %v\n", err)
		}
	}()

	fmt.Printf("[SERVER ] Listenin on :8080 and Waiting time is %s", wait)

	q := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(q, os.Interrupt)

	<-q

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	fmt.Println("Server Shutdown Gracefully!")
	os.Exit(0)
}

