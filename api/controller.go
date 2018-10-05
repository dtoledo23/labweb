package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// func Create

func PlayerController() *mux.Router {
	r := mux.NewRouter()

	// Get player
	r.
		PathPrefix("/player/{id}/").
		Methods(http.MethodGet).
		HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			i, err := strconv.ParseInt(mux.Vars(req)["id"], 10, 64)
			if err != nil {
				resp.WriteHeader(http.StatusBadRequest)
				resp.Write([]byte("Bad request"))
				return
			}

			fmt.Println("executed this")
			player := db.Get(int(i))[0]
			json.NewEncoder(resp).Encode(player)
		})

	// List all players
	r.
		PathPrefix("/player/").
		Methods(http.MethodGet).
		HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			json.NewEncoder(resp).Encode(db.ListAll())
		})

	// Delete player
	r.
		PathPrefix("/player/{id}").
		Methods(http.MethodDelete).
		HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			id, err := strconv.ParseInt(mux.Vars(req)["id"], 10, 64)
			if err != nil {
				resp.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(resp, "Bad request")
				return
			}

			if err := db.Delete(int(id)); err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(resp, err)
				return
			}

			fmt.Fprintf(resp, "Deleted %d", id)
		})

	// Delete players passed on query
	r.
		PathPrefix("/player").
		Methods(http.MethodDelete).
		Queries("ids", "{ids}").
		HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			ids := make([]int, 0)
			for _, query := range strings.Split(mux.Vars(req)["ids"], ",") {
				id, err := strconv.ParseInt(query, 10, 64)
				if err != nil {
					resp.WriteHeader(http.StatusBadRequest)
					fmt.Fprintln(resp, "Invalid player id", query)
					return
				}
				ids = append(ids, int(id))
			}

			if err := db.Delete(ids...); err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(resp, err)
				return
			}

			fmt.Fprintf(resp, "Deleted %v", ids)
		})

	// Edit player
	r.
		PathPrefix("/player/{id}").
		Methods(http.MethodPut).
		HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(resp, "Edit player")
		})

	// Add player
	r.
		PathPrefix("/player").
		Methods(http.MethodPost).
		HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(resp, "Add player")
		})

	return r
}
