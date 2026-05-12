package styles

import "github.com/charmbracelet/lipgloss"

var (
	ColorBg          = lipgloss.Color("#090D2E")
	ColorPanelBg     = lipgloss.Color("#0C1036")
	ColorSurface     = lipgloss.Color("#171B49")
	ColorSurfaceAlt  = lipgloss.Color("#111542")
	ColorBorder      = lipgloss.Color("#2E347D")
	ColorBorderFocus = lipgloss.Color("#7C4DFF")
	ColorText        = lipgloss.Color("#D7DBF7")
	ColorDim         = lipgloss.Color("#8B90B3")
	ColorHeaderBg    = lipgloss.Color("#090D2E")
	ColorHeaderText  = lipgloss.Color("#D7DBF7")
	ColorAccent      = lipgloss.Color("#8B5CF6")
	ColorSuccess     = lipgloss.Color("#4ADE80")
	ColorWarn        = lipgloss.Color("#FBBF24")
	ColorError       = lipgloss.Color("#F87171")
	ColorInfo        = lipgloss.Color("#22D3EE")
	ColorStatusBg    = lipgloss.Color("#0A1033")
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

func CardStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(ColorSurface).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorBorder).
		Padding(0, 1)
}
