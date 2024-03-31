package main

import (
	"reflect"
	"testing"
)

type TSetup struct {
	path       string
	collection string
	repo       Repository
}

var tset = TSetup{}

func setUp() {
	tset.path = "./data"
	tset.collection = "todolist_test"
	tset.repo = NewRepo(tset.path, tset.collection)
}

func tearDown() {
	tset.repo.driver.Delete(tset.collection, "items")
}

func TestNewRepo(t *testing.T) {
	setUp()
	defer tearDown()

	// Verify that the repository driver is not nil
	if tset.repo.driver == nil {
		t.Error("Expected non-nil repository driver")
	}

	// Verify that the repository collection is set correctly
	if tset.repo.collection != tset.collection {
		t.Errorf("Expected collection '%s', got '%s'", tset.collection, tset.repo.collection)
	}
}
func TestRepository_GetAll(t *testing.T) {
	setUp()
	defer tearDown()

	t.Run("Verify GetAll returns all items", func(t *testing.T) {
		expectedItems := map[int]string{1: "item1", 2: "item2", 3: "item3"}
		tset.repo.driver.Write(tset.collection, "items", expectedItems)

		actualItems := tset.repo.GetAll()
		if !reflect.DeepEqual(actualItems, expectedItems) {
			t.Errorf("Expected items %v, got %v", expectedItems, actualItems)
		}
	})

	t.Run("Verify GetAll returns empty map when no items exist", func(t *testing.T) {
		expectedItems := map[int]string{1: "item1", 2: "item2", 3: "item3"}
		tset.repo.driver.Write(tset.collection, "items", expectedItems)

		// Empty the collection
		tset.repo.driver.Write(tset.collection, "items", map[int]string{})
		actualItems := tset.repo.GetAll()
		if len(actualItems) != 0 {
			t.Errorf("Expected empty map, got %v", actualItems)
		}
	})
}
func TestRepository_Save(t *testing.T) {
	setUp()
	defer tearDown()

	// Test case 1: Verify items are saved successfully
	items := map[int]string{1: "item1", 2: "item2", 3: "item3"}
	tset.repo.Save(items)

	// Verify that the saved items are retrieved correctly
	actual := make(map[int]string)
	tset.repo.driver.Read(tset.collection, "items", &actual) // Add nil as the third argument
	if !reflect.DeepEqual(actual, items) {
		t.Errorf("Expected saved items %v, got %v", items, actual)
	}
}
