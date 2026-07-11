package domain

import "time"

type RoadMapItems struct {
	ID        int
	ProjectID int
	Title     string
	Goal      string
	Position  int
	Status    ItemsStatus
	Deadline  time.Time
	CreatedBY int
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type ItemsStatus string

const (
	ItemsStatusCreated    ItemsStatus = "created"
	ItemsStatusInProgress ItemsStatus = "in_progress"
	ItemsStatusFinished   ItemsStatus = "finished"
	ItemsStatusDelayed    ItemsStatus = "delayed"
	ItemsStatusExtended   ItemsStatus = "extended"
)
