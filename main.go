package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(500, 280))

	output := ""
	input := widget.NewLabel(output)

	isHistory := false
	historyStr := ""
	history := widget.NewLabel(historyStr)

	var historyArr []string;

	// Hitsory button
	historyBtn := widget.NewButton("History", func() {
		if isHistory{
			historyStr = ""
		}else{
			for i := len(historyArr)-1; i>=0; i--{
				historyStr = historyStr + historyArr[i]
				historyStr+="\n"
			}
		}
		isHistory = !isHistory
		history.SetText(historyStr)
	})

	// Back button
	backBtn := widget.NewButton("Back", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})

	// Clear button
	clearBtn := widget.NewButton("Clear", func() {
		output = ""
		input.SetText(output)
	})

	// Open button
	openBtn := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})

	// Close button
	closeBtn := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})

	// Divide button
	divideBtn := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})

	// 7 button
	sevenBtn := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})

	// Eight button
	eightBtn := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})

	// Nine button
	nineBtn := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})

	// Multiply button
	multiplyBtn := widget.NewButton("*", func() {
		output = output + "x"
		input.SetText(output)
	})

	// Four button
	fourBtn := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})

	// Five button
	fiveBtn := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})

	// Six button
	sixBtn := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})

	// Minus button
	minusBtn := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})

	// One button
	oneBtn := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})

	// Two button
	twoBtn := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})

	// Three button
	threeBtn := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})

	// Plus button
	plusBtn := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})

	// Zero button
	zeroBtn := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})

	// Dot button
	dotBtn := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})

	// Equal button
	equalBtn := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAppend := output+" = "+ans;
				historyArr = append(historyArr, strToAppend)
				output = ans 
			} else {
				output = "ERROR !!"
			}
		} else {
			output = "ERROR !!"
		}

		input.SetText(output)
	})

	w.SetContent(container.NewVBox(
		input,
		history,
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
