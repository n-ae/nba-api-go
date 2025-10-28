package parameters

import "net/url"

type RequestOption func(params url.Values)

func WithLeagueID(leagueID LeagueID) RequestOption {
	return func(params url.Values) {
		if leagueID != "" {
			params.Set("LeagueID", leagueID.String())
		}
	}
}

func WithSeason(season Season) RequestOption {
	return func(params url.Values) {
		if season != "" {
			params.Set("Season", season.String())
		}
	}
}

func WithSeasonType(seasonType SeasonType) RequestOption {
	return func(params url.Values) {
		if seasonType != "" {
			params.Set("SeasonType", seasonType.String())
		}
	}
}

func WithPerMode(perMode PerMode) RequestOption {
	return func(params url.Values) {
		if perMode != "" {
			params.Set("PerMode", perMode.String())
		}
	}
}

func WithMeasureType(measureType MeasureType) RequestOption {
	return func(params url.Values) {
		if measureType != "" {
			params.Set("MeasureType", measureType.String())
		}
	}
}
