package report

import (
	"log"

	ui "github.com/gizak/termui/v3"
)

var percent int

// initLayout ...
func initLayout() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	initMainBlock()
	initProgressBlock()
	initSummaryTable()
	initFlowBlock()
}

func closeLayout() {
	ui.Close()
}
