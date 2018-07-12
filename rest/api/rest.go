package api

import (
	"encoding/json"
	"fmt"
	"log"
	"my-movie-list/service/tmdb"
	"net/http"
	"remark/backend/app/rest/auth"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gorilla/mux"
)

// Rest is a http api service
type Rest struct {
	Version string

	TmbdClient    *tmdb.TmdbAPI
	Authenticator *auth.Authenticator

	httpServer *http.Server
}

// Run starts Rest api http server
func (rs *Rest) Run(port int) {
	log.Printf("[INFO] server started at :%d", port)

	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1/").Subrouter()

	// Check alive
	v1.HandleFunc("/ping", rs.pingHndlr).Methods("GET")

	// Movie autocomplete
	v1.HandleFunc("/autocomplete/{query}", rs.autocompleteHndlr).Methods("GET")

	rs.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	err := rs.httpServer.ListenAndServe()
	log.Printf("[WARN] http server terminated, %s", err)
}

func (rs *Rest) autocompleteHndlr(w http.ResponseWriter, r *http.Request) {
	response := rs.makeResponseItem()

	vars := mux.Vars(r)
	query, ok := vars["query"]
	if !ok {
		response.Errors = make([]Error, 1)
		response.Errors[0] = Error{
			Message: "query is mandatory field",
		}

		rs.makeJSONResponse(w, response)
		return
	}

	response.AutocompleteMovie = make([]*AutocompleteMovie, 0)

	if utf8.RuneCountInString(query) < 4 {
		rs.makeJSONResponse(w, response)
		return
	}

	searchMovieRS, err := rs.TmbdClient.SearchMovie(query)
	if err != nil {
		response.Errors = make([]Error, 1)
		response.Errors[0] = Error{
			Message: err.Error(),
		}

		rs.makeJSONResponse(w, response)
		return
	}

	resultCounts := len(searchMovieRS.Results)

	autocompleteCount := 4
	for i := 0; (i < autocompleteCount && resultCounts > autocompleteCount) || (i < resultCounts && resultCounts < autocompleteCount); i++ {
		dateSlice := strings.Split(searchMovieRS.Results[i].ReleaseDate, "-")
		autocompleteItem := &AutocompleteMovie{
			Name:  searchMovieRS.Results[i].Title,
			Year:  dateSlice[0],
			Image: "https://image.tmdb.org/t/p/w500" + searchMovieRS.Results[i].PosterPath,
		}

		response.AutocompleteMovie = append(response.AutocompleteMovie, autocompleteItem)

		credits, _ := rs.TmbdClient.GetMovieCredits(strconv.Itoa(searchMovieRS.Results[i].ID))
		castCount := len(credits.Cast)
		if credits != nil {
			for j := 0; (castCount >= 3 && j < 3) || (castCount < 3 && castCount > j); j++ {
				autocompleteItem.Actors = append(autocompleteItem.Actors, credits.Cast[j].Name)
			}
		}
	}

	rs.makeJSONResponse(w, response)
}

func (rs *Rest) pingHndlr(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

func (rs *Rest) makeJSONResponse(w http.ResponseWriter, response interface{}) {
	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (rs *Rest) makeResponseItem() *Response {
	return &Response{
		SystemInfo: &SystemInfo{
			Version: rs.Version,
		},
	}
}

// AutocompleteMovie is a response DTO of autocomplete movie request
type AutocompleteMovie struct {
	Name   string   `json:"name,omitempty"`
	Actors []string `json:"actors,omitempty"`
	Year   string   `json:"release_year,omitempty"`
	Image  string   `json:"image_url,omitempty"`
}

type Response struct {
	AutocompleteMovie []*AutocompleteMovie `json:"autocomplete_movie"`

	Errors     []Error     `json:"errors,omitempty"`
	SystemInfo *SystemInfo `json:"system_info"`
}

type Error struct {
	Message string `json:"message,omitempty"`
}

type SystemInfo struct {
	Version string `json:"version,omitempty"`
}
