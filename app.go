package main

import (
	"context"
	"math"
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

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		Repo: NewData("./data", "todolist"),
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

	a.Tlist.items = a.Repo.GetAll()
	a.Tlist.lastKey = maxKey(a.Tlist.items)
}

// Load returns the current todolist
func (a *App) Load() map[int]string {
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
