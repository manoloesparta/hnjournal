package main

import (
	"api/database"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var m *database.Connection

func main() {
	m = database.NewConnection("mongodb://localhost:27017", "journal", "hackernews")
	defer m.Close()

	router := mux.NewRouter()
	router.HandleFunc("/random", articlesHandler)
	fmt.Println(":8080 serving")
	http.ListenAndServe(":8080", router)
}

type out struct {
	Title1 string `json:"title1"`
	URL1   string `json:"URL1"`
	Title2 string `json:"title2"`
	URL2   string `json:"URL2"`
	Title3 string `json:"title3"`
	URL3   string `json:"URL3"`
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	values := m.GetRandom()

	out := out{
		values[0].Title,
		values[0].URL,
		values[1].Title,
		values[1].URL,
		values[2].Title,
		values[2].URL,
	}

	json.NewEncoder(w).Encode(out)
}
