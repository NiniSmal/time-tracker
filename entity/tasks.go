package entity

import "time"

type Task struct {
	ID         int64         `json:"id"`
	Name       string        `json:"name"`
	Status     bool          `json:"status"`
	CreatedAt  time.Time     `json:"created_at"`
	FinishedAt time.Time     `json:"finished_at,omitempty"`
	LeadTime   time.Duration `json:"lead_time,omitempty"`
	UserID     int64         `json:"user_id"`
}

type TasksFilter struct {
	Status     bool          `json:"status,omitempty"`
	CreatedAt  time.Time     `json:"created_at,omitempty"`
	FinishedAt time.Time     `json:"finished_at,omitempty"`
	LeadTime   time.Duration `json:"lead_time"`
	UserID     int64         `json:"user_id"`
}
