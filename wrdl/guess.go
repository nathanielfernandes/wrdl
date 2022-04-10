package wrdl

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Guess struct {
	c string
	v Value
}

func (g Guess) tile() string {
	c := strings.ToUpper(g.c)
	switch g.v {
	case Incorrect:
		return WhiteTile.Render(c)

	case Valid:
		return YellowTile.Render(c)

	case Determined:
		return GreenTile.Render(c)

	case Empty:
		return GreyTile.Render(c)
	}
	return ""
}

type FullGuess []Guess
type Guesses []FullGuess

func fresh_guesses() Guesses {
	gs := Guesses{}
	for i := 0; i < 6; i++ {
		gs = append(gs, FullGuess{{" ", Empty}, {" ", Empty}, {" ", Empty}, {" ", Empty}, {" ", Empty}})
	}

	return gs
}

func (gs FullGuess) Tiles() string {
	return lipgloss.JoinHorizontal(lipgloss.Center, gs[0].tile(), gs[1].tile(), gs[2].tile(), gs[3].tile(), gs[4].tile())
}

func (gs Guesses) Tiles() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		gs[0].Tiles(),
		gs[1].Tiles(),
		gs[2].Tiles(),
		gs[3].Tiles(),
		gs[4].Tiles(),
		gs[5].Tiles(),
	)
}
