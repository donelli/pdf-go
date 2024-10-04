package main

import (
	"strings"
	"tpdf"
)

func main() {
	content := tpdf.Column(
		tpdf.Text(strings.Repeat("Hello ", 20)),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Text("Hello").WithFontSize(20),
		tpdf.Column(
			tpdf.Text("With spacing"),
			tpdf.Text("With spacing"),
			tpdf.Text("With spacing"),
		).WithSpacing(12),
	)

	theme := tpdf.NewTheme()
	writer := tpdf.NewWriter(8, 8, 8, 8, theme)

	writer.SetMainWidget(content)

	writer.GenerateToFile("column.pdf")
}
