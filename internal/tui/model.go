package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/prxsto/sunnyside/internal/cfg"
)

var _ tea.Model = &InitPromptModel{}

const (
	confKey = `Config File Path`
)

type InitPromptModel struct {
	inputs  map[string]textinput.Model
	cfgPath string
	done    bool
}

func NewInitPrompt(cfgPath string, userHomeDir string) *InitPromptModel {
	configFilePrompt := textinput.New()
	configFilePrompt.Placeholder = userHomeDir + "/sunnyside/config.toml"
	configFilePrompt.Focus()
	return &InitPromptModel{
		cfgPath: cfgPath,
		inputs: map[string]textinput.Model{
			confKey: configFilePrompt,
		},
	}
}

func (i InitPromptModel) Init() tea.Cmd {
	return textinput.Blink
}

func (i InitPromptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q", "Q":
			return i, tea.Quit
		case "enter":
			i.done = true
			return i, tea.Quit
		}
	}
	cmd := i.updateInputs(msg)
	return i, cmd
}

func (i *InitPromptModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, 0)
	for k := range i.inputs {
		if i.inputs[k].Focused() {
			m, cmd := i.inputs[k].Update(msg)
			i.inputs[k] = m
			cmds = append(cmds, cmd)
		}
	}
	return tea.Batch(cmds...)
}

func (i InitPromptModel) View() string {
	// write output file
	if i.done {
		v := i.inputs[confKey]
		if v.Value() == "" {
			v.SetValue(v.Placeholder)
		}
		config := &cfg.Config{
			ConfigFile: v.Value(),
		}
		err := cfg.ToFile(i.cfgPath, config)
		if err != nil {
			return err.Error()
		}
		return "Setup complete! \n"
	}
	output := strings.Builder()
	// write output to screen
	for k, v := range i.inputs {
		output.WriteString(k + "\n")
		output.WriteString(v.View())
	}
	return output.String()
}
