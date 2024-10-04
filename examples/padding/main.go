package main

import (
	"math"
	tpdf "tpdf"
)

func main() {
	content := tpdf.Column(
		tpdf.Container(
			tpdf.Padding(
				tpdf.PaddingAll(8),
				tpdf.Text("Padding all 8"),
			),
		).
			WithBackgroundColor(tpdf.HexToRGBA("#ff0000")).
			WithSize(120, 120),
		tpdf.Container(
			tpdf.Padding(
				tpdf.PaddingEach(8, 8, 8, 8),
				tpdf.Container(
					tpdf.Padding(
						tpdf.PaddingEach(8, 8, 8, 8),
						tpdf.Container(
							nil,
						).
							WithBackgroundColor(tpdf.HexToRGBA("#E5FF00")).
							WithSize(math.MaxFloat64, math.MaxFloat64),
					),
				).
					WithBackgroundColor(tpdf.HexToRGBA("#0ff000")).
					WithSize(math.MaxFloat64, math.MaxFloat64),
			),
		).
			WithBackgroundColor(tpdf.HexToRGBA("#ff0000")).
			WithSize(120, 120),
	).WithSpacing(8)

	theme := tpdf.NewTheme()
	writer := tpdf.NewWriter(8, 8, 8, 8, theme)

	writer.SetMainWidget(content)

	writer.GenerateToFile("padding.pdf")
}
