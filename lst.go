package main

import (
	"os"
	"time"

	g "github.com/AllenDang/giu"
)

var (
	t0                       time.Time
	content, text, timestamp string
	filename                 = "timecode.txt"
	ticker                   *time.Ticker
	edit                     bool = false
	autosavebool             bool = true
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
	f, _ := os.Create(filename)
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

func autoSave() {
	for {
		if autosavebool {
			go saveTofile()
		}
		time.Sleep(30 * time.Second)
	}
}

func loop() {
	timestamp = t0.Format("15:04:05")
	g.SingleWindow().RegisterKeyboardShortcuts(
		g.WindowShortcut{Key: g.KeyE, Modifier: g.ModControl, Callback: func() { edit = !edit }},
		g.WindowShortcut{Key: g.KeyS, Modifier: g.ModControl, Callback: saveTofile},
		g.WindowShortcut{Key: g.KeyF1, Modifier: g.ModControl, Callback: minusOne},
		g.WindowShortcut{Key: g.KeyF2, Modifier: g.ModControl, Callback: plusOne},
	).Layout(
		g.Row(
			g.Label(timestamp),
			g.Button("+1 (F1)").OnClick(plusOne),
			g.Button("-1 (F2)").OnClick(minusOne),
			g.Button("start").OnClick(startTicker),
			g.Button("stop").OnClick(stopTicker),
			g.Button("reset").OnClick(resetTicker),
			g.Checkbox("keyboard only (crtl-e)", &edit),
		),
		g.Custom(func() {
			if edit {
				g.SetKeyboardFocusHere()
			}
		}),
		g.Row(
			g.InputText(&content).Size(g.Auto-150),
			g.Event().OnKeyPressed(g.KeyEnter, newLine),
			g.Button("New comment").OnClick(newLine),
		),
		g.InputTextMultiline(&text).Size(g.Auto, g.Auto-30),
		g.Row(
			g.Button("Save (ctrl-s)").OnClick(saveTofile),
			g.InputText(&filename),
			g.Checkbox("autosave", &autosavebool),
		),
	)
}

func main() {
	ticker = time.NewTicker(time.Second)
	go tickerUpdate()
	go autoSave()
	wnd := g.NewMasterWindow("Live Shill Timecode", 500, 500, g.MasterWindowFlagsFloating)
	wnd.Run(loop)
}
