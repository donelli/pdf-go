package main

import (
	tpdf "tpdf"
)

func main() {
	content := tpdf.Column(
		tpdf.Row(
			tpdf.Container(nil).
				WithSize(100, 100).
				WithBackgroundColor(tpdf.HexToRGBA("#ff0000")),

			tpdf.Container(nil).
				WithSize(50, 50).
				WithBackgroundColor(tpdf.HexToRGBA("#0000ff")),
		).
			WithMainAxisSize(tpdf.MainAxisSizeMin).
			WithSpacing(8),

		tpdf.Container(nil).
			WithSize(100, 20).
			WithBackgroundColor(tpdf.HexToRGBA("#00ff00")),

		tpdf.Container(
			tpdf.Text("Container with child and padding"),
		).
			WithBackgroundColor(tpdf.HexToRGBA("#bbffbb")).
			PaddingAll(8),

		tpdf.Container(
			tpdf.Text("Container with child and padding and width"),
		).
			WithWidth(100).
			WithBackgroundColor(tpdf.HexToRGBA("#dddddd")).
			PaddingAll(8),

		tpdf.Container(
			tpdf.Text("A"),
		).
			WithSize(40, 40).
			WithBackgroundColor(tpdf.HexToRGBA("#dddddd")).
			BorderRadius(tpdf.BorderRadiusAll(4)),

		tpdf.Container(
			tpdf.Text("Border radius in specific corders"),
		).
			WithBackgroundColor(tpdf.HexToRGBA("#ff00ff")).
			BorderRadius(tpdf.BorderRadiusEach(4, 0, 0, 4)).
			PaddingAll(2),

		tpdf.Container(
			tpdf.Text("Bordered"),
		).
			PaddingAll(4).
			Bordered(tpdf.HexToRGBA("#00ff00")),

		tpdf.Container(
			tpdf.Text("Rounded Bordered"),
		).
			PaddingAll(4).
			BorderRadius(tpdf.BorderRadiusAll(4)).
			Bordered(tpdf.HexToRGBA("#00ff00")),
	).WithSpacing(8)

	generator := tpdf.NewGenerator()
	generator.SetMargins(8, 8, 8, 8)
	generator.SetMainWidget(content)

	generator.GenerateToFile("container.pdf")
}
