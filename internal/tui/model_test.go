package tui

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"cashflow/internal/tui/screens"
)

func TestNewModelStartsOnWelcomeScreen(t *testing.T) {
	model := NewModel()

	if _, ok := model.Active.(*screens.Welcome); !ok {
		t.Fatalf("expected welcome screen, got %T", model.Active)
	}
}

func TestWelcomeTransitionsToNewVault(t *testing.T) {
	model := NewModel()

	updated, cmd := model.Update(tea.KeyPressMsg(tea.Key{Text: "n"}))
	if cmd != nil {
		t.Fatal("welcome transition should not return a command")
	}

	model = updated.(*Model)
	if _, ok := model.Active.(*screens.VaultCreationScreen); !ok {
		t.Fatalf("expected new vault screen, got %T", model.Active)
	}
}

func TestNewVaultReturnsToWelcome(t *testing.T) {
	model := NewModel()

	updated, _ := model.Update(tea.KeyPressMsg(tea.Key{Text: "n"}))
	model = updated.(*Model)

	updated, cmd := model.Update(tea.KeyPressMsg(tea.Key{Code: tea.KeyEscape}))
	if cmd != nil {
		t.Fatal("back transition should not return a command")
	}

	model = updated.(*Model)
	if _, ok := model.Active.(*screens.Welcome); !ok {
		t.Fatalf("expected welcome screen, got %T", model.Active)
	}
}

func TestGlobalQuitIsHandledByRootModel(t *testing.T) {
	model := NewModel()

	_, cmd := model.Update(tea.KeyPressMsg(tea.Key{Mod: tea.ModCtrl, Code: 'c'}))
	if cmd == nil {
		t.Fatal("ctrl+c should return a quit command")
	}

	if msg := cmd(); msg != (tea.QuitMsg{}) {
		t.Fatalf("expected quit message, got %T", msg)
	}
}

func TestViewUsesTheAlternateScreen(t *testing.T) {
	model := NewModel()

	model.Update(tea.WindowSizeMsg{Width: 100, Height: 30})
	view := model.View()

	if !view.AltScreen {
		t.Fatal("the TUI should render in the alternate screen")
	}
	if view.Content == "" {
		t.Fatal("the welcome screen should render content")
	}
}

func TestReturningToWelcomePreservesTerminalSize(t *testing.T) {
	model := NewModel()
	model.Update(tea.WindowSizeMsg{Width: 100, Height: 30})

	updated, _ := model.Update(tea.KeyPressMsg(tea.Key{Text: "n"}))
	model = updated.(*Model)
	updated, _ = model.Update(tea.KeyPressMsg(tea.Key{Code: tea.KeyEscape}))
	model = updated.(*Model)

	view := model.View()
	if got := lipgloss.Width(view.Content); got != 100 {
		t.Fatalf("expected welcome view width 100 after returning, got %d", got)
	}
	if got := lipgloss.Height(view.Content); got != 30 {
		t.Fatalf("expected welcome view height 30 after returning, got %d", got)
	}
}
