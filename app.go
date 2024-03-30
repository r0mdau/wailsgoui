package main

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

var items []string

// Reset clears the todolist
func (a *App) Reset() []string {
	items = []string{}
	return items
}

// Add adds an item to the todolist
func (a *App) Add(item string) []string {
	items = append(items, item)
	return items
}
