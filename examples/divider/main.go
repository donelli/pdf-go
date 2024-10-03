package main

import (
	tpdf "tpdf"
)

func main() {
	content := tpdf.Column(
		tpdf.Text("Normal divider"),
		tpdf.Divider(),
		tpdf.Text("Red divider"),
		tpdf.Divider().Color(tpdf.HexToRGBA("#ff0000")),
		tpdf.Text("Green divider with 3 as line height and square cap style"),
		tpdf.Divider().Color(tpdf.HexToRGBA("#00ff00")).LineHeight(3).CapStyle(tpdf.DividerCapStyleSquare),
		tpdf.Container(
			tpdf.Row(
				tpdf.Text("A"),
				tpdf.VerticalDivider(),
				tpdf.Text("B"),
				tpdf.VerticalDivider().Color(tpdf.HexToRGBA("#ff0000")).CapStyle(tpdf.DividerCapStyleSquare),
				tpdf.Text("C"),
				tpdf.VerticalDivider().Color(tpdf.HexToRGBA("#00ff00")).LineWidth(3),
			),
		).WithHeight(100),
	).WithSpacing(8)

	generator := tpdf.NewGenerator()
	generator.SetMargins(8, 8, 8, 8)
	generator.SetMainWidget(content)

	generator.SetDefaultLineColor(tpdf.HexToRGBA("#dddddd"))
	generator.SetDefaultCapStyle(tpdf.DividerCapStyleRound)

	generator.GenerateToFile("divider.pdf")
}
