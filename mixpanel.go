package mixpanel

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Mixpanel client config
type Config struct {
	Verbose int 
	Token string 
}


const (
	TrackEndpoint = "https://api.mixpanel.com/track#live-event"
)

//  NewClient initializes new config
func NewClient(token string, verbose int) *Config {
 return &Config{
	Verbose: verbose,
	Token: token,
 }
}

// Track is mixpanel tracking event
type Track struct {
	Event string `json:"event"`
	Properties TrackProperties  `json:"properties"`
}

type TrackProperties struct {
	// Unique user ID
	DistinctID string `json:"distinct_id"`
	// Project token
	Token string `json:"token"`
	// User remote IP address
	IP string  `json:"ip"`
	// Event time
	Time int64 `json:"time"`
	// Unique event ID
	InsertID string `json:"$insert_id"`
}

// Track sends mixpanel trackevent - https://developer.mixpanel.com/reference/events#track-event
func (con Config) Track(track Track) error {

	track.Properties.Token = con.Token


	data, err := json.Marshal(track)
	if err != nil {
		return err
	}

	reqBody := fmt.Sprintf("data=%s", string(data))


	req, err := http.NewRequest("POST", TrackEndpoint, strings.NewReader(reqBody))
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

