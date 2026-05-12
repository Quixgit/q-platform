package panels

import (
	"strings"

	"q-platform/internal/ui/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LogsPanel struct {
	width, height, offset int
	focused               bool
	lines                 []string
}

func NewLogsPanel() *LogsPanel {
	return &LogsPanel{lines: []string{"14:29:11 [INFO] terraform init --backend=true", "14:29:14 [INFO] Initializing modules...", "14:29:15 [INFO] Initializing provider plugins...", "14:29:16 [INFO] Finding hashicorp/aws versions...", "14:29:17 [WARN] Using previously-installed hashicorp/aws v5.36.0", "14:29:18 [INFO] Terraform has been successfully initialized!"}}
}
func (p *LogsPanel) Init() tea.Cmd     { return nil }
func (p *LogsPanel) SetSize(w, h int)  { p.width, p.height = w, h }
func (p *LogsPanel) SetFocused(f bool) { p.focused = f }
func (p *LogsPanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	viewport := max(1, p.height-8)
	switch m := msg.(type) {
	case tea.KeyMsg:
		if p.focused {
			switch m.String() {
			case "up", "k":
				if p.offset > 0 {
					p.offset--
				}
			case "down", "j":
				maxOff := max(0, len(p.lines)-viewport)
				if p.offset < maxOff {
					p.offset++
				}
			case "L", "l":
				p.lines = []string{}
				p.offset = 0
			}
		}
	case tea.MouseMsg:
		if p.focused && m.Button == tea.MouseButtonWheelUp && p.offset > 0 {
			p.offset--
		}
		if p.focused && m.Button == tea.MouseButtonWheelDown {
			maxOff := max(0, len(p.lines)-viewport)
			if p.offset < maxOff {
				p.offset++
			}
		}
	}
	return p, nil
}
func (p *LogsPanel) View() string {
	viewport := max(1, p.height-8)
	start, end := p.offset, p.offset+viewport
	if start > len(p.lines) {
		start = len(p.lines)
	}
	if end > len(p.lines) {
		end = len(p.lines)
	}
	vis := []string{}
	if len(p.lines) > 0 {
		vis = p.lines[start:end]
	}
	body := strings.Join(vis, "\n")
	content := lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.JoinHorizontal(lipgloss.Top, lipgloss.NewStyle().Bold(true).Foreground(styles.ColorAccent).Render("EXECUTION LOGS"), "   ", "Command:", lipgloss.NewStyle().Foreground(styles.ColorSuccess).Render(" plan"), "   00:00:18"),
		body,
		"",
		lipgloss.NewStyle().Foreground(styles.ColorDim).Render("[Clear Logs L]  [i INFO: 12]  [⚠ WARN: 1]  [✗ ERROR: 0]"),
	)
	inner := lipgloss.NewStyle().Width(max(1, p.width-4)).Height(max(1, p.height-4)).Render(content)
	return styles.BasePanelStyle(p.focused).Width(p.width).Height(p.height).Render(inner)
}
