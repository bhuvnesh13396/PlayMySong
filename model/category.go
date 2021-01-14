package model

import (
	"errors"
)

var (
	CategoryNotFound = errors.New("Category not found")
)

type Category struct {
	ID      		string   `json:"id"`
	Title   		string   `json:"title"`
	Type	      string	`json:"type"`
	SongIDs 		[]string `json:"songs_ids"`
}

type CategoryRepo interface {
	Add(category Category) (err error)
	Get(ID string) (category Category, err error)
	List() (category []Category, err error)
	Update(ID string, songIDs []string) (err error)
}
