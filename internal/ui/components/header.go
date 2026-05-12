package components

import (
	"time"

	"q-platform/internal/ui/styles"

	"github.com/charmbracelet/lipgloss"
)

type Header struct{ width int }

func (h *Header) SetSize(width int) { h.width = width }

func (h Header) View() string {
	left := lipgloss.NewStyle().Bold(true).Render("TFORGE") + " • " + lipgloss.NewStyle().Foreground(styles.ColorInfo).Bold(true).Render("terraform") + "  • ws: default"
	mid := "dir: /home/quix/work/infrajump-gl/terraform-modules"
	right := "engine: " + lipgloss.NewStyle().Foreground(styles.ColorSuccess).Render("terraform") + "   time: " + time.Now().Format("15:04:05") + "  • " + lipgloss.NewStyle().Foreground(styles.ColorSuccess).Render("0 running")
	row := lipgloss.JoinHorizontal(lipgloss.Top,
		lipgloss.NewStyle().Width(h.width*30/100).Render(left),
		lipgloss.NewStyle().Width(h.width*43/100).Align(lipgloss.Center).Render(mid),
		lipgloss.NewStyle().Width(h.width-(h.width*30/100+h.width*43/100)-4).Align(lipgloss.Right).Render(right),
	)
	return lipgloss.NewStyle().Width(h.width).Padding(0, 2).Bold(true).Foreground(styles.ColorHeaderText).Background(styles.ColorHeaderBg).Render(row)
}
