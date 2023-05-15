package entities

import "time"

type Profile struct {
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
}
