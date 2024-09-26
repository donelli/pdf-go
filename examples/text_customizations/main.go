package main

import (
	tpdf "tpdf"
)

func main() {
	content := tpdf.Column(
		tpdf.Text("Red with font size 20").WithColor(tpdf.HexToRGBA("#ff0000")).WithFontSize(20),
		tpdf.Text("Bold text").InBold(),
	).WithSpacing(8)

	generator := tpdf.NewGenerator()
	generator.SetMargins(8, 8, 8, 8)
	generator.SetMainWidget(content)

	generator.GenerateToFile("text_customizations.pdf")
}
