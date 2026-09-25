package tui

import "charm.land/bubbles/v2/key"

// GlobalKeyMap contains bindings that can be handled by the application shell.
// Screen-specific actions belong to the screen that owns them.
type GlobalKeyMap struct {
	Quit key.Binding
}

// NavigationKeyMap contains actions shared by screens that navigate a list or
// move between controls.
type NavigationKeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Left   key.Binding
	Right  key.Binding
	Select key.Binding
	Back   key.Binding
}

func NewGlobalKeyMap() GlobalKeyMap {
	return GlobalKeyMap{
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "quit"),
		),
	}
}

func NewNavigationKeyMap() NavigationKeyMap {
	return NavigationKeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
		Left: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "move left"),
		),
		Right: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "move right"),
		),
		Select: key.NewBinding(
			key.WithKeys("enter", "space"),
			key.WithHelp("enter/space", "select"),
		),
		Back: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "back"),
		),
	}
}
