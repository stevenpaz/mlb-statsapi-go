package mlbstatsapi

import "fmt"

type Venue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type League struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Division struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Sport struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetTeamsResponse struct {
	Copyright string  `json:"copyright"`
	Teams     []*Team `json:"teams"`
}

type Team struct {
	ID              int       `json:"id"`
	AllStarStatus   string    `json:"allStarStatus"`
	Name            string    `json:"name"`
	Season          int       `json:"season"`
	Venue           *Venue    `json:"venue"`
	TeamCode        string    `json:"teamCode"`
	FileCode        string    `json:"fileCode"`
	Abbreviation    string    `json:"abbreviation"`
	TeamName        string    `json:"teamName"`
	LocationName    string    `json:"locationName"`
	FirstYearOfPlay string    `json:"firstYearOfPlay"`
	League          *League   `json:"league"`
	Division        *Division `json:"division"`
	Sport           *Sport    `json:"sport"`
	ShortName       string    `json:"shortName"`
	ParentOrgName   string    `json:"parentOrgName"`
	ParentOrgID     int       `json:"parentOrgId"`
	FranchiseName   string    `json:"franchiseName"`
	ClubName        string    `json:"clubName"`
	Active          bool      `json:"active"`
}

func (api *MLBStatsAPI) GetTeams() ([]*Team, error) {
	var resp GetTeamsResponse

	err := api.get("teams", &resp)
	if err != nil {
		return nil, fmt.Errorf("error getting teams: %w", err)
	}

	return resp.Teams, nil
}
