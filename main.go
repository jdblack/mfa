package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func tick(clock binding.Float) {
	for {
		time.Sleep(100 * time.Millisecond)
		var left float64
		left = float64(29 - (time.Now().Unix() % 30))
		clock.Set(left / 30.0)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("MFA Tool")
	r, _ := fyne.LoadResourceFromURLString("https://picsum.photos/200")
	w.SetIcon(r)

	clock := binding.NewFloat()
	go tick(clock)

	countdown := widget.NewProgressBarWithData(clock)

	box := container.NewVBox(countdown)

	cfg := OtpConfig{}
	cfg.init("/Users/jblack/.otps")
	for _, o := range cfg.otps.get() {
		label := canvas.NewText(o.name, color.Black)
		label.TextStyle = fyne.TextStyle{Bold: true}

		box.Add(label)
		box.Add(widget.NewLabelWithData(o.token))
		box.Add(layout.NewSpacer())
	}
	box.Add(widget.NewButton("Quit", func() { a.Quit() }))
	w.SetContent(box)
	w.ShowAndRun()
}
