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

	generator := tpdf.NewGenerator()

	generator.SetMainWidget(content)
	generator.GenerateToFile("column.pdf")
}
