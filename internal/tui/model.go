package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

var _ tea.Model = &InitPromptModel{}

const (
	histKey = `History File Path`
)

type InitPromptModel struct{}

func (i InitPromptModel) Init() tea.Cmd {
	return nil
}

func (i InitPromptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (i InitPromptModel) View() string {
	return ""
}
