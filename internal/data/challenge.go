package data

import (
	"time"

	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
)

type CreateChallengeRequest struct {
	Type            string    `json:"type"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	EndDate         time.Time `json:"end_date"`
	Invited         []string  `json:"invited"`
	Prize           string    `json:"prize"`
	Penalty         int       `json:"penalty"`
	CigarettesLimit int       `json:"cigarettes_limit"`
}

type CreateChallengeResponse struct {
	ID string `json:"id"`
}

type GetChallengeByIDRequest struct {
	ID string `json:"id"`
}

type GetChallengeByIDResponse struct {
	*models.Challenge
}

type GetChallengesByUserIDRequest struct {
	UserID        string    `json:"user_id"`
	ChallengeType string    `json:"challenge_type"`
	Invite        bool      `json:"invite"`
	Limit         int       `json:"limit"`
	Page          int       `json:"page"`
	From          time.Time `json:"from"`
	To            time.Time `json:"to"`
}

type GetChallengesByUserIDResponse struct {
	Challenges []*models.Challenge `json:"challenges"`
}

type UpdateChallengeRequest struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	EndDate         time.Time `json:"end_date"`
	Invited         []string  `json:"invited"`
	Prize           string    `json:"prize"`
	Penalty         int       `json:"penalty"`
	CigarettesLimit int       `json:"cigarettes_limit"`
}

type UpdateChallengeResponse struct {
	ID string `json:"id"`
}