package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
)

// VideoEventsRequest contains parameters for the VideoEvents endpoint
type VideoEventsRequest struct {
	GameID      string
	GameEventID string
}

// VideoEventsVideo represents the Video result set for VideoEvents
type VideoEventsVideo struct {
	uuid string `json:"uuid"`
	vl   string `json:"vl"`
	vt   string `json:"vt"`
	gc   string `json:"gc"`
	surl string `json:"surl"`
	durl string `json:"durl"`
	vurl string `json:"vurl"`
	purl string `json:"purl"`
}

// VideoEventsResponse contains the response data from the VideoEvents endpoint
type VideoEventsResponse struct {
	Video []VideoEventsVideo
}

// GetVideoEvents retrieves data from the videoevents endpoint
func GetVideoEvents(ctx context.Context, client *stats.Client, req VideoEventsRequest) (*models.Response[*VideoEventsResponse], error) {
	params := url.Values{}
	if req.GameID == "" {
		return nil, fmt.Errorf("GameID is required")
	}
	params.Set("GameID", string(req.GameID))
	if req.GameEventID == "" {
		return nil, fmt.Errorf("GameEventID is required")
	}
	params.Set("GameEventID", string(req.GameEventID))

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "videoevents", params, &rawResp); err != nil {
		return nil, err
	}

	response := &VideoEventsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.Video = make([]VideoEventsVideo, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 8 {
				item := VideoEventsVideo{
					uuid: toString(row[0]),
					vl:   toString(row[1]),
					vt:   toString(row[2]),
					gc:   toString(row[3]),
					surl: toString(row[4]),
					durl: toString(row[5]),
					vurl: toString(row[6]),
					purl: toString(row[7]),
				}
				response.Video = append(response.Video, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
