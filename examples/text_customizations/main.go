package main

import (
	tpdf "tpdf"
)

func main() {
	content := tpdf.Column(
		tpdf.Text("Hello", tpdf.FontSize(20), tpdf.FontColor("#ff0000")),
	)

	generator := tpdf.NewGenerator()
	generator.SetMargins(8, 8, 8, 8)
	generator.SetMainWidget(content)

	generator.GenerateToFile("text_customizations.pdf")
}
