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
	guesses   *Guesses
	primary   int
	secondary int
	keymap    KeyMap
	help      help.Model
}

func InitialModel() Model {
	guesses := fresh_guesses()
	return Model{
		guesses:   &guesses,
		primary:   0,
		secondary: 0,
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

func (m *Model) input(s string) {
	if m.primary < 6 {
		if m.secondary < 5 {
			(*m.guesses)[m.primary][m.secondary] = Guess{s, Incorrect}
			m.secondary++
		} else {
			m.primary++
			m.secondary = 0
			m.input(s)
		}
	}
}

func (m *Model) delete() {
	if 0 <= m.primary {
		if 0 < m.secondary {
			(*m.guesses)[m.primary][m.secondary-1] = Guess{" ", Empty}
			m.secondary--
		} else {
			m.secondary = 5
			m.primary--
			m.delete()
		}
	}
}

func (m *Model) change_value() {
	if m.primary >= 0 && m.secondary >= 0 {
		if m.secondary-1 < 0 {
			m.secondary = 5
			m.primary--
		}

		if (*m.guesses)[m.primary][m.secondary-1].v < 2 {
			(*m.guesses)[m.primary][m.secondary-1].v++
		} else {
			(*m.guesses)[m.primary][m.secondary-1].v = 0
		}
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
			if strings.Contains("abcdefghijklmnopqrstuvwxtz", s) {
				m.input(s)
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
		if unique(word) {
			s.WriteString(darkgreytext.Render(word))
		} else {
			s.WriteString(greentext.Render(word))
		}
		s.WriteRune(' ')

		if i == 71 {
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
