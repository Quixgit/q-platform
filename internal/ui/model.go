package ui

import tea "github.com/charmbracelet/bubbletea"

type PanelFocus int

const (
	FocusNav PanelFocus = iota
	FocusMain
	FocusContext
	FocusLogs
)

type PanelComponent interface {
	tea.Model
	SetSize(width, height int)
	SetFocused(bool)
}
