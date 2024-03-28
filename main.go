package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// API endpoints
var (
	artistsEndpoint   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsEndpoint = "https://groupietrackers.herokuapp.com/api/locations"
)

// Structs to unmarshal API responses
type Artist struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Year         int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDate"`
	Relations    string   `json:"relations"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Handler function to fetch artists data
func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	var artistsData []Artist
	// Make GET request to artists API endpoint
	resp, err := http.Get(artistsEndpoint)
	if err != nil {
		http.Error(w, "Failed to fetch artists data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode JSON response
	err = json.NewDecoder(resp.Body).Decode(&artistsData)
	if err != nil {
		http.Error(w, "Failed to decode artists data", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "Failed to parse file index.html", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artistsData)
	if err != nil {
		http.Error(w, "Failed to execute file index.html", http.StatusInternalServerError)
		return
	}
}

func relationPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path != "/relation" {
		http.NotFound(w, r)
		return
	}
	relationsEndpoint := r.FormValue("relationlink")
	fmt.Println(relationsEndpoint)
	var relationData Relation
	resp, err := http.Get(relationsEndpoint)
	if err != nil {
		http.Error(w, "Failed to fetch relation data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode JSON response
	err = json.NewDecoder(resp.Body).Decode(&relationData)
	if err != nil {
		http.Error(w, "Failed to decode artists data", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("template/relation.html")
	if err != nil {
		http.Error(w, "Failed to parse file relation.html", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, relationData)
	if err != nil {
		http.Error(w, "Failed to execute file relation.html", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Define HTTP routes
	http.HandleFunc("/", homePage)
	http.HandleFunc("/relation", relationPage)

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
