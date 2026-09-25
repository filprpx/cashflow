package tui

import (
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

func TestGlobalKeyMapOnlyContainsUnconditionalQuit(t *testing.T) {
	keys := NewGlobalKeyMap()

	if !key.Matches(tea.KeyPressMsg(tea.Key{Mod: tea.ModCtrl, Code: 'c'}), keys.Quit) {
		t.Fatal("ctrl+c should match the global quit binding")
	}

	if key.Matches(tea.KeyPressMsg(tea.Key{Text: "q"}), keys.Quit) {
		t.Fatal("q should not match the global quit binding")
	}
}

func TestNavigationKeyMapSupportsViAndArrowKeys(t *testing.T) {
	keys := NewNavigationKeyMap()

	tests := []struct {
		name string
		msg  tea.KeyPressMsg
		want key.Binding
	}{
		{name: "arrow up", msg: tea.KeyPressMsg(tea.Key{Code: tea.KeyUp}), want: keys.Up},
		{name: "vi up", msg: tea.KeyPressMsg(tea.Key{Text: "k"}), want: keys.Up},
		{name: "arrow down", msg: tea.KeyPressMsg(tea.Key{Code: tea.KeyDown}), want: keys.Down},
		{name: "vi down", msg: tea.KeyPressMsg(tea.Key{Text: "j"}), want: keys.Down},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !key.Matches(test.msg, test.want) {
				t.Fatalf("%s should match its navigation binding", test.name)
			}
		})
	}
}
