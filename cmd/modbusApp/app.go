package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/simonvetter/modbus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type groupDefinition struct {
	StartReg uint `json:"startReg"`
	NumReg   uint `json:"numReg"`
}

// App struct
type App struct {
	ctx        context.Context
	Hostname   string `json:"hostname"`
	ReadGroups []groupDefinition
	client     *modbus.ModbusClient
	quit       chan int
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
	a.Hostname = s
	fmt.Println(a.Hostname)
	var err error
	a.client, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL:     a.Hostname,
		Timeout: 1 * time.Second,
	})
	if err != nil {
		// error out if client creation failed
		fmt.Println("Failed to create")

	}
	a.quit = make(chan int)
	go modbusReads(a)
}

func (a *App) StopReading() {
	a.quit <- 0
}

type returnData struct {
	StartReg uint     `json:"startReg"`
	RegData  []uint16 `json:"RegData"`
}

func modbusReads(a *App) {
	for {
		select {
		case <-a.quit:
			return
		case <-time.After(time.Second):
			err := a.client.Open()
			if err != nil {
				return
			}
			values := make([]returnData, len(a.ReadGroups))
			for i, v := range a.ReadGroups {
				vals, _ := a.client.ReadRegisters(uint16(v.StartReg), uint16(v.NumReg), modbus.HOLDING_REGISTER)
				values[i].StartReg = v.StartReg
				values[i].RegData = vals
			}
			if len(a.ReadGroups) > 0 {
				runtime.EventsEmit(a.ctx, "RegData", values)
			}

			a.client.Close()
		}

	}
}

func (a *App) GetGroups() []groupDefinition {
	return a.ReadGroups
}

func (a *App) UpdateGroups(s string) {
	json.Unmarshal([]byte(s), &a.ReadGroups)
}
