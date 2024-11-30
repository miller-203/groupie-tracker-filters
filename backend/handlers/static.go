package handlers

import (
	"html/template"
	"net/http"
	"os"
)

// Static is an HTTP handler that serves static files from the server.

func Static(w http.ResponseWriter, r *http.Request) {
	// Use defer to recover from potential panics and return a 404 page not found error.
	tmp2, err := template.ParseFiles("./frontend/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusNotFound)
			tmp2.Execute(w, "page not found")
			return
		}
	}()

	// Check if the requested file exists on the server and is not a directory.

	file, err := os.Stat("." + r.URL.Path)
	if !os.IsNotExist(err) && !file.IsDir() {
		http.FileServer(http.Dir("./")).ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusForbidden)
		tmp2.Execute(w, "Forbidden access to a directory via URL")
		return
	}
}
