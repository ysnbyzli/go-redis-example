package models

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Food struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float32   `json:"price"`
}

type Basket struct {
	Foods []BasketItem `json:"foods"`
}

type BasketItem struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
	Count int       `json:"count"`
}

var Foods = [5]Food{
	{Id: uuid.New(), Name: "Adana Kebab", Price: 120},
	{Id: uuid.New(), Name: "Mercimek Çorbası", Price: 30},
	{Id: uuid.New(), Name: "İrmik Helvası", Price: 30},
	{Id: uuid.New(), Name: "Ekmek Arası Köfte", Price: 60},
	{Id: uuid.New(), Name: "Ayran", Price: 20},
}

func (b *Food) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &b)
}

func (b *Basket) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &b)
}

func (b *BasketItem) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &b)
}
