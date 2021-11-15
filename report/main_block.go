package report

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var mainTable *widgets.Table

func initMainBlock() {
	mainTable = widgets.NewTable()
	mainTable.Rows = [][]string{
		{"Workers", "Speed", "Failed Rate", "Total Size"},
	}
	mainTable.TextStyle = ui.NewStyle(ui.ColorWhite)
	mainTable.RowSeparator = false
	mainTable.BorderStyle = ui.NewStyle(ui.ColorGreen)
	mainTable.SetRect(0, 3, 70, len(SummaryHandler.Workers)+7)
	mainTable.FillRow = true
	mainTable.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
	for i := 0; i < len(SummaryHandler.Workers); i++ {
		mainTable.Rows = append(mainTable.Rows, []string{
			fmt.Sprintf("worker-%d", i), "", "", "",
		})
	}
	mainTable.Rows = append(mainTable.Rows, []string{
		"Total", "-", "-", "-",
	})
	mainTable.RowStyles[len(mainTable.Rows)-1] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
}

func updateMainBlock(data [][]string) {
	for i := 0; i < len(data); i++ {
		mainTable.Rows[i+1] = data[i]
	}
}
