package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	fmt.Println("test fyne")

	a := app.New()

	w := a.NewWindow("title") //this is title
	w.Resize(fyne.NewSize(200,400))

	w.ShowAndRun()// Finally Running our App
}