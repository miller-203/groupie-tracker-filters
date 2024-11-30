package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiURL = "https://groupietrackers.herokuapp.com/api"

// FetchData sends a GET request to the specified API endpoint (appending the link to apiURL)

func FetchData(data interface{}, link string) (err error) {
	resp, err := http.Get(apiURL + link)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(data)
	return nil
}
