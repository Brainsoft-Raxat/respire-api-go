package data

type GetRecommendationsRequest struct {
	CravingLevel int    `json:"craving"`
	Context      string `json:"context"`
	Mood         string `json:"mood"`
}

type GetRecommendationsResponse struct{
	Reccomendations []string `json:"reccomendations"`
}
