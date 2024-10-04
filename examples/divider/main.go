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

	theme := tpdf.NewTheme()
	theme.DefaultDividerColor = tpdf.HexToRGBA("#dddddd")
	theme.DefaultDividerCapStyle = tpdf.DividerCapStyleRound
	writer := tpdf.NewWriter(8, 8, 8, 8, theme)

	writer.SetMainWidget(content)

	writer.GenerateToFile("divider.pdf")
}
