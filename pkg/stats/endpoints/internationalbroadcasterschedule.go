package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
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

	var apiResp struct {
		ResultSets []map[string]interface{} `json:"resultSets"`
	}
	if err := json.Unmarshal(rawResp.Body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(apiResp.ResultSets) == 0 {
		return &InternationalBroadcasterScheduleResponse{Games: []ScheduledGame{}}, nil
	}

	nextGameListRaw, ok := apiResp.ResultSets[0]["NextGameList"]
	if !ok {
		return &InternationalBroadcasterScheduleResponse{Games: []ScheduledGame{}}, nil
	}

	nextGameListJSON, err := json.Marshal(nextGameListRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal NextGameList: %w", err)
	}

	var games []ScheduledGame
	if err := json.Unmarshal(nextGameListJSON, &games); err != nil {
		return nil, fmt.Errorf("failed to unmarshal games: %w", err)
	}

	return &InternationalBroadcasterScheduleResponse{Games: games}, nil
}
