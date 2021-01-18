package main

import (
	"api/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var m *database.Connection

func main() {
	time.Sleep(10)
	m = database.NewConnection("mongodb://database:27017", "journal", "hackernews")
	defer m.Close()

	router := mux.NewRouter()
	router.HandleFunc("/random", articlesHandler)
	fmt.Println("Listening at 8081")
	http.ListenAndServe(":8081", router)
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

	json.NewEncoder(w).Encode(out)
}
