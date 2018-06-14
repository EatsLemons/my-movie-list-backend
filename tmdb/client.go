package tmdb

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// TmdbAPI is a http service client to Tmdb service
type TmdbAPI struct {
	apiKey string

	domain     string
	httpClient *http.Client
}

// MakeTmdbAPIClient is a TmdbAPI's constructor
func MakeTmdbAPIClient(apiKey string) *TmdbAPI {
	return &TmdbAPI{
		apiKey: apiKey,

		domain: "https://api.themoviedb.org/3",
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// SearchMovie request movies list from TMDB by query string
func (c *TmdbAPI) SearchMovie(query string) (*SearchMovieResponse, error) {
	requestURI := "/search/movie?language=ru-RU&page=1&include_adult=true&query=" + query

	response := &SearchMovieResponse{}
	err := c.makeGetRequest(requestURI, response)
	if err != nil {
		log.Printf("[WARN] search movie request to tmdb has failed %s for query %s", err, query)
		return nil, err
	}

	return response, nil
}

// GetMovieCredits request a team of movie developers, cast and crew
func (c *TmdbAPI) GetMovieCredits(movieID string) (*MovieCreditsResponse, error) {
	requestURI := "/movie/" + movieID + "/credits?"

	response := &MovieCreditsResponse{}
	err := c.makeGetRequest(requestURI, response)
	if err != nil {
		log.Printf("[WARN] get movie credits request to tmdb has failed %s for movie %s", err, movieID)
		return nil, err
	}

	return response, nil
}

func (c *TmdbAPI) makeGetRequest(url string, result interface{}) error {
	r, err := c.httpClient.Get(c.domain + url + "&api_key=" + c.apiKey)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(result)
}

type MovieCreditsResponse struct {
	ID   int `json:"id"`
	Cast []struct {
		CastID      int    `json:"cast_id"`
		Character   string `json:"character"`
		CreditID    string `json:"credit_id"`
		Gender      int    `json:"gender"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Order       int    `json:"order"`
		ProfilePath string `json:"profile_path"`
	} `json:"cast"`
	Crew []struct {
		CreditID    string `json:"credit_id"`
		Department  string `json:"department"`
		Gender      int    `json:"gender"`
		ID          int    `json:"id"`
		Job         string `json:"job"`
		Name        string `json:"name"`
		ProfilePath string `json:"profile_path"`
	} `json:"crew"`
}

type SearchMovieResponse struct {
	Page    int `json:"page"`
	Results []struct {
		PosterPath       string  `json:"poster_path"`
		Adult            bool    `json:"adult"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		GenreIds         []int   `json:"genre_ids"`
		ID               int     `json:"id"`
		OriginalTitle    string  `json:"original_title"`
		OriginalLanguage string  `json:"original_language"`
		Title            string  `json:"title"`
		BackdropPath     string  `json:"backdrop_path"`
		Popularity       float64 `json:"popularity"`
		VoteCount        int     `json:"vote_count"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
	} `json:"results"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
}
