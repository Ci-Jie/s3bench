package report

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var progressGauge *widgets.Gauge

func initProgressBlock() {
	progressGauge = widgets.NewGauge()
	progressGauge.Title = "Progress Rate"
	progressGauge.SetRect(0, 0, 70, 3)
	progressGauge.Percent = 0
	progressGauge.BarColor = ui.ColorRed
	progressGauge.BorderStyle.Fg = ui.ColorWhite
	progressGauge.TitleStyle.Fg = ui.ColorCyan
}

func updateProgress(percent int) {
	progressGauge.Percent = percent
}
