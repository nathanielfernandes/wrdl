package wrdl

import "github.com/charmbracelet/lipgloss"

var (
	grey   = lipgloss.Color("#787c7e")
	yellow = lipgloss.Color("#c9b458")
	green  = lipgloss.Color("#6aaa64")
	white  = lipgloss.Color("#d3d6da")

	greytext     = lipgloss.NewStyle().Foreground(lipgloss.Color("#bababa"))
	darkgreytext = lipgloss.NewStyle().Foreground(grey)
	greentext    = lipgloss.NewStyle().Foreground(green)

	GreyTile   = lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.ThickBorder()).Foreground(grey).BorderForeground(grey)
	YellowTile = lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.ThickBorder()).Foreground(yellow).BorderForeground(yellow)
	GreenTile  = lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.ThickBorder()).Foreground(green).BorderForeground(green)
	WhiteTile  = lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.ThickBorder()).Foreground(white).BorderForeground(white)

	title        = lipgloss.JoinHorizontal(lipgloss.Center, GreenTile.Render("W"), YellowTile.Render("R"), GreenTile.Render("D"), GreenTile.Render("L"))
	results      = lipgloss.NewStyle().Width(23).Height(15).MarginLeft(1).Border(lipgloss.HiddenBorder()).BorderForeground(white)
	outer_border = lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.ThickBorder()).BorderForeground(white)
)
