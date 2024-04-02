package main

import (
	"fmt"
	"testing"

	"github.com/sdomino/scribble"
)

const TestCollection = "todolist_test"

func TestNewApp(t *testing.T) {
	app := NewApp()

	// Verify that the repository is initialized correctly
	if app.Repo.GetCollection() != defaultCollection {
		t.Error("Expected non-nil repository")
	}

	// Verify that the todolist is initialized correctly
	if len(app.Tlist.items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(app.Tlist.items))
	}
}

func TestLoad(t *testing.T) {
	app := NewApp()
	app.Tlist = Todolist{
		items: map[int]string{
			1: "item1",
			2: "item2",
			3: "item3",
		},
		lastKey: 3,
	}

	// Call the Load function
	result := app.GetItems()

	// Verify that the loaded items are correct
	if len(result) != 3 {
		t.Errorf("Expected 3 items, got %d", len(result))
	}
	expectedResults := []string{"", "item1", "item2", "item3"}

	for i, expectedResult := range expectedResults {
		if result[i] != expectedResult {
			t.Errorf("Expected '%s', got '%s'", expectedResult, result[i])
		}
	}
}

func TestAdd(t *testing.T) {
	app := &App{
		Repo: NewRepo(defaultStore, TestCollection),
		Tlist: Todolist{
			items:   make(map[int]string),
			lastKey: 0,
		},
	}

	// Add an item
	app.Add("item1")

	// Verify that the item was added correctly
	if len(app.Tlist.items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(app.Tlist.items))
	}

	// Add another item
	app.Add("item2")

	// Verify that the second item was added correctly
	if len(app.Tlist.items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(app.Tlist.items))
	}
}
func TestRemove(t *testing.T) {
	app := &App{
		Repo: NewRepo(defaultStore, TestCollection),
		Tlist: Todolist{
			items:   make(map[int]string),
			lastKey: 0,
		},
	}

	// Add items to the list
	app.Add("item1")
	app.Add("item2")

	// Remove an item
	app.Remove(1)

	// Verify that the item was removed correctly
	if len(app.Tlist.items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(app.Tlist.items))
	}

	// Remove another item
	app.Remove(2)

	// Verify that the second item was removed correctly
	if len(app.Tlist.items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(app.Tlist.items))
	}
}

func TestReset(t *testing.T) {
	app := &App{
		Repo: NewRepo(defaultStore, TestCollection),
		Tlist: Todolist{
			items:   make(map[int]string),
			lastKey: 0,
		},
	}

	// Add items to the list
	app.Add("item1")
	app.Add("item2")

	// Reset the list
	app.Reset()

	// Verify that the list is empty
	if len(app.Tlist.items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(app.Tlist.items))
	}
}

func TestMaxKey(t *testing.T) {
	type testCase struct {
		items          map[int]string
		expectedResult int
	}

	testCases := []testCase{
		{
			items:          make(map[int]string),
			expectedResult: 0,
		},
		{
			items: map[int]string{
				1: "item1",
				2: "item2",
				3: "item3",
			},
			expectedResult: 3,
		},
		{
			items: map[int]string{
				-1: "item1",
				-2: "item2",
				-3: "item3",
			},
			expectedResult: -1,
		},
		{
			items: map[int]string{
				-5: "item1",
				2:  "item2",
				0:  "item3",
				7:  "item4",
				-3: "item5",
			},
			expectedResult: 7,
		},
	}

	for _, tc := range testCases {
		result := maxKey(tc.items)
		if result != tc.expectedResult {
			t.Errorf("Expected %d, got %d", tc.expectedResult, result)
		}
	}
}
func TestGetDatastorePath(t *testing.T) {
	app := &App{
		Repo: NewRepo(defaultStore, TestCollection),
	}

	expectedResult := defaultStore
	result := app.GetDatastorePath()

	if result != expectedResult {
		t.Errorf("Expected '%s', got '%s'", expectedResult, result)
	}
}
func TestChangeDatastorePath(t *testing.T) {
	app := &App{
		Repo: NewRepo(defaultStore, TestCollection),
		Tlist: Todolist{
			items:   make(map[int]string),
			lastKey: 0,
		},
	}

	expected := defaultStore
	actual := app.GetDatastorePath()

	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}

	// Change the datastore path
	expected = "/tmp"
	actual = app.ChangeDatastorePath(expected)

	// Verify that the datastore path is updated correctly
	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}

type MockRepository struct {
	datastorePath string
	collection    string
	driver        *scribble.Driver
}

func (m *MockRepository) GetAll() map[int]string {
	return map[int]string{
		1: "item1",
		2: "item2",
		3: "item3",
	}
}

func (m *MockRepository) Save(items map[int]string) {
	// Do nothing
}

func (m *MockRepository) GetCollection() string {
	return m.collection
}

func (m *MockRepository) GetDatastorePath() string {
	return m.datastorePath
}

func (m *MockRepository) GetDriver() *scribble.Driver {
	return m.driver
}

func TestLoadDatastore(t *testing.T) {
	db, err := scribble.New(defaultStore, nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	var repo = MockRepository{
		datastorePath: defaultStore,
		collection:    TestCollection,
		driver:        db,
	}
	app := &App{
		Repo: &repo,
		Tlist: Todolist{
			items:   make(map[int]string),
			lastKey: 0,
		},
	}

	// Call the LoadDatastore function
	app.LoadDatastore()

	// Verify that the items are loaded correctly
	if len(app.Tlist.items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(app.Tlist.items))
	}

	expectedResults := []string{"", "item1", "item2", "item3"}

	for i, expectedResult := range expectedResults {
		if app.Tlist.items[i] != expectedResult {
			t.Errorf("Expected '%s', got '%s'", expectedResult, app.Tlist.items[i])
		}
	}

	// Verify that the last key is updated correctly
	if app.Tlist.lastKey != 3 {
		t.Errorf("Expected last key to be 3, got %d", app.Tlist.lastKey)
	}
}
