package main

import (
	"context"
	"fmt"

	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var cnt int

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

func (a *App) domReady(ctx context.Context) {
	go countUp(ctx)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Greet2(name string) string {
	return fmt.Sprintf("Goodbye %s, It's show time!", name)
}

func (a *App) GetCount() int {
	return cnt
}

func countUp(ctx context.Context) {
	for {
		cnt++
		runtime.EventsEmit(ctx, "EmitCount", cnt)
		time.Sleep(time.Second)
	}
}
