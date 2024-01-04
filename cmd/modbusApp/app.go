package main

import (
	"context"
	"encoding/json"
	"fmt"
)

var cnt int

// App struct
type App struct {
	ctx      context.Context
	Hostname string `json:"hostname"`
	Port     uint   `json:"port"`
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

func (a *App) domReady(ctx context.Context) {
}

func (a *App) CheckSubmit(s string) {
	fmt.Println("Submitted", s)
	json.Unmarshal([]byte(s), a)
	fmt.Println(a.Hostname, a.Port)
}
