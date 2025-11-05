package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// LeagueStandingsRequest contains parameters for the LeagueStandings endpoint
type LeagueStandingsRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	LeagueID   *parameters.LeagueID
}

// LeagueStandingsStandings represents the Standings result set for LeagueStandings
type LeagueStandingsStandings struct {
	LeagueID                string  `json:"LeagueID"`
	SeasonID                string  `json:"SeasonID"`
	TeamID                  string  `json:"TeamID"`
	TeamCity                string  `json:"TeamCity"`
	TeamName                string  `json:"TeamName"`
	Conference              string  `json:"Conference"`
	ConferenceRecord        string  `json:"ConferenceRecord"`
	PlayoffRank             int     `json:"PlayoffRank"`
	ClinchIndicator         string  `json:"ClinchIndicator"`
	Division                string  `json:"Division"`
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
	DivisionGamesBack       string  `json:"DivisionGamesBack"`
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
	Score_80_Plus           string  `json:"Score_80_Plus"`
	Opp_Score_80_Plus       string  `json:"Opp_Score_80_Plus"`
	Score_Below_80          string  `json:"Score_Below_80"`
	Opp_Score_Below_80      string  `json:"Opp_Score_Below_80"`
	TotalPoints             string  `json:"TotalPoints"`
	OppTotalPoints          string  `json:"OppTotalPoints"`
	DiffTotalPoints         string  `json:"DiffTotalPoints"`
}

// LeagueStandingsResponse contains the response data from the LeagueStandings endpoint
type LeagueStandingsResponse struct {
	Standings []LeagueStandingsStandings
}

// GetLeagueStandings retrieves data from the leaguestandings endpoint
func GetLeagueStandings(ctx context.Context, client *stats.Client, req LeagueStandingsRequest) (*models.Response[*LeagueStandingsResponse], error) {
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
	if err := client.GetJSON(ctx, "leaguestandings", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueStandingsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.Standings = make([]LeagueStandingsStandings, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 86 {
				item := LeagueStandingsStandings{
					LeagueID:                toString(row[0]),
					SeasonID:                toString(row[1]),
					TeamID:                  toString(row[2]),
					TeamCity:                toString(row[3]),
					TeamName:                toString(row[4]),
					Conference:              toString(row[5]),
					ConferenceRecord:        toString(row[6]),
					PlayoffRank:             toInt(row[7]),
					ClinchIndicator:         toString(row[8]),
					Division:                toString(row[9]),
					DivisionRecord:          toString(row[10]),
					DivisionRank:            toInt(row[11]),
					WINS:                    toString(row[12]),
					LOSSES:                  toString(row[13]),
					WinPCT:                  toString(row[14]),
					LeagueRank:              toInt(row[15]),
					Record:                  toString(row[16]),
					HOME:                    toString(row[17]),
					ROAD:                    toString(row[18]),
					L10:                     toString(row[19]),
					Last10Home:              toFloat(row[20]),
					Last10Road:              toFloat(row[21]),
					OT:                      toString(row[22]),
					ThreePTSOrLess:          toFloat(row[23]),
					TenPTSOrMore:            toFloat(row[24]),
					LongHomeStreak:          toString(row[25]),
					strLongHomeStreak:       toString(row[26]),
					LongRoadStreak:          toString(row[27]),
					strLongRoadStreak:       toString(row[28]),
					LongWinStreak:           toString(row[29]),
					LongLossStreak:          toString(row[30]),
					CurrentHomeStreak:       toString(row[31]),
					strCurrentHomeStreak:    toString(row[32]),
					CurrentRoadStreak:       toString(row[33]),
					strCurrentRoadStreak:    toString(row[34]),
					CurrentStreak:           toString(row[35]),
					strCurrentStreak:        toString(row[36]),
					ConferenceGamesBack:     toString(row[37]),
					DivisionGamesBack:       toString(row[38]),
					ClinchedConferenceTitle: toString(row[39]),
					ClinchedDivisionTitle:   toString(row[40]),
					ClinchedPlayoffBirth:    toString(row[41]),
					EliminatedConference:    toFloat(row[42]),
					EliminatedDivision:      toFloat(row[43]),
					AheadAtHalf:             toString(row[44]),
					BehindAtHalf:            toString(row[45]),
					TiedAtHalf:              toString(row[46]),
					AheadAtThird:            toString(row[47]),
					BehindAtThird:           toString(row[48]),
					TiedAtThird:             toString(row[49]),
					Score100PTS:             toFloat(row[50]),
					OppScore100PTS:          toFloat(row[51]),
					OppOver500:              toString(row[52]),
					LeadInFGPCT:             toFloat(row[53]),
					LeadInReb:               toFloat(row[54]),
					FewerTurnovers:          toString(row[55]),
					PointsPG:                toString(row[56]),
					OppPointsPG:             toString(row[57]),
					DiffPointsPG:            toString(row[58]),
					vsEast:                  toFloat(row[59]),
					vsAtlantic:              toString(row[60]),
					vsCentral:               toString(row[61]),
					vsSoutheast:             toFloat(row[62]),
					vsWest:                  toString(row[63]),
					vsNorthwest:             toString(row[64]),
					vsPacific:               toString(row[65]),
					vsSouthwest:             toString(row[66]),
					Jan:                     toString(row[67]),
					Feb:                     toString(row[68]),
					Mar:                     toString(row[69]),
					Apr:                     toString(row[70]),
					May:                     toString(row[71]),
					Jun:                     toString(row[72]),
					Jul:                     toString(row[73]),
					Aug:                     toString(row[74]),
					Sep:                     toString(row[75]),
					Oct:                     toString(row[76]),
					Nov:                     toString(row[77]),
					Dec:                     toString(row[78]),
					Score_80_Plus:           toString(row[79]),
					Opp_Score_80_Plus:       toString(row[80]),
					Score_Below_80:          toString(row[81]),
					Opp_Score_Below_80:      toString(row[82]),
					TotalPoints:             toString(row[83]),
					OppTotalPoints:          toString(row[84]),
					DiffTotalPoints:         toString(row[85]),
				}
				response.Standings = append(response.Standings, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
