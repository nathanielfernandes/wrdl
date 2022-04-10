package wrdl

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	quit   key.Binding
	change key.Binding
}

type Model struct {
	guesses *Guesses
	row     int
	column  int
	keymap  KeyMap
	help    help.Model
}

func InitialModel() Model {
	guesses := fresh_guesses()
	return Model{
		guesses: &guesses,
		row:     0,
		column:  0,
		keymap: KeyMap{
			quit: key.NewBinding(
				key.WithKeys("ctrl+c"),
				key.WithHelp("ctrl+c", "quit"),
			),
			change: key.NewBinding(
				key.WithKeys(" "),
				key.WithHelp("space", "cycle"),
			),
		},
		help: help.NewModel(),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m Model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.change,
		m.keymap.quit,
	})
}

func (m *Model) input(s string, v Value) {
	if m.column < 5 {
		(*m.guesses)[m.row][m.column] = Guess{s, v}
	} else {
		m.column = 4
	}

	m.column++
	if m.column == 5 {
		if m.row < 5 {
			m.column = 0
			m.row++
		}
	}
}

func (m *Model) delete() Guess {
	m.column--
	if m.column == -1 {
		if m.row > 0 {
			m.column = 4
			m.row--
		}
	}

	old := Guess{" ", Empty}
	if m.column >= 0 {
		old = (*m.guesses)[m.row][m.column]
		(*m.guesses)[m.row][m.column] = Guess{" ", Empty}
	} else {
		m.column = 0
	}

	return old
}

func (m *Model) change_value() {
	old := m.delete()
	if old.v != 3 {
		if old.v == 2 {
			old.v = 0
		} else {
			old.v++
		}
		m.input(old.c, old.v)
	}

}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		s := msg.String()

		switch s {
		case "ctrl+c":
			return m, tea.Quit

		case " ":
			m.change_value()

		case "backspace":
			m.delete()

		default:
			if strings.Contains("abcdefghijklmnopqrstuvwxyz", s) {
				m.input(s, Incorrect)
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	var (
		res = Results(*m.guesses)
		n   = strconv.Itoa(len(res))
		s   = strings.Builder{}
		i   = 0
	)

	for _, word := range res {
		if !unique(word) {
			s.WriteString(darkgreytext.Render(word))
		} else {
			s.WriteString(greentext.Render(word))
		}
		s.WriteRune(' ')

		if i == 67 {
			break
		}

		if i == 3 {
			s.WriteRune('\n')
		}
		i++
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		outer_border.Render(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				lipgloss.JoinVertical(
					lipgloss.Center,
					title,
					m.guesses.Tiles(),
				),
				results.Render("Possible Words: "+greytext.Render(n)+"\n\n"+s.String()),
			),
		),
		m.helpView(),
	)
}
