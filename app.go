package main

import (
	cmdref "changeme/internal"
	"context"
	"fmt"
	"io/ioutil"
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

// GetCommands returns a commands from the .json file
func (a *App) GetCommands() (string, error) {
	fileOperator := cmdref.NewCmdFileOps()
	fileBytes, err := ioutil.ReadFile(fileOperator.GetFilePath())
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}

	fileContent := string(fileBytes)
	return fileContent, nil
}
