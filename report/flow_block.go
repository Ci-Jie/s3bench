package report

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var flowLine *widgets.Sparkline
var flowLineGroup *widgets.SparklineGroup

func initFlowBlock() {
	flowLine = widgets.NewSparkline()
	flowLine.Data = []float64{}
	flowLine.LineColor = ui.ColorRed

	flowLineGroup = widgets.NewSparklineGroup(flowLine)
	flowLineGroup.SetRect(70, 0, 140, len(SummaryHandler.Workers)+11)
	flowLineGroup.BorderStyle.Fg = ui.ColorCyan
}

func updateFlowBlock(data []float64) {
	flowLine.Data = data
}

func getFlowBlockData() (data []float64) {
	return flowLine.Data
}
