package models

import "time"

type Challenge struct {
	ID              string    `json:"id" firestore:"-"`
	Type            string    `json:"type" firestore:"type,omitempty"`
	Name            string    `json:"name" firestore:"name,omitempty"`
	Description     string    `json:"description" firestore:"description,omitempty"`
	EndDate         time.Time `json:"end_date" firestore:"end_date,omitempty"`
	OwnerID         string    `json:"owner_id" firestore:"owner_id,omitempty"`
	Participants    []string  `json:"participants" firestore:"participants,omitempty"`
	Invited         []string  `json:"invited" firestore:"invited,omitempty"`
	Prize           string    `json:"prize" firestore:"prize,omitempty"`
	Penalty         int       `json:"penalty" firestore:"penalty,omitempty"`
	CigarettesLimit int       `json:"cigarettes_limit" firestore:"cigarettes_limit,omitempty"`
}

const (
	ChallengeTypeLimitCigarettes       = "limit_cigarettes"
	ChallengeTypeMaximizeSmokeFreeTime = "maximize_smoke_free_time"
	ChallengeTypeTasks                 = "tasks"
)
