package handlers

import (
	"html/template"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, r *http.Request, errMsg string, errCode int) {
	tmp2, err := template.ParseFiles("./frontend/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmp2.Execute(w, "Error parsing template")
		return
	}
	tmpl, err := template.ParseFiles("./frontend/artist.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmp2.Execute(w, "Error parsing template")
		return
	}
	w.WriteHeader(errCode)

	err = tmpl.Execute(w, errMsg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmp2.Execute(w, "Internal server error!")
		return
	}
}
