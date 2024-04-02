package main

import (
	"fmt"

	"github.com/sdomino/scribble"
)

type Repository interface {
	GetAll() map[int]string
	Save(map[int]string)
	GetDatastorePath() string
	GetCollection() string
	GetDriver() *scribble.Driver
}

type repository struct {
	datastorePath string
	driver        *scribble.Driver
	collection    string
}

// NewRepo creates a new Repository from path and collection
func NewRepo(path, collection string) Repository {
	db, err := scribble.New(path, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	repo := &repository{
		datastorePath: path,
		driver:        db,
		collection:    collection,
	}

	return repo
}

// GetAll reads the entire collection from disk
func (r *repository) GetAll() map[int]string {
	items := make(map[int]string)
	if err := r.driver.Read(r.collection, "items", &items); err != nil {
		fmt.Println("Error", err)
	}
	return items
}

// Save rewrites the entire collection on disk
func (r *repository) Save(items map[int]string) {
	if err := r.driver.Write(r.collection, "items", items); err != nil {
		fmt.Println("Error", err)
	}
}

// GetDatastorePath returns the path to the datastore
func (r *repository) GetDatastorePath() string {
	return r.datastorePath
}

// ChangeDatastorePath changes the path to the datastore
func (r *repository) GetCollection() string {
	return r.collection
}

// GetDriver returns the scribble driver
func (r *repository) GetDriver() *scribble.Driver {
	return r.driver
}
