package panels

import (
	"strings"

	"q-platform/internal/ui/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LogsPanel struct { width,height,offset int; focused bool; lines []string }
func NewLogsPanel()*LogsPanel{return &LogsPanel{lines:[]string{"14:29:11 [INFO] terraform init --backend=true","14:29:14 [INFO] Initializing modules...","14:29:15 [INFO] Initializing provider plugins...","14:29:17 [WARN] Using previously-installed hashicorp/aws","14:29:18 [INFO] Terraform has been initialized"}}}
func (p *LogsPanel) Init() tea.Cmd { return nil }
func (p *LogsPanel) SetSize(w,h int){p.width=w;p.height=h}
func (p *LogsPanel) SetFocused(f bool){p.focused=f}
func (p *LogsPanel) Update(msg tea.Msg)(tea.Model,tea.Cmd){
	viewport:=max(1,p.height-8)
	switch m:=msg.(type){
	case tea.KeyMsg:
		if p.focused { switch m.String(){case "up","k": if p.offset>0{p.offset--}; case "down","j": maxOff:=max(0,len(p.lines)-viewport); if p.offset<maxOff{p.offset++}; case "L": p.lines=[]string{}; p.offset=0}}
	case tea.MouseMsg:
		if p.focused { if m.Button==tea.MouseButtonWheelUp && p.offset>0 {p.offset--}; if m.Button==tea.MouseButtonWheelDown { maxOff:=max(0,len(p.lines)-viewport); if p.offset<maxOff{p.offset++}} }
	}
	return p,nil
}
func (p *LogsPanel) View() string {
	viewport:=max(1,p.height-8); start:=p.offset; if start>len(p.lines){start=len(p.lines)}; end:=start+viewport; if end>len(p.lines){end=len(p.lines)}
	visible:=[]string{}
	if len(p.lines)>0 { visible=p.lines[start:end] }
	header:=lipgloss.NewStyle().Bold(true).Foreground(styles.ColorAccent).Render("EXECUTION LOGS   Command: plan   00:00:18")
	body:=strings.Join(visible,"\n")
	footer:="[Clear Logs L] [i INFO: 12] [⚠ WARN: 1] [✗ ERROR: 0]"
	content:=lipgloss.JoinVertical(lipgloss.Left, header, body, "", lipgloss.NewStyle().Foreground(styles.ColorDim).Render(footer))
	inner:=lipgloss.NewStyle().Width(max(1,p.width-4)).Height(max(1,p.height-4)).Render(content)
	return styles.BasePanelStyle(p.focused).Width(p.width).Height(p.height).Render(inner)
}
