package endpoints

import (
	"testing"
)

func BenchmarkToInt(b *testing.B) {
	tests := []interface{}{
		float64(42),
		int(42),
		"42",
		nil,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			_ = toInt(test)
		}
	}
}

func BenchmarkToFloat(b *testing.B) {
	tests := []interface{}{
		float64(42.5),
		int(42),
		"42.5",
		nil,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			_ = toFloat(test)
		}
	}
}

func BenchmarkToString(b *testing.B) {
	tests := []interface{}{
		"hello",
		float64(42.5),
		int(42),
		nil,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			_ = toString(test)
		}
	}
}

func BenchmarkParseSeasonStats(b *testing.B) {
	rows := [][]interface{}{
		{
			203999, "2023-24", "00", 1610612743, "DEN", 25, 82, 82,
			34.5, 9.2, 16.5, 0.558, 1.5, 4.2, 0.357, 5.8, 7.1, 0.817,
			2.9, 9.6, 12.5, 9.0, 1.4, 0.9, 3.0, 2.5, 26.4,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parseSeasonStats(rows)
	}
}

func BenchmarkParseCareerTotals(b *testing.B) {
	rows := [][]interface{}{
		{
			203999, "00", 0, 750, 740,
			26000.0, 6900.0, 12300.0, 0.561, 1100.0, 3100.0, 0.355,
			4350.0, 5400.0, 0.806, 2200.0, 7200.0, 9400.0, 6750.0, 1050.0, 675.0, 2250.0, 1875.0, 19800.0,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parseCareerTotals(rows)
	}
}

func BenchmarkParseGameLogs(b *testing.B) {
	rows := [][]interface{}{
		{
			"22023", 203999, "0022300001", "2023-10-24", "DEN vs. LAL", "W",
			35, 9, 16, 0.563, 2, 5, 0.400, 7, 8, 0.875,
			3, 10, 13, 8, 2, 1, 3, 2, 27, 5, 1,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parseGameLogs(rows)
	}
}
