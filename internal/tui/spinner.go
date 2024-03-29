package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Loading struct {
	spinner    spinner.Model
	loading    bool
	loaderText string
	response   chan struct{}
}

func (l Loading) Init() tea.Cmd {
	return l.spinner.Tick
}

func (l Loading) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	l.spinner, _ = l.spinner.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return l, tea.Quit
		}
	}
	select {
	case <-l.response:
		l.loading = false
		return l, tea.Quit
	default:
		return l, l.spinner.Tick
	}
}

func (l Loading) View() string {
	if l.loading {
		return fmt.Sprintf("%s %s ...", l.spinner.View(), l.loaderText)
	} else {
		return ""
	}
}

func RenderLoad(loaderText string, response chan struct{}) {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	p := tea.NewProgram(Loading{spinner: s, loading: true, response: response, loaderText: loaderText})
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
