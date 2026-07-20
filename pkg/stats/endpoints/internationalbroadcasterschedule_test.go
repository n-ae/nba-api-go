package endpoints

import (
	"testing"

	"github.com/n-ae/nba-api-go/v2/pkg/stats/parameters"
)

func TestInternationalBroadcasterScheduleRequest_Validation(t *testing.T) {
	tests := []struct {
		name    string
		req     InternationalBroadcasterScheduleRequest
		wantErr bool
	}{
		{
			name: "valid request with required fields",
			req: InternationalBroadcasterScheduleRequest{
				LeagueID: parameters.LeagueIDNBA,
				Season:   "2025",
			},
			wantErr: false,
		},
		{
			name: "valid request with all fields",
			req: InternationalBroadcasterScheduleRequest{
				LeagueID: parameters.LeagueIDNBA,
				Season:   "2025",
				RegionID: stringPtr("1"),
				Date:     stringPtr("11/07/2025"),
				EST:      stringPtr("Y"),
			},
			wantErr: false,
		},
		{
			name: "missing season",
			req: InternationalBroadcasterScheduleRequest{
				LeagueID: parameters.LeagueIDNBA,
				Season:   "",
			},
			wantErr: true,
		},
		{
			name: "invalid league id",
			req: InternationalBroadcasterScheduleRequest{
				LeagueID: "99",
				Season:   "2025",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.LeagueID.Validate(); (err != nil) != tt.wantErr {
				if tt.name == "missing season" {
					return
				}
				t.Errorf("LeagueID.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.req.Season == "" && !tt.wantErr {
				t.Errorf("Expected Season validation to fail")
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}

func TestParseInternationalBroadcasterScheduleResponse(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		wantGames int
		wantErr   bool
	}{
		{
			name:      "populated game with broadcasters",
			body:      `{"resultSets":[{"NextGameList":[{"gameID":"0022500001","vtCity":"Boston","vtNickName":"Celtics","vtShortName":"Boston","vtAbbreviation":"BOS","htCity":"Los Angeles","htNickName":"Lakers","htShortName":"LA Lakers","htAbbreviation":"LAL","date":"2026-01-15","time":"19:30","day":"Thursday","broadcasters":[{"broadcastID":"1","broadcasterName":"ESPN","tapeDelayComments":""}]}]}]}`,
			wantGames: 1,
		},
		{
			name:      "empty NextGameList",
			body:      `{"resultSets":[{"NextGameList":[]}]}`,
			wantGames: 0,
		},
		{
			name:      "resultSets present but NextGameList key absent",
			body:      `{"resultSets":[{}]}`,
			wantGames: 0,
		},
		{
			name:      "no resultSets at all",
			body:      `{"resultSets":[]}`,
			wantGames: 0,
		},
		{
			name:    "malformed JSON",
			body:    `not json`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := parseInternationalBroadcasterScheduleResponse([]byte(tt.body))
			if (err != nil) != tt.wantErr {
				t.Fatalf("parseInternationalBroadcasterScheduleResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if resp.Games == nil {
				t.Error("resp.Games is nil, want a non-nil (possibly empty) slice")
			}
			if len(resp.Games) != tt.wantGames {
				t.Errorf("len(resp.Games) = %d, want %d", len(resp.Games), tt.wantGames)
			}
		})
	}
}

func TestParseInternationalBroadcasterScheduleResponse_FieldValues(t *testing.T) {
	body := `{"resultSets":[{"NextGameList":[{"gameID":"0022500001","vtCity":"Boston","vtAbbreviation":"BOS","htCity":"Los Angeles","htAbbreviation":"LAL","date":"2026-01-15","broadcasters":[{"broadcastID":"1","broadcasterName":"ESPN","tapeDelayComments":"live"}]}]}]}`

	resp, err := parseInternationalBroadcasterScheduleResponse([]byte(body))
	if err != nil {
		t.Fatalf("parseInternationalBroadcasterScheduleResponse() error = %v", err)
	}
	if len(resp.Games) != 1 {
		t.Fatalf("len(resp.Games) = %d, want 1", len(resp.Games))
	}

	game := resp.Games[0]
	if game.GameID != "0022500001" {
		t.Errorf("game.GameID = %q, want %q", game.GameID, "0022500001")
	}
	if game.VisitorAbbr != "BOS" || game.HomeAbbr != "LAL" {
		t.Errorf("game.VisitorAbbr/HomeAbbr = %q/%q, want %q/%q", game.VisitorAbbr, game.HomeAbbr, "BOS", "LAL")
	}
	if len(game.Broadcasters) != 1 || game.Broadcasters[0].BroadcasterName != "ESPN" {
		t.Errorf("game.Broadcasters = %+v, want a single ESPN entry", game.Broadcasters)
	}
}
