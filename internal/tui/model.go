package tui

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"

	"cashflow/internal/tui/screens"
)

type Model struct {
	GlobalKeys GlobalKeyMap
	Active     screens.Screen
	Width      int
	Height     int
}

func NewModel() *Model {
	return &Model{
		GlobalKeys: NewGlobalKeyMap(),
		Active:     screens.NewWelcomeScreen(),
	}
}

func (m *Model) Init() tea.Cmd {
	return m.Active.Init()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyPressMsg); ok && key.Matches(keyMsg, m.GlobalKeys.Quit) {
		return m, tea.Quit
	}

	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.Width = size.Width
		m.Height = size.Height
	}

	next, cmd := m.Active.Update(msg)
	if next == nil {
		next = m.Active
	}

	if next != m.Active {
		cmd = tea.Sequence(cmd, next.Init())
		if m.Width > 0 && m.Height > 0 {
			resized, sizeCmd := next.Update(tea.WindowSizeMsg{
				Width:  m.Width,
				Height: m.Height,
			})
			if resized != nil {
				next = resized
			}
			cmd = tea.Sequence(cmd, sizeCmd)
		}
	}
	m.Active = next

	return m, cmd
}

func (m *Model) View() tea.View {
	view := m.Active.View()
	view.AltScreen = true
	return view
}
