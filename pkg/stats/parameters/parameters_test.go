package parameters

import "testing"

func TestPerMode_Validate(t *testing.T) {
	tests := []struct {
		name    string
		perMode PerMode
		wantErr bool
	}{
		{
			name:    "valid Totals",
			perMode: PerModeTotals,
			wantErr: false,
		},
		{
			name:    "valid PerGame",
			perMode: PerModePerGame,
			wantErr: false,
		},
		{
			name:    "invalid",
			perMode: PerMode("Invalid"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.perMode.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("PerMode.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLeagueID_Validate(t *testing.T) {
	tests := []struct {
		name     string
		leagueID LeagueID
		wantErr  bool
	}{
		{
			name:     "valid NBA",
			leagueID: LeagueIDNBA,
			wantErr:  false,
		},
		{
			name:     "valid ABA",
			leagueID: LeagueIDABA,
			wantErr:  false,
		},
		{
			name:     "empty (allowed)",
			leagueID: "",
			wantErr:  false,
		},
		{
			name:     "invalid",
			leagueID: LeagueID("99"),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.leagueID.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("LeagueID.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewSeason(t *testing.T) {
	tests := []struct {
		name string
		year int
		want string
	}{
		{
			name: "2023-24 season",
			year: 2023,
			want: "2023-24",
		},
		{
			name: "1999-00 season",
			year: 1999,
			want: "1999-00",
		},
		{
			name: "2000-01 season",
			year: 2000,
			want: "2000-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeason(tt.year)
			if got.String() != tt.want {
				t.Errorf("NewSeason() = %v, want %v", got, tt.want)
			}
		})
	}
}
