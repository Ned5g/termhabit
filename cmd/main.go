package main

import (
	"log"

	tui "github.com/Ned5g/termhabit/internal/tui"
	"github.com/Ned5g/termhabit/storage"
	tea "github.com/charmbracelet/bubbletea"
)

const appName = "termhabit"

func main() {
	cfg, err := storage.LoadConfig(appName)
	if err != nil {
		log.Printf("failed to load config file: %v", err)
	}

	p := tea.NewProgram(tui.InitialModel(appName, cfg))
	_, err = p.Run()
	if err != nil {
		log.Printf("failed to run termhabit: ", err)
	}
}
