package main

import (
	"log"

	"github.com/kr/pretty"
	"github.com/stevenpaz/mlb-statsapi-go/src/pkg/mlbstatsapi"
)

func main() {
	mlbstatsapi := mlbstatsapi.New()

	teams, err := mlbstatsapi.GetTeams()
	if err != nil {
		log.Fatal(err)

		return
	}

	// pretty print teams to console output
	pretty.Println(teams)
}
