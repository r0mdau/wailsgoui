package main

import (
	"fmt"

	"github.com/sdomino/scribble"
)

// Repository struct
type Repository struct {
	driver     *scribble.Driver
	collection string
}

// NewRepo creates a new Repository from path and collection
func NewRepo(path, collection string) Repository {
	db, err := scribble.New(path, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	repo := Repository{
		driver:     db,
		collection: collection,
	}

	return repo
}

// GetAll reads the entire collection from disk
func (r *Repository) GetAll() map[int]string {
	items := make(map[int]string)
	if err := r.driver.Read(r.collection, "items", &items); err != nil {
		fmt.Println("Error", err)
	}
	return items
}

// Save rewrites the entire collection on disk
func (r *Repository) Save(items map[int]string) {
	if err := r.driver.Write(r.collection, "items", items); err != nil {
		fmt.Println("Error", err)
	}
}
