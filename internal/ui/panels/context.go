package panels

import (
	"strings"

	"q-platform/internal/ui/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ContextPanel struct {
	width, height int
	focused       bool
}

func NewContextPanel() *ContextPanel                            { return &ContextPanel{} }
func (p *ContextPanel) Init() tea.Cmd                           { return nil }
func (p *ContextPanel) SetSize(w, h int)                        { p.width, p.height = w, h }
func (p *ContextPanel) SetFocused(f bool)                       { p.focused = f }
func (p *ContextPanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return p, nil }

func (p *ContextPanel) View() string {
	w := p.width - 4
	content := lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.NewStyle().Bold(true).Foreground(styles.ColorAccent).Render("CONTEXT & ACTIONS"),
		"",
		styles.CardStyle().Width(w).Render("Module Info\n\nName:          autoscaling-group\nPath:          AWS/autoscaling-group\nEngine:        terraform\nType:          module\nUpdated:       2h ago\nWorkspace:     default\nRoot:          /home/quix/work/infrajump-gl/terraform-modules"),
		"",
		styles.CardStyle().Width(w).Render("Quick Actions\n\n[p Plan]  [a Apply]  [d Destroy]  [v Validate]  [s Security]\n[i Init]  [g Graph]  [o Open in $EDITOR]  [x Show in Explorer]"),
		"",
		styles.CardStyle().Width(w).Render("Providers\n\naws       -> 5.0              Required: 3     Configured: 3\nrandom    -> 3.6\nnull      -> 3.0"),
		"",
		styles.CardStyle().Width(w).Render("Risk Breakdown\n\nCritical: 3  "+lipgloss.NewStyle().Foreground(styles.ColorError).Render("██████")+"\nHigh:     6  "+lipgloss.NewStyle().Foreground(lipgloss.Color("#FB923C")).Render("██████████")+"\nMedium:   8  "+lipgloss.NewStyle().Foreground(styles.ColorWarn).Render("████████")+"\nLow:      7  "+lipgloss.NewStyle().Foreground(styles.ColorSuccess).Render("███████")),
		"",
		styles.CardStyle().Width(w).Render("Dependencies\n\nModules: 2\nResources: 24\nData Sources: 3\n\nRelated Workspaces\ndev      last run: 2h ago  ✓\nstaging  last run: 1d ago  ✓\nprod     last run: 2d ago  ⚠"),
	)
	inner := lipgloss.NewStyle().Width(w).Height(p.height - 4).Render(content)
	return styles.BasePanelStyle(p.focused).Width(p.width).Height(p.height).Render(strings.TrimRight(inner, "\n"))
}
