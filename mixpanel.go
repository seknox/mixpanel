package mixpanel

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Config struct {
	APIEndpoint string
	Verbose int 
	Token string 
}


const (
	MixpanelEndpoint = "https://api.mixpanel.com"
)


func NewClient(token string, verbose int) *Config {
 return &Config{
	APIEndpoint: MixpanelEndpoint,
	Verbose: verbose,
	Token: token,
 }
}

// Track is mixpanel tracking event
type Track struct {
	Event string `json:"event"`
	Properties TrackProperties  `json:"properties"`
	// Returns data if req is successfull.Value should be 0 or 1.
	Verbose int `json:"verbose"`
}

type TrackProperties struct {
	// Unique user ID
	DistinctID string 
	// Project token
	Token string 
	// User remote IP address
	IP string 
	// Event time
	Time int64
	// Unique event ID
	InsertID string
}

// Send Track request
func (con Config) Track(track Track) error {

	track.Properties.Token = con.Token
	reqBody, err := json.Marshal(track)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", con.APIEndpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

