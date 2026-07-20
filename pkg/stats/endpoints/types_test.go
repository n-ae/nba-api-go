package endpoints

import "testing"

func TestFindResultSet(t *testing.T) {
	sets := []resultSet{
		{Name: "PlayerStats", Headers: []string{"A"}},
		{Name: "TeamStats", Headers: []string{"B"}},
	}

	got, ok := findResultSet(sets, "TeamStats")
	if !ok {
		t.Fatal("findResultSet(sets, \"TeamStats\") returned ok=false, want true")
	}
	if got.Name != "TeamStats" || len(got.Headers) != 1 || got.Headers[0] != "B" {
		t.Errorf("findResultSet(sets, \"TeamStats\") = %+v, want the TeamStats entry", got)
	}

	if _, ok := findResultSet(sets, "DoesNotExist"); ok {
		t.Error("findResultSet(sets, \"DoesNotExist\") returned ok=true, want false")
	}

	if _, ok := findResultSet(nil, "PlayerStats"); ok {
		t.Error("findResultSet(nil, ...) returned ok=true, want false")
	}
}

// TestFindResultSet_KeysByNameNotPosition is the specific regression this
// function exists to prevent: rawResp.ResultSets[0]-style positional
// indexing silently reads the wrong result set into the wrong struct field
// if NBA.com ever returns result sets in a different order than expected.
func TestFindResultSet_KeysByNameNotPosition(t *testing.T) {
	sets := []resultSet{
		{Name: "TeamStats", Headers: []string{"team-data"}},
		{Name: "PlayerStats", Headers: []string{"player-data"}},
	}

	got, ok := findResultSet(sets, "PlayerStats")
	if !ok {
		t.Fatal("findResultSet(sets, \"PlayerStats\") returned ok=false, want true")
	}
	if got.Headers[0] != "player-data" {
		t.Errorf("findResultSet(sets, \"PlayerStats\") = %+v, want the entry actually named PlayerStats regardless of its position in the slice", got)
	}
}

func TestValidateHeaders(t *testing.T) {
	tests := []struct {
		name    string
		got     []string
		want    []string
		wantErr bool
	}{
		{name: "exact match", got: []string{"GAME_ID", "TEAM_ID"}, want: []string{"GAME_ID", "TEAM_ID"}},
		{name: "empty both", got: nil, want: nil},
		{name: "wrong length", got: []string{"GAME_ID"}, want: []string{"GAME_ID", "TEAM_ID"}, wantErr: true},
		{name: "reordered columns", got: []string{"TEAM_ID", "GAME_ID"}, want: []string{"GAME_ID", "TEAM_ID"}, wantErr: true},
		{name: "renamed column", got: []string{"GAME_ID", "SQUAD_ID"}, want: []string{"GAME_ID", "TEAM_ID"}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateHeaders(tt.got, tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateHeaders(%v, %v) error = %v, wantErr %v", tt.got, tt.want, err, tt.wantErr)
			}
		})
	}
}

func TestJSONTags(t *testing.T) {
	type sample struct {
		GAME_ID string  `json:"GAME_ID"`
		TEAM_ID int     `json:"TEAM_ID"`
		PTS     float64 `json:"PTS"`
	}

	got := jsonTags(sample{})
	want := []string{"GAME_ID", "TEAM_ID", "PTS"}

	if len(got) != len(want) {
		t.Fatalf("jsonTags(sample{}) = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("jsonTags(sample{})[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
