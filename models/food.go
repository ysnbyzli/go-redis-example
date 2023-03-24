package models

import "github.com/google/uuid"

type Food struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float32   `json:"price"`
}
