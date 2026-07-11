package endpoints

import (
	"testing"
	"time"
)

func TestParseBirthDate(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    time.Time
		wantErr bool
	}{
		{
			name: "commonplayerinfo ISO format",
			raw:  "1988-03-14T00:00:00",
			want: time.Date(1988, time.March, 14, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "commonteamroster uppercase month format",
			raw:  "APR 17, 2001",
			want: time.Date(2001, time.April, 17, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "commonteamroster zero-padded day",
			raw:  "JAN 09, 2001",
			want: time.Date(2001, time.January, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			name:    "empty value",
			raw:     "",
			wantErr: true,
		},
		{
			name:    "unrecognized format",
			raw:     "not a date",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBirthDate(tt.raw)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("parseBirthDate(%q) expected error, got nil", tt.raw)
				}
				return
			}
			if err != nil {
				t.Fatalf("parseBirthDate(%q) unexpected error: %v", tt.raw, err)
			}
			if !got.Equal(tt.want) {
				t.Errorf("parseBirthDate(%q) = %v, want %v", tt.raw, got, tt.want)
			}
		})
	}
}

func TestDateOfBirth(t *testing.T) {
	t.Run("PlayerInfo", func(t *testing.T) {
		p := PlayerInfo{Birthdate: "1988-03-14T00:00:00"}
		got, err := p.DateOfBirth()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !got.Equal(time.Date(1988, time.March, 14, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("got %v", got)
		}
	})

	t.Run("CommonPlayerInfoV2CommonPlayerInfo", func(t *testing.T) {
		p := CommonPlayerInfoV2CommonPlayerInfo{BIRTHDATE: "1988-03-14T00:00:00"}
		got, err := p.DateOfBirth()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !got.Equal(time.Date(1988, time.March, 14, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("got %v", got)
		}
	})

	t.Run("CommonTeamRosterCommonTeamRoster", func(t *testing.T) {
		r := CommonTeamRosterCommonTeamRoster{BIRTH_DATE: "APR 17, 2001"}
		got, err := r.DateOfBirth()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !got.Equal(time.Date(2001, time.April, 17, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("got %v", got)
		}
	})

	t.Run("CommonTeamRosterV2CommonTeamRoster", func(t *testing.T) {
		r := CommonTeamRosterV2CommonTeamRoster{BIRTH_DATE: "APR 17, 2001"}
		got, err := r.DateOfBirth()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !got.Equal(time.Date(2001, time.April, 17, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("got %v", got)
		}
	})
}
