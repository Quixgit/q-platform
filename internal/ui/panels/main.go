package panels

import (
	"strings"

	"q-platform/internal/ui/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainPanel struct {
	width, height, tab int
	focused            bool
}

func NewMainPanel() *MainPanel         { return &MainPanel{} }
func (p *MainPanel) Init() tea.Cmd     { return nil }
func (p *MainPanel) SetSize(w, h int)  { p.width, p.height = w, h }
func (p *MainPanel) SetFocused(f bool) { p.focused = f }
func (p *MainPanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if k, ok := msg.(tea.KeyMsg); ok && p.focused {
		switch k.String() {
		case "1", "2", "3", "4", "5", "6":
			p.tab = int(k.String()[0] - '1')
		}
	}
	return p, nil
}

func (p *MainPanel) View() string {
	tabs := []string{"1 Overview", "2 Providers", "3 Risk", "4 Execution", "5 Docs", "6 Graph"}
	tabRow := make([]string, 0, len(tabs))
	for i, t := range tabs {
		st := lipgloss.NewStyle().Padding(0, 2).Border(lipgloss.RoundedBorder()).BorderForeground(styles.ColorBorder)
		if i == p.tab {
			st = st.Background(styles.ColorAccent).Foreground(lipgloss.Color("#FFFFFF")).Bold(true)
		}
		tabRow = append(tabRow, st.Render(t))
	}
	w := p.width - 4
	col := (w - 4) / 3
	summary := styles.CardStyle().Width(col).Render("Summary\n\nResources:      24\nVariables:      18\nOutputs:        6\nData Sources:   3\nModules:        2")
	risk := styles.CardStyle().Width(col).Render("Risk Level\n\n   " + lipgloss.NewStyle().Bold(true).Foreground(styles.ColorError).Render("HIGH") + "\n\n   " + lipgloss.NewStyle().Foreground(styles.ColorError).Render("██") + lipgloss.NewStyle().Foreground(lipgloss.Color("#FB923C")).Render("██") + lipgloss.NewStyle().Foreground(styles.ColorWarn).Render("██") + lipgloss.NewStyle().Foreground(styles.ColorDim).Render("██") + "\n\nScore: 78/100")
	exec := styles.CardStyle().Width(col).Render("Last Execution\n\n" + lipgloss.NewStyle().Foreground(styles.ColorSuccess).Render("Plan") + "\n2h ago\n" + lipgloss.NewStyle().Foreground(styles.ColorSuccess).Render("✓ Success") + "\nDuration: 18s\nChanges: " + lipgloss.NewStyle().Foreground(styles.ColorSuccess).Render("+12") + " " + lipgloss.NewStyle().Foreground(styles.ColorWarn).Render("~3") + " " + lipgloss.NewStyle().Foreground(styles.ColorError).Render("-1"))
	cards := lipgloss.JoinHorizontal(lipgloss.Top, summary, "  ", risk, "  ", exec)
	history := styles.CardStyle().Width(w).Height(11).Render("Recent History\n\nTime      Action    User    Status      Changes\n2h ago    plan      quix    ✓ Success   +12 ~3 -1\n1d ago    apply     quix    ✓ Success   +12 ~3 -1\n2d ago    plan      quix    ✓ Success   +8 ~2 -0\n3d ago    validate  quix    ✓ Success   -\n5d ago    plan      quix    ✓ Success   +5 ~1 -0")
	content := lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.NewStyle().Bold(true).Foreground(styles.ColorAccent).Render("MODULE WORKSPACE"),
		"",
		lipgloss.NewStyle().Bold(true).Render("◈ autoscaling-group ")+lipgloss.NewStyle().Foreground(styles.ColorError).Render("❖ High risk"),
		lipgloss.NewStyle().Foreground(styles.ColorDim).Render("path: AWS/autoscaling-group      type: module      engine: terraform      updated: 2h ago"),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top, tabRow...),
		"",
		cards,
		"",
		styles.CardStyle().Width(w).Render("Description\n\nCreates an Auto Scaling Group with configurable capacity,\nhealth checks, and load balancer integration.\nSupports mixed instances, spot instances, and scaling policies."),
		"",
		styles.CardStyle().Width(w).Render("Tags\n\n[aws]  [compute]  [autoscaling]  [high-availability]  [network]"),
		"",
		history,
	)
	inner := lipgloss.NewStyle().Width(w).Height(p.height - 4).Render(content)
	return styles.BasePanelStyle(p.focused).Width(p.width).Height(p.height).Render(strings.TrimRight(inner, "\n"))
}
