package main

import (
	"context"
	"math"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Todolist struct {
	items   map[int]string
	lastKey int
}

// App struct
type App struct {
	ctx   context.Context
	Repo  Repository
	Tlist Todolist
}

const defaultCollection = "todolist"

var defaultStore = "./data"

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		Repo: &repository{
			datastorePath: defaultStore,
			collection:    defaultCollection,
		},
		Tlist: Todolist{
			items:   make(map[int]string),
			lastKey: 0,
		},
	}
}

// maxKey returns the maximum key value in a map
func maxKey(items map[int]string) int {
	if len(items) == 0 {
		return 0
	}
	maxNumber := math.MinInt
	for n := range items {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.Repo = NewRepo(defaultStore, defaultCollection)
	a.LoadDatastore()
}

// LoadDatastore loads the datastore into memory
func (a *App) LoadDatastore() {
	a.Tlist.items = a.Repo.GetAll()
	a.Tlist.lastKey = maxKey(a.Tlist.items)
}

// GetItems returns the current todolist
func (a *App) GetItems() map[int]string {
	return a.Tlist.items
}

// Add adds an item to the todolist
func (a *App) Add(item string) map[int]string {
	a.Tlist.lastKey++
	a.Tlist.items[a.Tlist.lastKey] = item
	a.Repo.Save(a.Tlist.items)
	return a.Tlist.items
}

// Remove removes an item from the todolist
func (a *App) Remove(key int) map[int]string {
	delete(a.Tlist.items, key)
	a.Repo.Save(a.Tlist.items)
	return a.Tlist.items
}

// Reset clears the todolist
func (a *App) Reset() map[int]string {
	a.Tlist = Todolist{
		items:   make(map[int]string),
		lastKey: 0,
	}
	a.Repo.Save(a.Tlist.items)
	return a.Tlist.items
}

// GetDatastorePath returns the current datastore path
func (a *App) GetDatastorePath() string {
	return a.Repo.GetDatastorePath()
}

// ChangeDatastorePath changes the datastore path and trigger a reload items event
func (a *App) ChangeDatastorePath(path string) string {
	a.Repo = NewRepo(path, defaultCollection)
	a.LoadDatastore()

	// TODO find a more elegant hack for unit tests
	if a.ctx != nil {
		runtime.EventsEmit(a.ctx, "ReloadItems", path)
	}
	return a.Repo.GetDatastorePath()
}
