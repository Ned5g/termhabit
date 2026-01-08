package tui

import "fmt"

func (m Model) View() string {
	styles := MakeStyles(m.Cfg)

	s := styles["title"].Render("TermHabit") + "\n\n"

	for i, h := range m.Habits.Habits {
		cursor := " "
		if m.Cursor == i {
			cursor = styles["cursor"].Render(">")
		}

		check := "\u0078"
		name := h.Name

		if h.Done {
			check = "\u2713"
			name = styles["done"].Render(h.Name)
		} else {
			name = styles["name"].Render(h.Name)
		}

		s += fmt.Sprintf("%s %s %s\n", cursor, check, name)
	}

	s += "\n"
	switch m.Mode {
	case modeNormal:
		s += "j/k: move | space: toggle | a: add | d: delete | q: quit"

	case modeAdd:
		s += styles["input"].Render("New Habit: " + m.Input)
	}

	return styles["border"].Render(s)
}
