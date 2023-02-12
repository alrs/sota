package sota

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

const timeFormat = "2006-01-02T15:04:05.999999999"

type Spot struct {
	ID                int       `json:"id"`
	UserID            int       `json:"userID"`
	TimeStamp         time.Time `json:"timeStamp"`
	Comments          string    `json:"comments"`
	Callsign          string    `json:"callsign"`
	AssociationCode   string    `json:"associationCode"`
	SummitCode        string    `json:"summitCode"`
	ActivatorCallsign string    `json:"activatorCallsign"`
	ActivatorName     string    `json:"activatorName"`
	Frequency         float32   `json:"frequency"`
	Mode              string    `json:"mode"`
	SummitDetails     string    `json:"summitDetails"`
	HighlightColor    string    `json:"higlightColor"`
}

func (s *Spot) UnmarshalJSON(data []byte) error {
	type rawBlob struct {
		ID                int    `json:"id"`
		UserID            int    `json:"userID"`
		TimeStamp         string `json:"timeStamp"`
		Comments          string `json:"comments"`
		Callsign          string `json:"callsign"`
		AssociationCode   string `json:"associationCode"`
		SummitCode        string `json:"summitCode"`
		ActivatorCallsign string `json:"activatorCallsign"`
		ActivatorName     string `json:"activatorName"`
		Frequency         string `json:"frequency"`
		Mode              string `json:"mode"`
		SummitDetails     string `json:"summitDetails"`
		HighlightColor    string `json:"higlightColor"`
	}
	var rb rawBlob
	err := json.Unmarshal(data, &rb)
	if err != nil {
		return err
	}
	s.ID = rb.ID
	s.UserID = rb.UserID
	s.TimeStamp, err = time.Parse(timeFormat, rb.TimeStamp)
	if err != nil {
		return err
	}
	s.Comments = rb.Comments
	s.Callsign = rb.Callsign
	s.AssociationCode = rb.AssociationCode
	s.SummitCode = rb.SummitCode
	s.ActivatorCallsign = rb.ActivatorCallsign
	s.ActivatorName = rb.ActivatorName
	f, err := strconv.ParseFloat(rb.Frequency, 32)
	s.Frequency = float32(f)
	s.Mode = strings.ToLower(rb.Mode)
	s.SummitDetails = rb.SummitDetails
	s.HighlightColor = rb.HighlightColor
	return nil
}