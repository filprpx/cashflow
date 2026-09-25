package screens

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

type NewVaultKeyMap struct {
	Back key.Binding
}

type VaultCreationScreen struct {
	keys NewVaultKeyMap
}

func NewVaultCreationScreen() *VaultCreationScreen {
	return &VaultCreationScreen{
		keys: NewVaultKeyMap{
			Back: key.NewBinding(
				key.WithKeys("esc"),
				key.WithHelp("esc", "back"),
			),
		},
	}
}

func (m *VaultCreationScreen) Init() tea.Cmd {
	return nil
}

func (m *VaultCreationScreen) Update(msg tea.Msg) (Screen, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyPressMsg)
	if ok && key.Matches(keyMsg, m.keys.Back) {
		return NewWelcomeScreen(), nil
	}

	return m, nil
}

func (m *VaultCreationScreen) View() tea.View {
	return tea.NewView("Create a new vault\n\nVault creation will be implemented here.\n\n[esc] Back\n")
}
