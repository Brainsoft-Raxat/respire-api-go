package data

type Accepted struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type FilterRequest struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
