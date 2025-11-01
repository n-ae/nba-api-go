package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// LeagueStandingsV3Request contains parameters for the LeagueStandingsV3 endpoint
type LeagueStandingsV3Request struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	LeagueID   *parameters.LeagueID
}

// LeagueStandingsV3Standings represents the Standings result set for LeagueStandingsV3
type LeagueStandingsV3Standings struct {
	TeamID                  string  `json:"TeamID"`
	TeamCity                string  `json:"TeamCity"`
	TeamName                string  `json:"TeamName"`
	Conference              string  `json:"Conference"`
	ConferenceRecord        string  `json:"ConferenceRecord"`
	PlayoffRank             int     `json:"PlayoffRank"`
	ClinchIndicator         string  `json:"ClinchIndicator"`
	DivisionRecord          string  `json:"DivisionRecord"`
	DivisionRank            int     `json:"DivisionRank"`
	WINS                    string  `json:"WINS"`
	LOSSES                  string  `json:"LOSSES"`
	WinPCT                  string  `json:"WinPCT"`
	LeagueRank              int     `json:"LeagueRank"`
	Record                  string  `json:"Record"`
	HOME                    string  `json:"HOME"`
	ROAD                    string  `json:"ROAD"`
	L10                     string  `json:"L10"`
	Last10Home              float64 `json:"Last10Home"`
	Last10Road              float64 `json:"Last10Road"`
	OT                      string  `json:"OT"`
	ThreePTSOrLess          float64 `json:"ThreePTSOrLess"`
	TenPTSOrMore            float64 `json:"TenPTSOrMore"`
	LongHomeStreak          string  `json:"LongHomeStreak"`
	strLongHomeStreak       string  `json:"strLongHomeStreak"`
	LongRoadStreak          string  `json:"LongRoadStreak"`
	strLongRoadStreak       string  `json:"strLongRoadStreak"`
	LongWinStreak           string  `json:"LongWinStreak"`
	LongLossStreak          string  `json:"LongLossStreak"`
	CurrentHomeStreak       string  `json:"CurrentHomeStreak"`
	strCurrentHomeStreak    string  `json:"strCurrentHomeStreak"`
	CurrentRoadStreak       string  `json:"CurrentRoadStreak"`
	strCurrentRoadStreak    string  `json:"strCurrentRoadStreak"`
	CurrentStreak           string  `json:"CurrentStreak"`
	strCurrentStreak        string  `json:"strCurrentStreak"`
	ConferenceGamesBack     string  `json:"ConferenceGamesBack"`
	ClinchedConferenceTitle string  `json:"ClinchedConferenceTitle"`
	ClinchedDivisionTitle   string  `json:"ClinchedDivisionTitle"`
	ClinchedPlayoffBirth    string  `json:"ClinchedPlayoffBirth"`
	EliminatedConference    float64 `json:"EliminatedConference"`
	EliminatedDivision      float64 `json:"EliminatedDivision"`
	AheadAtHalf             string  `json:"AheadAtHalf"`
	BehindAtHalf            string  `json:"BehindAtHalf"`
	TiedAtHalf              string  `json:"TiedAtHalf"`
	AheadAtThird            string  `json:"AheadAtThird"`
	BehindAtThird           string  `json:"BehindAtThird"`
	TiedAtThird             string  `json:"TiedAtThird"`
	Score100PTS             float64 `json:"Score100PTS"`
	OppScore100PTS          float64 `json:"OppScore100PTS"`
	OppOver500              string  `json:"OppOver500"`
	LeadInFGPCT             float64 `json:"LeadInFGPCT"`
	LeadInReb               float64 `json:"LeadInReb"`
	FewerTurnovers          string  `json:"FewerTurnovers"`
	PointsPG                string  `json:"PointsPG"`
	OppPointsPG             string  `json:"OppPointsPG"`
	DiffPointsPG            string  `json:"DiffPointsPG"`
	vsEast                  float64 `json:"vsEast"`
	vsAtlantic              string  `json:"vsAtlantic"`
	vsCentral               string  `json:"vsCentral"`
	vsSoutheast             float64 `json:"vsSoutheast"`
	vsWest                  string  `json:"vsWest"`
	vsNorthwest             string  `json:"vsNorthwest"`
	vsPacific               string  `json:"vsPacific"`
	vsSouthwest             string  `json:"vsSouthwest"`
	Jan                     string  `json:"Jan"`
	Feb                     string  `json:"Feb"`
	Mar                     string  `json:"Mar"`
	Apr                     string  `json:"Apr"`
	May                     string  `json:"May"`
	Jun                     string  `json:"Jun"`
	Jul                     string  `json:"Jul"`
	Aug                     string  `json:"Aug"`
	Sep                     string  `json:"Sep"`
	Oct                     string  `json:"Oct"`
	Nov                     string  `json:"Nov"`
	Dec                     string  `json:"Dec"`
}

