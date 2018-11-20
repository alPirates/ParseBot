package main

import (
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/extra"
)

type UiService struct {
	Tabpane *extra.Tabpane
	// main layout
	Layout1  []*ui.Row
	Commands *ui.List
	Logs     *ui.List
	Status   *ui.BarChart
	// config layout
	Layout2        []*ui.Row
	ConfigCommands *ui.List
	Config         *ui.List
	ConfigStatus   *ui.Par
}

func NewUiService() UiService {
	err := ui.Init()
	if err != nil {
		panic(err)
	}

	tab1 := extra.NewTab("Статистика")
	tab2 := extra.NewTab("Конфигурация")
	tabpane := extra.NewTabpane()
	tabpane.Y = 1
	tabpane.Width = 20
	tabpane.SetTabs(*tab1, *tab2)

	listCommands := ui.NewList()
	strs := []string{
		"[q] [Выход](fg-red)",
		"[u] [Обновление состояния](fg-blue)",
		"[t] [Тестовое письмо](fg-red)",
		"---------------------",
		"[1] [Главная панель](fg-green)",
		"[2] [Конфигурация](fg-yellow)",
	}
	listCommands.Items = strs
	listCommands.ItemFgColor = ui.ColorYellow
	listCommands.BorderLabel = "Возможные комманды"
	listCommands.Height = 10
	listCommands.Width = 25
	listCommands.Y = 0

	bc := ui.NewBarChart()
	bc.Data = []int{87, 54, 100, 50, 88}
	bc.DataLabels = []string{"S1", "S2", "S3", "S4", "S5"}
	bc.BorderLabel = "Состояние"
	bc.Width = 26
	bc.Height = 10
	bc.TextColor = ui.ColorGreen
	bc.BarColor = ui.ColorRed
	bc.NumColor = ui.ColorYellow

	lg := ui.NewList()
	logs := []string{
		"[Email] [шлю тестовый email](fg-red)",
		"[Parsing] [Запаршено 18:32](fg-blue)",
		"[TEMP] [Температура очень высокая](fg-red)",
	}
	lg.Items = logs
	lg.ItemFgColor = ui.ColorYellow
	lg.BorderLabel = "Логи"
	lg.Height = 15
	lg.Width = 25
	lg.Y = 0

	configCommands := ui.NewList()
	commandsList := []string{
		"[r] [Обновить конфигурацию](fg-red)",
		"---------------------",
		"[1] [Главная панель](fg-green)",
		"[2] [Конфигурация](fg-yellow)",
	}
	configCommands.Items = commandsList
	configCommands.ItemFgColor = ui.ColorYellow
	configCommands.BorderLabel = "Меню"
	configCommands.Height = 10
	configCommands.Width = 25
	configCommands.Y = 0

	config := ui.NewList()
	configList := []string{
		"[email] [test@mail.ru](fg-green)",
		"[password] [test](fg-yellow)",
	}
	config.Items = configList
	config.ItemFgColor = ui.ColorCyan
	config.BorderLabel = "Конфигурация"
	config.Height = 10
	config.Width = 25
	config.Y = 0

	configStatus := ui.NewPar("[Конфигурация прочитана успешно](fg-green)")
	configStatus.Height = 3
	configStatus.Width = 37
	configStatus.Y = 4
	configStatus.BorderFg = ui.ColorGreen

	return UiService{
		Tabpane:        tabpane,
		Commands:       listCommands,
		Status:         bc,
		Logs:           lg,
		ConfigCommands: configCommands,
		Config:         config,
		ConfigStatus:   configStatus,
	}
}

func (u UiService) Init() {
	layout1 := []*ui.Row{
		ui.NewRow(
			ui.NewCol(12, 0, u.Tabpane),
		),
		ui.NewRow(
			ui.NewCol(6, 0, u.Commands),
			ui.NewCol(6, 0, u.Status),
		),
		ui.NewRow(
			ui.NewCol(12, 0, u.Logs),
		),
	}

	layout2 := []*ui.Row{
		ui.NewRow(
			ui.NewCol(12, 0, u.Tabpane),
		),
		ui.NewRow(
			ui.NewCol(12, 0, u.ConfigCommands),
		),
		ui.NewRow(
			ui.NewCol(12, 0, u.Config),
		),
		ui.NewRow(
			ui.NewCol(12, 0, u.ConfigStatus),
		),
	}

	u.Layout1 = layout1
	u.Layout2 = layout2

	u.SetMainLayout()

	// checking keys
	u.CheckKeys()
	ui.Loop()
}

func (u UiService) SetMainLayout() {
	ui.Clear()
	ui.Body.Rows = u.Layout1
	ui.Body.Align()
	u.Tabpane.SetActiveLeft()
	ui.Render(ui.Body, u.Tabpane)
}

func (u UiService) SetConfigLayout() {
	ui.Clear()
	ui.Body.Rows = u.Layout2
	ui.Body.Align()
	u.Tabpane.SetActiveRight()
	ui.Render(ui.Body, u.Tabpane)
}

func (u UiService) CheckKeys() {
	ui.Handle("q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("u", func(ui.Event) {
		// обновление состояния
	})
	ui.Handle("r", func(ui.Event) {
		// обновление конфигурации
	})
	ui.Handle("t", func(ui.Event) {
		// тестовое  письмо
		if len(u.Logs.Items) == 7 {
			u.Logs.Items = u.Logs.Items[1:]
		}
		u.Logs.Items = append(u.Logs.Items, "[Email] [шлю тестовый email](fg-red)")
		ui.Render(u.Logs)
	})

	// switch layout
	ui.Handle("1", func(ui.Event) {
		u.SetMainLayout()
	})
	ui.Handle("2", func(ui.Event) {
		u.SetConfigLayout()
	})
}
