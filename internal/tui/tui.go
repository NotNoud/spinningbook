package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/NotNoud/spinningbook/internal/macro"
)

type Model struct {
	macros []macro.Macro
	cursor int
	width  int
	height int
}

func NewModel() Model {
	return Model{
		macros: []macro.Macro{
			{
				Name:        "moovment",
				Description: "Open Moovment workspace en start Claude Code",
				Steps: []macro.Step{
					{Type: macro.StepCD, Value: `C:\Users\Noud\Documents\Zzp\Moovment`, Platform: macro.PlatformWindows},
					{Type: macro.StepRun, Value: "claude", Platform: macro.PlatformAll},
				},
			},
		},
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j", "down":
			if m.cursor < len(m.macros)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "g":
			m.cursor = 0
		case "G":
			m.cursor = len(m.macros) - 1
		}
	}
	return m, nil
}

var (
	pane   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1)
	active = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	dim    = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	title  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("99"))
)

func (m Model) View() string {
	if m.width == 0 {
		return "starting..."
	}

	leftW := m.width / 3
	rightW := m.width - leftW - 4
	bodyH := m.height - 4

	var list string
	for i, mc := range m.macros {
		line := mc.Name
		if i == m.cursor {
			list += active.Render("▸ "+line) + "\n"
		} else {
			list += "  " + line + "\n"
		}
	}
	left := pane.Width(leftW).Height(bodyH).Render(title.Render("Macro's") + "\n\n" + list)

	var preview string
	if len(m.macros) > 0 {
		current := m.macros[m.cursor]
		preview = title.Render(current.Name) + "\n" + dim.Render(current.Description) + "\n\n"
		for i, step := range current.Steps {
			preview += dim.Render(fmt.Sprintf("%d. ", i+1)) + string(step.Type) + " " + step.Value + "\n"
		}
	}
	right := pane.Width(rightW).Height(bodyH).Render(title.Render("Preview") + "\n\n" + preview)

	help := dim.Render("j/k navigeren  •  g/G top/bottom  •  n nieuw  •  e edit  •  d delete  •  q afsluiten")

	return lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.JoinHorizontal(lipgloss.Top, left, right),
		help,
	)
}
