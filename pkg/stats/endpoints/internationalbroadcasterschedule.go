package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/v3/pkg/models"
	"github.com/n-ae/nba-api-go/v3/pkg/stats"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

type InternationalBroadcasterScheduleRequest struct {
	LeagueID parameters.LeagueID
	Season   string
	RegionID *string
	Date     *string
	EST      *string
}

type Broadcaster struct {
	BroadcastID       string `json:"broadcastID"`
	BroadcasterName   string `json:"broadcasterName"`
	TapeDelayComments string `json:"tapeDelayComments"`
}

type ScheduledGame struct {
	GameID           string        `json:"gameID"`
	VisitorCity      string        `json:"vtCity"`
	VisitorNickName  string        `json:"vtNickName"`
	VisitorShortName string        `json:"vtShortName"`
	VisitorAbbr      string        `json:"vtAbbreviation"`
	HomeCity         string        `json:"htCity"`
	HomeNickName     string        `json:"htNickName"`
	HomeShortName    string        `json:"htShortName"`
	HomeAbbr         string        `json:"htAbbreviation"`
	Date             string        `json:"date"`
	Time             string        `json:"time"`
	Day              string        `json:"day"`
	Broadcasters     []Broadcaster `json:"broadcasters"`
}

type InternationalBroadcasterScheduleResponse struct {
	Games []ScheduledGame
}

func GetInternationalBroadcasterSchedule(ctx context.Context, client *stats.Client, req InternationalBroadcasterScheduleRequest) (*InternationalBroadcasterScheduleResponse, error) {
	if err := req.LeagueID.Validate(); err != nil {
		return nil, fmt.Errorf("%w: LeagueID: %v", models.ErrInvalidRequest, err)
	}
	if req.Season == "" {
		return nil, fmt.Errorf("%w: Season is required", models.ErrInvalidRequest)
	}

	params := url.Values{}
	params.Set("LeagueID", req.LeagueID.String())
	params.Set("Season", req.Season)
	if req.RegionID != nil {
		params.Set("RegionID", *req.RegionID)
	}
	if req.Date != nil {
		params.Set("Date", *req.Date)
	}
	if req.EST != nil {
		params.Set("EST", *req.EST)
	}

	rawResp, err := client.Get(ctx, "internationalbroadcasterschedule", params)
	if err != nil {
		return nil, err
	}

	return parseInternationalBroadcasterScheduleResponse(rawResp.Body)
}

// parseInternationalBroadcasterScheduleResponse decodes a raw
// internationalbroadcasterschedule response body directly into
// ScheduledGame, rather than via an interface{}/map[string]interface{}
// intermediate re-marshaled and re-unmarshaled a second time: the
// "NextGameList" key is a fixed, known field name, so a plain nested
// struct with a json tag decodes it in one pass with no interface{}
// anywhere. Split out from GetInternationalBroadcasterSchedule so this
// parsing logic is testable without a live HTTP call.
func parseInternationalBroadcasterScheduleResponse(body []byte) (*InternationalBroadcasterScheduleResponse, error) {
	var apiResp struct {
		ResultSets []struct {
			NextGameList []ScheduledGame `json:"NextGameList"`
		} `json:"resultSets"`
	}
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(apiResp.ResultSets) == 0 {
		return &InternationalBroadcasterScheduleResponse{Games: []ScheduledGame{}}, nil
	}

	games := apiResp.ResultSets[0].NextGameList
	if games == nil {
		games = []ScheduledGame{}
	}

	return &InternationalBroadcasterScheduleResponse{Games: games}, nil
}
