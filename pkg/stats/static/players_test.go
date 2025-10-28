package static

import (
	"testing"
)

func TestGetAllPlayers(t *testing.T) {
	players, err := GetAllPlayers()
	if err != nil {
		t.Fatalf("GetAllPlayers() error = %v", err)
	}

	if len(players) == 0 {
		t.Error("GetAllPlayers() returned empty slice")
	}

	if len(players) < 5000 {
		t.Errorf("GetAllPlayers() returned %d players, expected at least 5000", len(players))
	}
}

func TestGetActivePlayers(t *testing.T) {
	players, err := GetActivePlayers()
	if err != nil {
		t.Fatalf("GetActivePlayers() error = %v", err)
	}

	for _, player := range players {
		if !player.IsActive {
			t.Errorf("GetActivePlayers() returned inactive player: %s", player.FullName)
		}
	}
}

func TestGetInactivePlayers(t *testing.T) {
	players, err := GetInactivePlayers()
	if err != nil {
		t.Fatalf("GetInactivePlayers() error = %v", err)
	}

	for _, player := range players {
		if player.IsActive {
			t.Errorf("GetInactivePlayers() returned active player: %s", player.FullName)
		}
	}
}

func TestFindPlayerByID(t *testing.T) {
	tests := []struct {
		name     string
		playerID int
		wantName string
		wantNil  bool
	}{
		{
			name:     "Kareem Abdul-Jabbar",
			playerID: 76003,
			wantName: "Kareem Abdul-Jabbar",
			wantNil:  false,
		},
		{
			name:     "non-existent player",
			playerID: 999999999,
			wantName: "",
			wantNil:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player, err := FindPlayerByID(tt.playerID)
			if err != nil {
				t.Fatalf("FindPlayerByID() error = %v", err)
			}

			if tt.wantNil {
				if player != nil {
					t.Errorf("FindPlayerByID() = %v, want nil", player)
				}
			} else {
				if player == nil {
					t.Fatal("FindPlayerByID() returned nil, want player")
				}
				if player.FullName != tt.wantName {
					t.Errorf("FindPlayerByID() name = %s, want %s", player.FullName, tt.wantName)
				}
			}
		})
	}
}

func TestFindPlayersByFullName(t *testing.T) {
	tests := []struct {
		name      string
		query     string
		wantCount int
		wantNames []string
	}{
		{
			name:      "exact match",
			query:     "^Kareem Abdul-Jabbar$",
			wantCount: 1,
			wantNames: []string{"Kareem Abdul-Jabbar"},
		},
		{
			name:      "partial match",
			query:     "Abdul",
			wantCount: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			players, err := FindPlayersByFullName(tt.query)
			if err != nil {
				t.Fatalf("FindPlayersByFullName() error = %v", err)
			}

			if tt.wantCount > 0 && len(players) < tt.wantCount {
				t.Errorf("FindPlayersByFullName() returned %d players, want at least %d", len(players), tt.wantCount)
			}

			if len(tt.wantNames) > 0 {
				found := make(map[string]bool)
				for _, player := range players {
					found[player.FullName] = true
				}

				for _, wantName := range tt.wantNames {
					if !found[wantName] {
						t.Errorf("FindPlayersByFullName() missing expected player: %s", wantName)
					}
				}
			}
		})
	}
}

func TestSearchPlayers(t *testing.T) {
	tests := []struct {
		name      string
		query     string
		wantFound bool
	}{
		{
			name:      "search by full name",
			query:     "kareem",
			wantFound: true,
		},
		{
			name:      "search by first name",
			query:     "michael",
			wantFound: true,
		},
		{
			name:      "search by last name",
			query:     "jordan",
			wantFound: true,
		},
		{
			name:      "case insensitive",
			query:     "LEBRON",
			wantFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			players, err := SearchPlayers(tt.query)
			if err != nil {
				t.Fatalf("SearchPlayers() error = %v", err)
			}

			if tt.wantFound && len(players) == 0 {
				t.Errorf("SearchPlayers(%s) returned no results, expected at least one", tt.query)
			}
		})
	}
}

func TestStripAccents(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "no accents",
			input: "Test",
			want:  "Test",
		},
		{
			name:  "french accents",
			input: "François",
			want:  "Francois",
		},
		{
			name:  "spanish accents",
			input: "José",
			want:  "Jose",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stripAccents(tt.input)
			if got != tt.want {
				t.Errorf("stripAccents() = %v, want %v", got, tt.want)
			}
		})
	}
}
