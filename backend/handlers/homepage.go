package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"groupie/backend"
	"groupie/backend/models"
)

// Homepage is an HTTP handler that serves the homepage of the web application.

func Homepage(w http.ResponseWriter, r *http.Request) {
	tmp2, err := template.ParseFiles("./frontend/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		tmp2.Execute(w, "Page Not Found")
		return
	}

	// Ensure the request method is GET; otherwise, return a 405 Method Not Allowed.

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		tmp2.Execute(w, "Method Not Allowed")
		return
	}

	// Fetch the list of artists from the backend and store it in the data slice.

	data := []models.Artist{}
	backend.FetchData(&data, "/artists")
	locationsidx := models.Locations{}
	backend.FetchData(&locationsidx, "/locations")
	for i := 0; i < len(locationsidx.Index); i++ {
		loca := []string{}
		for j := 0; j < len(locationsidx.Index[i].Locations); j++ {
			loca = append(loca, strings.ReplaceAll(locationsidx.Index[i].Locations[j], "-", ", "))
		}
		data[i].Locations = loca
		data[i].MembersLen = len(data[i].Members)
	}

	// Execute the template with the artist data and write it to the response.

	tmpl, err := template.ParseFiles("./frontend/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmp2.Execute(w, "Error parsing template")
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmp2.Execute(w, "Error executing template")
		return
	}
}
