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

type todolist struct {
	items   map[int]string
	lastKey int
}

var todo = todolist{
	items:   make(map[int]string),
	lastKey: 0,
}

// Reset clears the todolist
func (a *App) Reset() map[int]string {
	todo = todolist{
		items:   make(map[int]string),
		lastKey: 0,
	}
	return todo.items
}

// Add adds an item to the todolist
func (a *App) Add(item string) map[int]string {
	todo.lastKey++
	todo.items[todo.lastKey] = item
	return todo.items
}

func (a *App) Remove(key int) map[int]string {
	delete(todo.items, key)
	return todo.items
}
