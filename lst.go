package main

import (
	//"fmt"
	//"strconv"
	"os"
	"time"

	g "github.com/AllenDang/giu"
)

var content string
var t0 time.Time
var t float64
var text, s string
var filename = "timecode.txt"

func plusUn() {
	t0 = t0.Add(-time.Second)
}

func moinsUn() {
	t0 = t0.Add(time.Second)
}

func Raz() {
	t0 = time.Now()
}

func Newline() {
	text = text + s + " " + content + "\n"
	content = ""
}

func Savetofile() {
	f, _ := os.Create("timecode.txt")
	defer f.Close()
	f.WriteString(text)
}
func loop() {
	t = time.Since(t0).Seconds()
	e := time.Duration(int(t)) * time.Second
	t1 := time.Unix(int64(e.Seconds()), int64(t))
	s = t1.Format("03:04:05")
	g.SingleWindow().Layout(
		g.Row(
			g.Label(s),
			g.Button("+1").OnClick(plusUn),
			g.Button("-1").OnClick(moinsUn),
			g.Button("raz").OnClick(Raz),
		),
		g.InputText(&content).Size(g.Auto),
		g.Event().OnKeyPressed(g.KeyEnter, Newline),
		g.Event().OnKeyReleased(g.KeyEnter, g.SetKeyboardFocusHere),
		g.Button("New").OnClick(Newline),
		g.InputTextMultiline(&text).Size(g.Auto, g.Auto-50),
		g.Row(
			g.Button("Save").OnClick(Savetofile),
			g.InputText(&filename),
		),
	)
	g.Update()

}

func main() {
	t0 = time.Now()
	wnd := g.NewMasterWindow("Live Shill Timecode", 500, 350, g.MasterWindowFlagsFloating)
	wnd.Run(loop)
}
