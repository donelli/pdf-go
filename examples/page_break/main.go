package main

import (
	tpdf "tpdf"
)

func main() {
	content := tpdf.Column(
		tpdf.Text("First page"),
		tpdf.PageBreak(),
		tpdf.Text("Second page"),
		tpdf.PageBreak(),
		tpdf.Text("Third page"),
	).WithSpacing(8)

	theme := tpdf.NewTheme()
	theme.DefaultDividerColor = tpdf.HexToRGBA("#dddddd")
	theme.DefaultDividerCapStyle = tpdf.DividerCapStyleRound
	writer := tpdf.NewWriter(8, 8, 8, 8, theme)

	writer.SetMainWidget(content)

	writer.GenerateToFile("page_break.pdf")
}
