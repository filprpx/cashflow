package tui

type Model struct {
	KeyMap KeyMap

	ActiveVault string
}

func NewModel() Model {
	return Model{
		KeyMap:      DefaultKeyMap,
		ActiveVault: "",
	}
}
