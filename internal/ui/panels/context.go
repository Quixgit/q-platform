package panels

import (
	"strings"

	"q-platform/internal/ui/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ContextPanel struct { width,height int; focused bool }
func NewContextPanel()*ContextPanel{return &ContextPanel{}}
func (p *ContextPanel) Init() tea.Cmd { return nil }
func (p *ContextPanel) SetSize(w,h int){p.width=w;p.height=h}
func (p *ContextPanel) SetFocused(f bool){p.focused=f}
func (p *ContextPanel) Update(msg tea.Msg)(tea.Model,tea.Cmd){ return p,nil }
func (p *ContextPanel) View() string {
	header:=lipgloss.NewStyle().Bold(true).Foreground(styles.ColorAccent).Render("CONTEXT & ACTIONS")
	blocks:=[]string{box("Module Info\nName: autoscaling-group\nPath: AWS/autoscaling-group\nEngine: terraform\nType: module",p.width-4),box("Quick Actions\n[p Plan] [a Apply] [d Destroy]\n[v Validate] [s Security] [i Init] [g Graph]",p.width-4),box("Providers\naws -> 5.0 Required: 3 Configured: 3",p.width-4),box("Risk Breakdown\nCritical: ███\nHigh: ██████\nMedium: █████\nLow: ████",p.width-4)}
	inner:=lipgloss.NewStyle().Width(max(1,p.width-4)).Height(max(1,p.height-4)).Render(strings.Join(append([]string{header,""},blocks...),"\n\n"))
	return styles.BasePanelStyle(p.focused).Width(p.width).Height(p.height).Render(inner)
}
