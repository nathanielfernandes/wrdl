package wrdl

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func Wrld(cmd *cobra.Command, args []string) {
	m := InitialModel()
	if err := tea.NewProgram(&m).Start(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}