// LeagueStandingsV3Response contains the response data from the LeagueStandingsV3 endpoint
type LeagueStandingsV3Response struct {
	Standings []LeagueStandingsV3Standings
}

// GetLeagueStandingsV3 retrieves data from the leaguestandingsv3 endpoint
func GetLeagueStandingsV3(ctx context.Context, client *stats.Client, req LeagueStandingsV3Request) (*models.Response[*LeagueStandingsV3Response], error) {
	params := url.Values{}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/leaguestandingsv3", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueStandingsV3Response{}
	if len(rawResp.ResultSets) > 0 {
		response.Standings = make([]LeagueStandingsV3Standings, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 75 {
				item := LeagueStandingsV3Standings{
					TeamID:                  toString(row[0]),
					TeamCity:                toString(row[1]),
					TeamName:                toString(row[2]),
					Conference:              toString(row[3]),
					ConferenceRecord:        toString(row[4]),
					PlayoffRank:             toInt(row[5]),
					ClinchIndicator:         toString(row[6]),
					DivisionRecord:          toString(row[7]),
					DivisionRank:            toInt(row[8]),
					WINS:                    toString(row[9]),
					LOSSES:                  toString(row[10]),
					WinPCT:                  toString(row[11]),
					LeagueRank:              toInt(row[12]),
					Record:                  toString(row[13]),
					HOME:                    toString(row[14]),
					ROAD:                    toString(row[15]),
					L10:                     toString(row[16]),
					Last10Home:              toFloat(row[17]),
					Last10Road:              toFloat(row[18]),
					OT:                      toString(row[19]),
					ThreePTSOrLess:          toFloat(row[20]),
					TenPTSOrMore:            toFloat(row[21]),
					LongHomeStreak:          toString(row[22]),
					strLongHomeStreak:       toString(row[23]),
					LongRoadStreak:          toString(row[24]),
					strLongRoadStreak:       toString(row[25]),
					LongWinStreak:           toString(row[26]),
					LongLossStreak:          toString(row[27]),
					CurrentHomeStreak:       toString(row[28]),
					strCurrentHomeStreak:    toString(row[29]),
					CurrentRoadStreak:       toString(row[30]),
					strCurrentRoadStreak:    toString(row[31]),
					CurrentStreak:           toString(row[32]),
					strCurrentStreak:        toString(row[33]),
					ConferenceGamesBack:     toString(row[34]),
					ClinchedConferenceTitle: toString(row[35]),
					ClinchedDivisionTitle:   toString(row[36]),
					ClinchedPlayoffBirth:    toString(row[37]),
					EliminatedConference:    toFloat(row[38]),
					EliminatedDivision:      toFloat(row[39]),
					AheadAtHalf:             toString(row[40]),
					BehindAtHalf:            toString(row[41]),
					TiedAtHalf:              toString(row[42]),
					AheadAtThird:            toString(row[43]),
					BehindAtThird:           toString(row[44]),
					TiedAtThird:             toString(row[45]),
					Score100PTS:             toFloat(row[46]),
					OppScore100PTS:          toFloat(row[47]),
					OppOver500:              toString(row[48]),
					LeadInFGPCT:             toFloat(row[49]),
					LeadInReb:               toFloat(row[50]),
					FewerTurnovers:          toString(row[51]),
					PointsPG:                toString(row[52]),
					OppPointsPG:             toString(row[53]),
					DiffPointsPG:            toString(row[54]),
					vsEast:                  toFloat(row[55]),
					vsAtlantic:              toString(row[56]),
					vsCentral:               toString(row[57]),
					vsSoutheast:             toFloat(row[58]),
					vsWest:                  toString(row[59]),
					vsNorthwest:             toString(row[60]),
					vsPacific:               toString(row[61]),
					vsSouthwest:             toString(row[62]),
					Jan:                     toString(row[63]),
					Feb:                     toString(row[64]),
					Mar:                     toString(row[65]),
					Apr:                     toString(row[66]),
					May:                     toString(row[67]),
					Jun:                     toString(row[68]),
					Jul:                     toString(row[69]),
					Aug:                     toString(row[70]),
					Sep:                     toString(row[71]),
					Oct:                     toString(row[72]),
					Nov:                     toString(row[73]),
					Dec:                     toString(row[74]),
				}
				response.Standings = append(response.Standings, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
