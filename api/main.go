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
	Title string `json:"title"`
	URL   string `json:"URL"`
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	values := m.GetRandom()

	out := out{
		values.Title,
		values.URL,
	}

	fmt.Printf("%d ", 1)

	json.NewEncoder(w).Encode(out)
}
