package models

type Task struct {
	ID            string   `json:"id" firestore:"-"`
	Title         string   `json:"title" firestore:"title"`
	Description   string   `json:"description" firestore:"description"`
	Points        int      `json:"points" firestore:"points"`
	RequiresPhoto bool     `json:"requires_photo" firestore:"requires_photo"`
	CompletedBy   []string `json:"completed_by" firestore:"completed_by"`
}
