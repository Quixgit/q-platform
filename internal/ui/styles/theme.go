package styles

import "github.com/charmbracelet/lipgloss"

var (
	ColorBg         = lipgloss.Color("#080D2B")
	ColorPanelBg    = lipgloss.Color("#0B1240")
	ColorBorder     = lipgloss.Color("#2A2F7A")
	ColorBorderFocus= lipgloss.Color("#7C4DFF")
	ColorText       = lipgloss.Color("#D7DBF7")
	ColorDim        = lipgloss.Color("#8B90B3")
	ColorHeaderBg   = lipgloss.Color("#00FFFF")
	ColorHeaderText = lipgloss.Color("#000000")
	ColorAccent     = lipgloss.Color("#8B5CF6")
	ColorSuccess    = lipgloss.Color("#4ADE80")
	ColorWarn       = lipgloss.Color("#FBBF24")
	ColorError      = lipgloss.Color("#F87171")
	ColorInfo       = lipgloss.Color("#22D3EE")
	ColorStatusBg   = lipgloss.Color("#0A1033")
)

func BasePanelStyle(focused bool) lipgloss.Style {
	border := ColorBorder
	if focused {
		border = ColorBorderFocus
	}
	return lipgloss.NewStyle().
		Background(ColorPanelBg).
		Foreground(ColorText).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(border).
		Padding(1)
}
