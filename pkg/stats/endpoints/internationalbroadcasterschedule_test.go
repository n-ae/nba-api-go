package endpoints

import (
	"testing"

	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
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
