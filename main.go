package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func printarr(in map[string]string) {
	fmt.Printf("----------------\n")
	for k, v := range in {
		fmt.Printf("%s : %s\n", v, k)
	}
	fmt.Printf("----------------\n")
}

func main() {
	a := app.New()
	w := a.NewWindow("MFA Tool")
	countdown := widget.NewLabel("Countdown")
	box := container.NewVBox(countdown)

	cfg := OtpConfig{}
	cfg.init("/Users/jblack/.otps")
	fmt.Println("---------")
	for _, o := range cfg.otps.get() {
		label := fmt.Sprintf("%s\n%s", o.name, o.token)
		box.Add(widget.NewLabel(label))
	}
	box.Add(widget.NewButton("quit", func() { a.Quit() }))
	w.SetContent(box)
	w.ShowAndRun()
}
