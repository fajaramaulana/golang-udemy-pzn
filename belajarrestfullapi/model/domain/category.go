package domain

import "time"

type Category struct {
	Id        int
	Name      string
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}
