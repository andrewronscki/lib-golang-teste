package helloworld

import "time"

type Model struct {
	Name      string    `json:"name"`
	SiteID    string    `json:"site_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
