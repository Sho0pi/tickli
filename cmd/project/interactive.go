package project

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/sho0pi/tickli/internal/types/project"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sho0pi/tickli/internal/types"
)

// Define styles for the UI
var (
	titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF79C6")).Bold(true)
	inputStyle = lipgloss.NewStyle().PaddingLeft(1)
	helpStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4"))
)

// model represents the UI state for the interactive mode.
type model struct {
	inputs  []textinput.Model
	focused int
	opts    *createProjectOptions
	done    bool
}

// initialModel initializes the model with the provided options.
func initialModel(opts *createProjectOptions) model {
	var inputs []textinput.Model

	// Initialize text inputs with placeholders from opts
	nameInput := textinput.New()
	nameInput.Placeholder = opts.name
	nameInput.Focus()

	colorInput := textinput.New()
	colorInput.Placeholder = opts.color.String()

	viewModeInput := textinput.New()
	viewModeInput.Placeholder = string(opts.viewMode)

	kindInput := textinput.New()
	kindInput.Placeholder = string(opts.kind)

	inputs = append(inputs, nameInput, colorInput, viewModeInput, kindInput)

	return model{
		inputs:  inputs,
		focused: 0,
		opts:    opts,
		done:    false,
	}
}

// Init is required by the bubbletea.Model interface.
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles input and updates the model.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if m.focused == len(m.inputs)-1 {
				m.done = true
				return m, tea.Quit
			}
			m.nextInput()
		case "tab", "shift+tab":
			if msg.String() == "tab" {
				m.nextInput()
			} else {
				m.prevInput()
			}
		}
	}

	// Update the focused input
	cmd := m.updateInputs(msg)
	return m, cmd
}

// View renders the UI.
func (m model) View() string {
	if m.done {
		return ""
	}

	var b strings.Builder

	b.WriteString(titleStyle.Render("Create a new project\n\n"))

	inputs := []string{
		inputStyle.Render(fmt.Sprintf("Name: %s", m.inputs[0].View())),
		inputStyle.Render(fmt.Sprintf("Color: %s", m.inputs[1].View())),
		inputStyle.Render(fmt.Sprintf("View Mode: %s", m.inputs[2].View())),
		inputStyle.Render(fmt.Sprintf("Kind: %s", m.inputs[3].View())),
	}

	b.WriteString(strings.Join(inputs, "\n"))
	b.WriteString("\n\n" + helpStyle.Render("(tab/shift+tab to navigate, enter to submit)"))

	return b.String()
}

// nextInput moves focus to the next input field.
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
	for i := range m.inputs {
		m.inputs[i].Blur()
	}
	m.inputs[m.focused].Focus()
}

// prevInput moves focus to the previous input field.
func (m *model) prevInput() {
	m.focused--
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
	for i := range m.inputs {
		m.inputs[i].Blur()
	}
	m.inputs[m.focused].Focus()
}

// updateInputs updates the focused input field.
func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return tea.Batch(cmds...)
}

// runInteractiveMode starts the interactive UI.
func runInteractiveMode(opts *createProjectOptions) (*types.Project, error) {
	p := tea.NewProgram(initialModel(opts))
	m, err := p.Run()
	if err != nil {
		return nil, err
	}

	model := m.(model)
	if !model.done {
		return nil, fmt.Errorf("interactive mode canceled")
	}

	// Create a project from the inputs
	project := &types.Project{
		Name:     model.inputs[0].Value(),
		Color:    project.Color(color.HEX(model.inputs[1].Value())),
		ViewMode: project.ViewMode(model.inputs[2].Value()),
		Kind:     project.Kind(model.inputs[3].Value()),
	}

	return project, nil
}
