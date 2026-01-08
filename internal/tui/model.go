package tui

import (
	"github.com/Ned5g/termhabit/storage"
)

type mode int

const (
	modeNormal mode = iota
	modeAdd
	modeDelete
)

type Model struct {
	Habits storage.HabitData
	Cursor int
	Input  string
	Mode   mode
	Cfg    *storage.Config
}
