package main

import (
	tpdf "tpdf"
)

func main() {
	content := tpdf.Column(
		tpdf.Text("Hello").WithColor(tpdf.HexToRGBA("#ff0000")).WithFontSize(20),
	)

	generator := tpdf.NewGenerator()
	generator.SetMargins(8, 8, 8, 8)
	generator.SetMainWidget(content)

	generator.GenerateToFile("text_customizations.pdf")
}
