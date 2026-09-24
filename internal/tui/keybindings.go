package tui

import "charm.land/bubbles/v2/key"

type KeyMap struct {
	// close application
	Quit key.Binding

	// navigation
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding

	// interaction
	Select key.Binding
}

// App default key bindings
var DefaultKeyMap = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("q", "quit"),
	),

	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move up"),
	),
	Left: key.NewBinding(
		key.WithKeys("", ""),
		key.WithHelp("", ""),
	),
	Right: key.NewBinding(
		key.WithKeys("", ""),
		key.WithHelp("", ""),
	),

	Select: key.NewBinding(
		key.WithKeys("space", "enter"),
		key.WithHelp("space/enter", "select"),
	),
}
