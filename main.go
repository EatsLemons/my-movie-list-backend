package main

import (
	"my-movie-list/rest/api"
	"my-movie-list/rest/auth"
	"my-movie-list/service/tmdb"
	"os"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Port int `long:"port" env:"MY_MOVIE_LIST_PORT" default:"8080" description:"port"`

	TmdbAPIKey string `long:"tmdb-api-key" env:"TMDB_API_KEY" default:"" description:"Tmdb api auth key"`
}

func main() {
	p := flags.NewParser(&opts, flags.Default)
	if _, e := p.ParseArgs(os.Args[1:]); e != nil {
		os.Exit(1)
	}

	tmbdClient := tmdb.MakeTmdbAPIClient(opts.TmdbAPIKey)
	authenticator := auth.NewAuthenticator()

	server := &api.Rest{
		TmbdClient:    tmbdClient,
		Authenticator: authenticator,
	}

	server.Run(opts.Port)
}
