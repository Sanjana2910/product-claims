package models

import "time"

// workflow tracks the claim identification job
type Workflow struct {
	ID        string    `json:"workflowId"`
	ProductID string    `json:"productId"`
	Status    string    `json:"status"`
	ErrorMsg  string    `json:"errorMsg,omitempty"`
	Product   *Product  `json:"product,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// workflow status values
const (
	StatusInProgress = "IN_PROGRESS"
	StatusCompleted  = "COMPLETED"
	StatusFailed     = "FAILED"
)