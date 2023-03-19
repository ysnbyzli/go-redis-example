package models

import "github.com/google/uuid"

type Todo struct {
	id        uuid.UUID `json:"id"`
	text      string    `json:"text"`
	completed bool      `json:"completed"`
}
