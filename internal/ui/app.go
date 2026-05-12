package ui

import (
	"q-platform/internal/ui/components"
	"q-platform/internal/ui/panels"
	"q-platform/internal/ui/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type App struct {
	width,height int
	focused PanelFocus
	header components.Header
	status components.StatusBar
	nav *panels.NavigationPanel
	main *panels.MainPanel
	ctx *panels.ContextPanel
	logs *panels.LogsPanel
}

func NewApp() *App { return &App{focused:FocusNav,nav:panels.NewNavigationPanel(),main:panels.NewMainPanel(),ctx:panels.NewContextPanel(),logs:panels.NewLogsPanel()} }
func (a *App) Init() tea.Cmd { return nil }
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m:=msg.(type){
	case tea.WindowSizeMsg:
		a.width,a.height=m.Width,m.Height
		a.layout()
	case tea.KeyMsg:
		switch m.String(){case "q","ctrl+c": return a, tea.Quit; case "tab": a.focused=(a.focused+1)%4; a.setFocus(); case "shift+tab": a.focused=(a.focused+3)%4; a.setFocus() }
	}
	for _,p := range []PanelComponent{a.nav,a.main,a.ctx,a.logs} { _,_ = p.Update(msg) }
	return a,nil
}
func (a *App) View() string {
	if a.width < 100 || a.height < 30 {
		return lipgloss.NewStyle().Foreground(styles.ColorWarn).Render("Terminal too small")
	}
	bodyH:=a.height-2
	topH:=bodyH*72/100
	bottomH:=bodyH-topH
	leftW:=a.width*19/100
	centerW:=a.width*48/100
	rightW:=a.width-leftW-centerW
	a.nav.SetSize(leftW,topH); a.main.SetSize(centerW,topH); a.ctx.SetSize(rightW,topH); a.logs.SetSize(a.width,bottomH)
	top:=lipgloss.JoinHorizontal(lipgloss.Top,a.nav.View(),a.main.View(),a.ctx.View())
	content:=lipgloss.JoinVertical(lipgloss.Left,a.header.View(),top,a.logs.View(),a.status.View())
	return lipgloss.NewStyle().Background(styles.ColorBg).Render(content)
}
func (a *App) layout(){ a.header.SetSize(a.width); a.status.SetSize(a.width); a.setFocus() }
func (a *App) setFocus(){ a.nav.SetFocused(a.focused==FocusNav); a.main.SetFocused(a.focused==FocusMain); a.ctx.SetFocused(a.focused==FocusContext); a.logs.SetFocused(a.focused==FocusLogs) }
