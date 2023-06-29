package mlbstatsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultTimeoutSeconds = 10
	defaultAPIVersion     = "v1"
	apiBaseURLFormat      = "https://statsapi.mlb.com/api/%s"
)

type Config struct {
	TimeoutSeconds int
	APIBaseURL     string
}

func DefaultConfig() *Config {
	return NewConfig(defaultAPIVersion)
}

func NewConfig(apiVersion string) *Config {
	return &Config{
		APIBaseURL: fmt.Sprintf(apiBaseURLFormat, apiVersion),
	}
}

type MLBStatsAPI struct {
	Config *Config
	client *http.Client
}

func New() *MLBStatsAPI {
	return &MLBStatsAPI{
		Config: DefaultConfig(),
		client: &http.Client{
			Timeout: defaultTimeoutSeconds * time.Second,
		},
	}
}

func (api *MLBStatsAPI) get(relURL string, v any) error {
	url := api.Config.APIBaseURL + "/" + relURL

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error building GET %s: %w", url, err)
	}

	resp, err := api.client.Do(req)
	if err != nil {
		return fmt.Errorf("error calling GET %s: %w", url, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response for GET %s: %w", url, err)
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &v)
	if err != nil {
		return fmt.Errorf("error reading response for GET %s: %w", url, err)
	}

	return nil
}
