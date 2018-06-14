package main

import (
	"my-movie-list/app/api"
	"my-movie-list/app/tmdb"
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

	server := &api.Rest{
		TmbdClient: tmbdClient,
	}

	server.Run(opts.Port)
}
