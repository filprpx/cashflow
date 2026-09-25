package screens

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type WelcomeKeyMap struct {
	NewVault key.Binding
	Select   key.Binding
}

type Welcome struct {
	keys    WelcomeKeyMap
	actions list.Model
	help    help.Model
	width   int
	height  int
}

func NewWelcomeScreen() *Welcome {
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = true
	delegate.SetSpacing(1)
	delegate.Styles = list.NewDefaultItemStyles(true)
	delegate.Styles.NormalTitle = lipgloss.NewStyle().Foreground(lipgloss.Color("#D8DEE9"))
	delegate.Styles.NormalDesc = lipgloss.NewStyle().Foreground(lipgloss.Color("#8290A5"))
	delegate.Styles.SelectedTitle = lipgloss.NewStyle().Foreground(lipgloss.Color("#8BD5CA")).Bold(true)
	delegate.Styles.SelectedDesc = lipgloss.NewStyle().Foreground(lipgloss.Color("#A7B6C9"))

	actions := list.New([]list.Item{
		welcomeAction{
			id:          "new-vault",
			title:       "Create your first vault",
			description: "Start with an encrypted local workspace",
		},
		welcomeAction{
			id:          "about",
			title:       "About Cashflow",
			description: "Personal financial data control, locally owned",
		},
	}, delegate, 52, 7)
	actions.Title = ""
	actions.SetShowTitle(false)
	actions.SetShowFilter(false)
	actions.SetShowStatusBar(false)
	actions.SetShowPagination(false)
	actions.SetShowHelp(false)

	return &Welcome{
		keys: WelcomeKeyMap{
			NewVault: key.NewBinding(
				key.WithKeys("n"),
				key.WithHelp("n", "new vault"),
			),
			Select: key.NewBinding(
				key.WithKeys("enter", "space"),
				key.WithHelp("enter/space", "select"),
			),
		},
		actions: actions,
		help:    help.New(),
	}
}

func (m *Welcome) Init() tea.Cmd {
	return nil
}

func (m *Welcome) Update(msg tea.Msg) (Screen, tea.Cmd) {
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = size.Width
		m.height = size.Height
		m.actions.SetSize(m.listWidth(), m.listHeight())
	}

	var cmd tea.Cmd
	m.actions, cmd = m.actions.Update(msg)

	keyMsg, ok := msg.(tea.KeyPressMsg)
	if ok && key.Matches(keyMsg, m.keys.NewVault) {
		return NewVaultCreationScreen(), nil
	}
	if ok && key.Matches(keyMsg, m.keys.Select) {
		if action, ok := m.actions.SelectedItem().(welcomeAction); ok && action.id == "new-vault" {
			return NewVaultCreationScreen(), nil
		}
	}

	return m, cmd
}

func (m *Welcome) View() tea.View {
	content := m.render()
	return tea.NewView(content)
}

type welcomeAction struct {
	id          string
	title       string
	description string
}

func (a welcomeAction) FilterValue() string {
	return a.title + " " + a.description
}

func (a welcomeAction) Title() string {
	return a.title
}

func (a welcomeAction) Description() string {
	return a.description
}

func (m *Welcome) ShortHelp() []key.Binding {
	return []key.Binding{m.keys.NewVault, m.keys.Select}
}

func (m *Welcome) FullHelp() [][]key.Binding {
	return [][]key.Binding{{m.keys.NewVault, m.keys.Select}}
}

func (m *Welcome) listWidth() int {
	width := m.width - 12
	if width < 32 {
		return 32
	}
	if width > 58 {
		return 58
	}
	return width
}

func (m *Welcome) listHeight() int {
	return 5
}

func (m *Welcome) render() string {
	const defaultWidth = 80
	const defaultHeight = 24

	width := m.width
	if width == 0 {
		width = defaultWidth
	}
	height := m.height
	if height == 0 {
		height = defaultHeight
	}

	accent := lipgloss.Color("#8BD5CA")
	muted := lipgloss.Color("#8290A5")
	text := lipgloss.Color("#D8DEE9")
	border := lipgloss.Color("#314158")

	header := lipgloss.NewStyle().
		Foreground(accent).
		Bold(true).
		Render("CASHFLOW")
	subtitleText := "PERSONAL FINANCIAL CONTROL"
	if width < 64 {
		subtitleText = "LOCAL-FIRST FINANCE"
	}
	subtitle := lipgloss.NewStyle().Foreground(muted).Render(subtitleText)
	headerRow := lipgloss.JoinHorizontal(lipgloss.Left, header, "  ", subtitle)

	heroTitle := lipgloss.NewStyle().Foreground(text).Bold(true).Render("Your money. Under your control.")
	heroCopy := lipgloss.NewStyle().Foreground(muted).Width(m.listWidth()).Render(
		"Connect your financial systems, preserve your data locally, and decide where it goes.",
	)

	panelStyle := lipgloss.NewStyle().
		Width(m.listWidth()).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(border)
	menu := panelStyle.Render(m.actions.View())

	helpView := m.help.ShortHelpView(m.ShortHelp())
	footer := lipgloss.NewStyle().Foreground(muted).Render(helpView + "    ctrl+c quit")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		headerRow,
		"",
		heroTitle,
		heroCopy,
		"",
		menu,
		"",
		footer,
	)

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
}
