package parameters

import "fmt"

type PerMode string

const (
	PerModeTotals        PerMode = "Totals"
	PerModePerGame       PerMode = "PerGame"
	PerModePer36         PerMode = "Per36"
	PerModePerMinute     PerMode = "PerMinute"
	PerModePer48         PerMode = "Per48"
	PerModePer40         PerMode = "Per40"
	PerModePerPossession PerMode = "PerPossession"
	PerModePer100Plays   PerMode = "Per100Plays"
	PerModePer100Poss    PerMode = "Per100Possessions"
)

func (p PerMode) Validate() error {
	switch p {
	case PerModeTotals, PerModePerGame, PerModePer36, PerModePerMinute,
		PerModePer48, PerModePer40, PerModePerPossession, PerModePer100Plays, PerModePer100Poss:
		return nil
	default:
		return fmt.Errorf("invalid PerMode: %s", p)
	}
}

func (p PerMode) String() string {
	return string(p)
}

type LeagueID string

const (
	LeagueIDNBA     LeagueID = "00"
	LeagueIDABA     LeagueID = "01"
	LeagueIDGLeague LeagueID = "20"
)

func (l LeagueID) Validate() error {
	switch l {
	case LeagueIDNBA, LeagueIDABA, LeagueIDGLeague, "":
		return nil
	default:
		return fmt.Errorf("invalid LeagueID: %s", l)
	}
}

func (l LeagueID) String() string {
	return string(l)
}

type Season string

const (
	SeasonAllTime Season = ""
)

func NewSeason(year int) Season {
	return Season(fmt.Sprintf("%d-%02d", year, (year+1)%100))
}

func (s Season) Validate() error {
	if s == "" {
		return nil
	}
	return nil
}

func (s Season) String() string {
	return string(s)
}

type SeasonType string

const (
	SeasonTypeRegular   SeasonType = "Regular Season"
	SeasonTypePlayoffs  SeasonType = "Playoffs"
	SeasonTypeAllStar   SeasonType = "All Star"
	SeasonTypePreseason SeasonType = "Pre Season"
)

func (s SeasonType) Validate() error {
	switch s {
	case SeasonTypeRegular, SeasonTypePlayoffs, SeasonTypeAllStar, SeasonTypePreseason, "":
		return nil
	default:
		return fmt.Errorf("invalid SeasonType: %s", s)
	}
}

func (s SeasonType) String() string {
	return string(s)
}

type StatCategory string

const (
	StatCategoryPoints    StatCategory = "PTS"
	StatCategoryRebounds  StatCategory = "REB"
	StatCategoryAssists   StatCategory = "AST"
	StatCategoryBlocks    StatCategory = "BLK"
	StatCategorySteals    StatCategory = "STL"
	StatCategoryTurnovers StatCategory = "TOV"
	StatCategoryFGPct     StatCategory = "FG_PCT"
	StatCategoryFG3Pct    StatCategory = "FG3_PCT"
	StatCategoryFTPct     StatCategory = "FT_PCT"
)

func (s StatCategory) Validate() error {
	switch s {
	case StatCategoryPoints, StatCategoryRebounds, StatCategoryAssists,
		StatCategoryBlocks, StatCategorySteals, StatCategoryTurnovers,
		StatCategoryFGPct, StatCategoryFG3Pct, StatCategoryFTPct, "":
		return nil
	default:
		return fmt.Errorf("invalid StatCategory: %s", s)
	}
}

func (s StatCategory) String() string {
	return string(s)
}

type MeasureType string

const (
	MeasureTypeBase     MeasureType = "Base"
	MeasureTypeAdvanced MeasureType = "Advanced"
	MeasureTypeMisc     MeasureType = "Misc"
	MeasureTypeScoring  MeasureType = "Scoring"
	MeasureTypeUsage    MeasureType = "Usage"
)

func (m MeasureType) Validate() error {
	switch m {
	case MeasureTypeBase, MeasureTypeAdvanced, MeasureTypeMisc,
		MeasureTypeScoring, MeasureTypeUsage, "":
		return nil
	default:
		return fmt.Errorf("invalid MeasureType: %s", m)
	}
}

func (m MeasureType) String() string {
	return string(m)
}

type PlayerOrTeam string

const (
	PlayerOrTeamPlayer PlayerOrTeam = "Player"
	PlayerOrTeamTeam   PlayerOrTeam = "Team"
)

func (p PlayerOrTeam) Validate() error {
	switch p {
	case PlayerOrTeamPlayer, PlayerOrTeamTeam:
		return nil
	default:
		return fmt.Errorf("invalid PlayerOrTeam: %s", p)
	}
}

func (p PlayerOrTeam) String() string {
	return string(p)
}
