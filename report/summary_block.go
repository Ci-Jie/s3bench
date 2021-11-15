package report

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var summaryTable *widgets.Table

func initSummaryTable() {
	summaryTable = widgets.NewTable()
	summaryTable.Rows = [][]string{
		{"Average Speed", "Total Count"},
		{"-", "-"},
	}
	summaryTable.TextStyle = ui.NewStyle(ui.ColorWhite)
	summaryTable.RowSeparator = false
	summaryTable.BorderStyle = ui.NewStyle(ui.ColorGreen)
	summaryTable.SetRect(0, len(SummaryHandler.Workers)+7, 70, len(SummaryHandler.Workers)+11)
	summaryTable.FillRow = true
	summaryTable.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
}

func updateSummaryTable(avgSpeed, totalSize string) {
	summaryTable.Rows[1] = []string{avgSpeed, totalSize}
}
