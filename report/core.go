package report

import (
	"fmt"
	"s3bench/object"
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
)

// Summary ...
type Summary struct {
	Workers    []*Worker
	Duration   int32
	TotalSpeed float64
	TotalCount int
	Stop       bool
}

// Worker ...
type Worker struct {
	Name      string
	Endpoint  string
	Object    *object.Object
	Speed     float64
	ErrorRate string
	TotalSize int
	Count     int
}

// SummaryHandler ...
var SummaryHandler *Summary

// Start ...
func Start(count *int) {
	initLayout()
	defer closeLayout()

	update(count)
	ui.Render(mainTable, progressGauge, flowLineGroup, summaryTable)
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second)
	for {
		if percent >= 100 {
			ticker.Stop()
			SummaryHandler.Stop = true
			avgSpeed := SummaryHandler.TotalSpeed / float64(*count)
			totalCount := SummaryHandler.TotalCount
			updateSummaryTable(fmt.Sprintf("%s/s", unitConvert(avgSpeed)), strconv.Itoa(totalCount))
			ui.Render(mainTable, progressGauge, flowLineGroup, summaryTable)
		}
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker.C:
			update(count)
			ui.Render(mainTable, progressGauge, flowLineGroup, summaryTable)
		}
	}
}

func update(count *int) {
	var totalSpeed float64 = 0
	var totalSize int
	var totalCount int
	var mainData [][]string
	for _, worker := range SummaryHandler.Workers {
		mainData = append(mainData, []string{
			worker.Name,
			fmt.Sprintf("%s/s", unitConvert(worker.Speed)),
			worker.ErrorRate,
			unitConvert(float64(worker.TotalSize)),
		})
		totalSpeed += worker.Speed
		totalSize += worker.TotalSize
		totalCount += worker.Count
	}
	mainData = append(mainData, []string{
		"Total",
		fmt.Sprintf("%s/s", unitConvert(totalSpeed)),
		"-",
		unitConvert(float64(totalSize)),
	})
	updateMainBlock(mainData)
	percent = int(float32(*count) / float32(SummaryHandler.Duration) * 100)
	updateProgress(percent)
	SummaryHandler.TotalSpeed += totalSpeed
	SummaryHandler.TotalCount = totalCount
	updateFlowBlock(append([]float64{totalSpeed}, getFlowBlockData()...))
}

func unitConvert(input float64) (output string) {
	switch {
	case input > 1048576:
		return fmt.Sprintf("%.2f MB", input/1048576)
	case input > 1024:
		return fmt.Sprintf("%.2f KB", input/1024)
	default:
		return fmt.Sprintf("%.2f bytes", input)
	}
}
