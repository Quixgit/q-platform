package components

import (
	"time"

	"q-platform/internal/ui/styles"

	"github.com/charmbracelet/lipgloss"
)

type Header struct { width int }

func (h *Header) SetSize(width int) { h.width = width }

func (h Header) View() string {
	content := lipgloss.JoinHorizontal(lipgloss.Top,
		"TFORGE • terraform • ws: default ",
		"dir: /home/quix/work/infrajump-gl/terraform-modules ",
		"engine: terraform ",
		"time: "+time.Now().Format("15:04:05"),
	)
	return lipgloss.NewStyle().Width(h.width).Padding(0, 2).Bold(true).
		Foreground(styles.ColorHeaderText).Background(styles.ColorHeaderBg).Render(content)
}
