package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupie/backend"
	"groupie/backend/models"
)

// ArtistPage is an HTTP handler that serves the artist's information page.

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	tmp2, err := template.ParseFiles("./frontend/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("./frontend/artist.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmp2.Execute(w, "Error parsing template")
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		tmp2.Execute(w, "Method Not Allowed")
		return
	}
	// Declare variables to store artist data fetched from the backend.

	var artist []models.Artist
	var artistRelation models.Relation
	var artistDate models.Date
	var artistLocation models.Location

	// Retrieve the "id" query parameter from the request URL.

	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if strId == "" {
		w.WriteHeader(http.StatusNotFound)
		tmp2.Execute(w, "Page Not Found")
		return
	}

	// Fetch artist-related data from multiple backend API endpoints.

	backend.FetchData(&artist, "/artists")
	backend.FetchData(&artistLocation, "/locations/"+strId)
	backend.FetchData(&artistDate, "/dates/"+strId)
	backend.FetchData(&artistRelation, "/relation/"+strId)

	// Validate the ID: check for conversion errors, non-positive IDs, or IDs out of range.
	if err != nil || id <= 0 || id > len(artist) {
		w.WriteHeader(http.StatusBadRequest)
		tmp2.Execute(w, "Invalid artist ID")
		return
	}

	// Editing data befor sending it
	for i := 0; i < len(artistDate.Dates); i++ {
		artistDate.Dates[i] = strings.Trim(artistDate.Dates[i], "*")
	}

	for i := 0; i < len(artistLocation.Locations); i++ {
		artistLocation.Locations[i] = strings.ReplaceAll(artistLocation.Locations[i], "_", " ")
		artistLocation.Locations[i] = strings.ReplaceAll(artistLocation.Locations[i], "-", " - ")
	}

	artistRelations := make(map[string][]string)
	for locations, dates := range artistRelation.DatesLocations {
		artistRelations[strings.ReplaceAll(strings.ReplaceAll(strings.Title(locations), "-", " ➔ "), "_", " ")] = dates
	}
	artistRelation.DatesLocations = artistRelations

	// Struct to store the specific artist's details, along with related data.

	data := struct {
		Myartist          models.Artist
		MyartistRelations models.Relation
		MyartistDates     models.Date
		MyartistLocations models.Location
	}{
		Myartist:          artist[id-1],
		MyartistRelations: artistRelation,
		MyartistDates:     artistDate,
		MyartistLocations: artistLocation,
	}

	// Execute the parsed template and pass the artist data to it.

	if err := tmpl.Execute(w, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmp2.Execute(w, "Error executing template")
		return
	}
}
