package static

import (
	"testing"
)

func BenchmarkGetAllPlayers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetAllPlayers()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetActivePlayers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetActivePlayers()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFindPlayerByID(b *testing.B) {
	playerID := 203999

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := FindPlayerByID(playerID)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFindPlayersByFullName(b *testing.B) {
	query := "LeBron"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := FindPlayersByFullName(query)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSearchPlayers(b *testing.B) {
	query := "jordan"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := SearchPlayers(query)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStripAccents(b *testing.B) {
	text := "José François Müller Øyvind"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = stripAccents(text)
	}
}

func BenchmarkFindTeamByID(b *testing.B) {
	teamID := 1610612747

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := FindTeamByID(teamID)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSearchTeams(b *testing.B) {
	query := "lakers"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := SearchTeams(query)
		if err != nil {
			b.Fatal(err)
		}
	}
}
