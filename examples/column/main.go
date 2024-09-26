package main

import (
	"strings"
	"tpdf"
)

func main() {
	content := tpdf.Column(
		tpdf.Text(strings.Repeat("Hello ", 20), tpdf.FontSize(14)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
		tpdf.Text("Hello", tpdf.FontSize(20)),
	)

	generator := tpdf.NewGenerator()

	generator.SetMainWidget(content)
	generator.GenerateToFile("column.pdf")
}
