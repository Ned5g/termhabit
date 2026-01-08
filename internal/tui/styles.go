package tui

import (
	"github.com/Ned5g/termhabit/storage"
	"github.com/charmbracelet/lipgloss"
)

func MakeStyles(cfg *storage.Config) map[string]lipgloss.Style {
	return map[string]lipgloss.Style{
		"title": lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(cfg.Styles.Title)),

		"cursor": lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(cfg.Styles.Cursor)),

		"habit": lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(cfg.Styles.Habit)),

		"done": lipgloss.NewStyle().
			Foreground(lipgloss.Color(cfg.Styles.Done)).
			Strikethrough(true),

		"help": lipgloss.NewStyle().
			Foreground(lipgloss.Color(cfg.Styles.Help)),

		"input": lipgloss.NewStyle().
			Foreground(lipgloss.Color(cfg.Styles.Input)),
	}
}
