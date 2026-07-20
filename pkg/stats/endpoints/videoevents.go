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
	UUID string `json:"uuid"`
	VL   string `json:"vl"`
	VT   string `json:"vt"`
	GC   string `json:"gc"`
	SURL string `json:"surl"`
	DURL string `json:"durl"`
	VURL string `json:"vurl"`
	PURL string `json:"purl"`
}

// VideoEventsResponse contains the response data from the VideoEvents endpoint
type VideoEventsResponse struct {
	Video []VideoEventsVideo
}

// GetVideoEvents retrieves data from the videoevents endpoint
func GetVideoEvents(ctx context.Context, client *stats.Client, req VideoEventsRequest) (*models.Response[*VideoEventsResponse], error) {
	params := url.Values{}
	if req.GameID == "" {
		return nil, fmt.Errorf("gameID is required")
	}
	params.Set("GameID", req.GameID)
	if req.GameEventID == "" {
		return nil, fmt.Errorf("gameEventID is required")
	}
	params.Set("GameEventID", req.GameEventID)

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "videoevents", params, &rawResp); err != nil {
		return nil, err
	}

	response := &VideoEventsResponse{}
	if rs, ok := findResultSet(rawResp.ResultSets, "Video"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(VideoEventsVideo{})); err != nil {
			return nil, fmt.Errorf("VideoEvents: Video result set: %w", err)
		}
		response.Video = make([]VideoEventsVideo, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 8 {
				item := VideoEventsVideo{
					UUID: toString(row[0]),
					VL:   toString(row[1]),
					VT:   toString(row[2]),
					GC:   toString(row[3]),
					SURL: toString(row[4]),
					DURL: toString(row[5]),
					VURL: toString(row[6]),
					PURL: toString(row[7]),
				}
				response.Video = append(response.Video, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
