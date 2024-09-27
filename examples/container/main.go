package main

import (
	tpdf "tpdf"
)

func main() {
	content := tpdf.Column(
		tpdf.Row(
			tpdf.Container().WithSize(100, 100).WithBackgroundColor(tpdf.HexToRGBA("#ff0000")),
			tpdf.Container().WithSize(50, 50).WithBackgroundColor(tpdf.HexToRGBA("#0000ff")),
		).WithMainAxisSize(tpdf.MainAxisSizeMin).WithSpacing(8),
		tpdf.Container().WithSize(100, 20).WithBackgroundColor(tpdf.HexToRGBA("#00ff00")),
	).WithSpacing(8)

	generator := tpdf.NewGenerator()
	generator.SetMargins(8, 8, 8, 8)
	generator.SetMainWidget(content)

	generator.GenerateToFile("container.pdf")
}
