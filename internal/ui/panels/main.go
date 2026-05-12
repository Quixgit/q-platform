package panels

import (
	"strings"

	"q-platform/internal/ui/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainPanel struct { width,height,tab int; focused bool }
func NewMainPanel()*MainPanel{return &MainPanel{}}
func (p *MainPanel) Init() tea.Cmd { return nil }
func (p *MainPanel) SetSize(w,h int){p.width=w;p.height=h}
func (p *MainPanel) SetFocused(f bool){p.focused=f}
func (p *MainPanel) Update(msg tea.Msg)(tea.Model,tea.Cmd){ if k,ok:=msg.(tea.KeyMsg); ok && p.focused { switch k.String(){case "1","2","3","4","5","6": p.tab=int(k.String()[0]-'1')} }; return p,nil }
func (p *MainPanel) View() string {
	tabs:=[]string{"1 Overview","2 Providers","3 Risk","4 Execution","5 Docs","6 Graph"}
	var tr []string
	for i,t:= range tabs{st:=lipgloss.NewStyle().Padding(0,1).Border(lipgloss.RoundedBorder()); if i==p.tab { st=st.Background(styles.ColorAccent).Foreground(lipgloss.Color("#FFFFFF"))}; tr=append(tr,st.Render(t))}
	header:=lipgloss.NewStyle().Bold(true).Foreground(styles.ColorAccent).Render("MODULE WORKSPACE")
	line1:=lipgloss.NewStyle().Bold(true).Render("autoscaling-group ⚠ High risk")
	line2:=lipgloss.NewStyle().Foreground(styles.ColorDim).Render("path: AWS/autoscaling-group  type: module  engine: terraform  updated: 2h ago")
	cards:=lipgloss.JoinHorizontal(lipgloss.Top, box("Summary\nResources: 24\nVariables: 18\nOutputs: 6", (p.width-8)/3), box("Risk Level\nHIGH\n██████░░░░\nScore: 78/100", (p.width-8)/3), box("Last Execution\nPlan 2h ago\n✓ Success\n+12 ~3 -1", (p.width-8)/3))
	table:=box("Recent History\nTime   Action    User   Status    Changes\n2h ago plan      quix   ✓ Success  +12 ~3 -1\n1d ago apply     quix   ✓ Success  +12 ~3 -1", p.width-4)
	content:=strings.Join([]string{header,"",line1,line2,"",lipgloss.JoinHorizontal(lipgloss.Top,tr...),"",cards,"",box("Description\nCreates an Auto Scaling Group with configurable capacity, health checks, and policies.",p.width-4),"",box("Tags\n[aws] [compute] [autoscaling] [high-availability] [network]",p.width-4),"",table},"\n")
	inner:=lipgloss.NewStyle().Width(max(1,p.width-4)).Height(max(1,p.height-4)).Render(content)
	return styles.BasePanelStyle(p.focused).Width(p.width).Height(p.height).Render(inner)
}
func box(c string,w int) string { return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(styles.ColorBorder).Padding(0,1).Width(max(1,w)).Render(c) }
