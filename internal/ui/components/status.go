package components

import (
	"q-platform/internal/ui/styles"

	"github.com/charmbracelet/lipgloss"
)

type StatusBar struct { width int }

func (s *StatusBar) SetSize(width int) { s.width = width }

func (s StatusBar) View() string {
	text := "✓ Ready | /:search | Tab:focus | Enter:select | q:quit"
	return lipgloss.NewStyle().Width(s.width).Padding(0, 2).
		Foreground(styles.ColorSuccess).Background(styles.ColorStatusBg).Render(text)
}
