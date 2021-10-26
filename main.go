package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")

	// Hitsory button 
	historyBtn := widget.NewButton("History", func() {

	})

	// Back button
	backBtn := widget.NewButton("Back", func() {

	})

	// Clear button
	clearBtn := widget.NewButton("Clear", func() {
		
	})

	// Open button
	openBtn := widget.NewButton("(", func() {
		
	})

	// Close button
	closeBtn := widget.NewButton(")", func() {
		
	})

	// Divide button
	divideBtn := widget.NewButton("/", func() {
		
	})

	// 7 button
	sevenBtn := widget.NewButton("7", func() {
		
	})

	// Eight button
	eightBtn := widget.NewButton("8", func() {
		
	})

	// Nine button
	nineBtn := widget.NewButton("9", func() {
		
	})

	// Multiply button
	multiplyBtn := widget.NewButton("*", func() {
		
	})

	// Four button
	fourBtn := widget.NewButton("4", func() {
		
	})

	// Five button
	fiveBtn := widget.NewButton("5", func() {
		
	})

	// Six button
	sixBtn := widget.NewButton("6", func() {
		
	})

	// Minus button
	minusBtn := widget.NewButton("-", func() {
		
	})

	// One button
	oneBtn := widget.NewButton("1", func() {
		
	})

	// Two button
	twoBtn := widget.NewButton("2", func() {
		
	})

	// Three button
	threeBtn := widget.NewButton("3", func() {
		
	})

	// Plus button
	plusBtn := widget.NewButton("+", func() {
		
	})

	// Zero button
	zeroBtn := widget.NewButton("0", func() {
		
	})

	// Dot button
	dotBtn := widget.NewButton(".", func() {
		
	})

	// Equal button
	equalBtn := widget.NewButton("=", func() {
		
	})

	input := widget.NewLabel("")
	w.SetContent(container.NewVBox(
		input,
		container.NewGridWithColumns(1, 
			container.NewGridWithColumns(2,
			historyBtn,
			backBtn,
			),
			container.NewGridWithColumns(4,
			clearBtn,
			openBtn,
			closeBtn,
			divideBtn,
			),
			container.NewGridWithColumns(4,
			nineBtn,
			eightBtn,
			sevenBtn,
			multiplyBtn,
			),
			container.NewGridWithColumns(4,
			fourBtn,
			fiveBtn,
			sixBtn,
			minusBtn,
			),
			container.NewGridWithColumns(4,
			oneBtn,
			twoBtn,
			threeBtn,
			plusBtn,
			),
			container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
			zeroBtn,
			dotBtn,
			),
			equalBtn,
			),

		),
	))

	w.ShowAndRun()
}