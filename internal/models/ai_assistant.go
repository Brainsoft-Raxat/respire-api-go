package models

import "time"

type GetRecommendationRequest struct {
	EventType string                       `json:"event_type"`
	Data      GetRecommendationRequestData `json:"data"`
}

type GetRecommendationRequestData struct {
	CravingLevel int       `json:"craving_level"`
	Context      string    `json:"context"`
	Mood         string    `json:"mood"`
	Timestamp    time.Time `json:"timestamp"`
}

type GetRecommendationResponse struct {
	Recommendations []string `json:"recommendations"`
}
