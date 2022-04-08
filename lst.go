package main

import (
	"fmt"
	"os"
	"time"

	g "github.com/AllenDang/giu"
)

var (
	t0                       time.Time
	content, text, timestamp string
	filename                 = "timecode.txt"
	ticker                   *time.Ticker
	auto                     bool
)

func plusOne() {
	t0 = t0.Add(time.Second)
}

func minusOne() {
	t0 = t0.Add(-time.Second)
}

func resetTicker() {
	t0, _ = time.Parse("05", "0")
}

func startTicker() {
	ticker.Reset(time.Second)
}

func stopTicker() {
	ticker.Stop()
}

func newLine() {
	text = text + timestamp + " " + content + "\n"
	content = ""
}

func saveTofile() {
	f, _ := os.Create("timecode.txt")
	defer f.Close()
	f.WriteString(text)
}

func tickerUpdate() {
	for {
		<-ticker.C
		t0 = t0.Add(time.Second)
		g.Update()
	}
}

func loop() {
	timestamp = t0.Format("15:04:05")
	g.SingleWindow().Layout(
		g.Row(
			g.Label(timestamp),
			g.Button("+1").OnClick(plusOne),
			g.Button("-1").OnClick(minusOne),
			g.Button("start").OnClick(startTicker),
			g.Button("stop").OnClick(stopTicker),
			g.Button("reset").OnClick(resetTicker),
			g.Checkbox("auto", &auto),
		),
		/*g.Custom(func() {
			if auto {
				g.SetKeyboardFocusHere()
			}
		}),*/
		g.InputText(&content).Size(g.Auto),
		g.Event().OnKeyPressed(g.KeyEnter, newLine),
		g.Button("New").OnClick(newLine),
		g.InputTextMultiline(&text).Size(g.Auto, g.Auto-30),
		g.Row(
			g.Button("Save").OnClick(saveTofile),
			g.InputText(&filename),
		),
	)
	fmt.Println(auto)
}

func main() {
	ticker = time.NewTicker(time.Second)
	go tickerUpdate()
	wnd := g.NewMasterWindow("Live Shill Timecode", 500, 500, g.MasterWindowFlagsFloating)
	wnd.Run(loop)
}
