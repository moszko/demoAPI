package domain

import "time"

type Trademark struct {
	Id         int
	Name       string
	StatusCode string
	StatusDate time.Time
}

type TrademarkRepository interface {
	GetByName(name string) (*Trademark, error)
	GetSimilarByName(name string) ([]string, error)
}
