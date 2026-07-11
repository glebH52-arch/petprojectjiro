package domain

import "time"

type Project struct {
	ID          int
	CreatedBY   int
	Title       string
	Goal        string
	Description string
	Status      ProjectStatus
	Archived    *bool
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	FinishedAt  *time.Time
}

type ProjectStatus string

const (
	ProjectStatusActive   ProjectStatus = "active"
	ProjectStatusArchived ProjectStatus = "archived"
	ProjectStatusFinished ProjectStatus = "finished"
)
