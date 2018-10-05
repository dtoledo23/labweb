package main

import (
	"log"
	"net/http"
)

var db PlayersDatabase

func init() {
	db = NewInMemoryPlayersDatabase()

	db.Add(
		Player{
			Name:   "Victor",
			Number: 10,
		},
		Player{
			Name:   "Kevin",
			Number: 5,
		},
		Player{
			Name:   "Samuel",
			Number: 8,
		},
	)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := PlayerController()
	r.Use(loggingMiddleware)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	log.Panic(server.ListenAndServe())
}
