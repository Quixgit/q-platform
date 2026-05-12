package panels

import (
	"strings"

	"q-platform/internal/ui/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NavigationPanel struct { width,height,cursor int; focused bool; items []string }

func NewNavigationPanel() *NavigationPanel { return &NavigationPanel{items: []string{"Terraform","  AWS","    autoscaling-group 🔴","    aws-backup 🟢","    cloudfront 🟡","OpenTofu","Helm","Projects"}} }
func (p *NavigationPanel) Init() tea.Cmd { return nil }
func (p *NavigationPanel) SetSize(w,h int){p.width=w;p.height=h}
func (p *NavigationPanel) SetFocused(f bool){p.focused=f}
func (p *NavigationPanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if k,ok:=msg.(tea.KeyMsg); ok && p.focused {
		switch k.String(){case "up","k": if p.cursor>0 {p.cursor--}; case "down","j": if p.cursor<len(p.items)-1 {p.cursor++}}
	}
	return p,nil
}
func (p *NavigationPanel) View() string {
	var rows []string
	title := lipgloss.NewStyle().Bold(true).Background(styles.ColorAccent).Padding(0,1).Render("NAVIGATION")
	rows=append(rows,title,"")
	for i,it := range p.items { st:=lipgloss.NewStyle(); if i==p.cursor { st=st.Background(styles.ColorAccent) }; rows=append(rows,st.Render(it)) }
	rows=append(rows,"",lipgloss.NewStyle().Foreground(styles.ColorDim).Render("102 modules"))
	inner:=lipgloss.NewStyle().Width(max(1,p.width-4)).Height(max(1,p.height-4)).Render(strings.Join(rows,"\n"))
	return styles.BasePanelStyle(p.focused).Width(p.width).Height(p.height).Render(inner)
}
func max(a,b int) int { if a>b {return a}; return b }
