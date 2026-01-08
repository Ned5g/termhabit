package tui

import (
	"time"

	"github.com/Ned5g/termhabit/storage"
	tea "github.com/charmbracelet/bubbletea"
)

func InitialModel(appName string, cfg *storage.Config) Model {
	return Model{
		Habits: storage.ReadHabitFile(appName),
		Cfg:    cfg,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) AddHabit(msg tea.KeyMsg) Model {
	switch msg.String() {
	case "enter":
		m.Habits.Habits = append(m.Habits.Habits, storage.Habit{
			Name: m.Input,
		})
		m.save()
		m.Mode = modeNormal

	case "esc":
		m.Mode = modeNormal

	case "backspace":
		if len(m.Input) > 0 {
			m.Input = m.Input[:len(m.Input)-1]
		}

	default:
		m.Input += msg.String()
	}

	return m
}

func (m Model) NormalMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q":
		return m, tea.Quit
	case "j":
		if m.Cursor < len(m.Habits.Habits)-1 {
			m.Cursor++
		}
	case "k":
		if m.Cursor > 0 {
			m.Cursor--
		}
	case "a":
		m.Mode = modeAdd
		m.Input = ""
	case "d":
		if len(m.Habits.Habits) > 0 {
			m.Habits.Habits = append(
				m.Habits.Habits[:m.Cursor],
				m.Habits.Habits[m.Cursor+1:]...,
			)

			if m.Cursor >= len(m.Habits.Habits) {
				m.Cursor = len(m.Habits.Habits) - 1
			}

			if m.Cursor < 0 {
				m.Cursor = 0
			}

			m.save()
		}
	case " ":
		if m.Cursor >= 0 && m.Cursor < len(m.Habits.Habits) {
			m.Habits.Habits[m.Cursor].Done = !m.Habits.Habits[m.Cursor].Done
			m.save()
		}
	}

	return m, nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.Mode {

		case modeNormal:
			return m.NormalMode(msg)

		case modeAdd:
			m = m.AddHabit(msg)
		}
	}

	return m, nil
}

func (m *Model) save() {
	m.Habits.LastUpdated = time.Now().Format(time.RFC3339)
	storage.WriteHabitFile("TermHabit", m.Habits)
}
