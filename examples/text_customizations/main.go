package main

import (
	"strings"
	tpdf "tpdf"
)

func main() {
	content := tpdf.Column(
		tpdf.Text("Red with font size 20").WithColor(tpdf.HexToRGBA("#ff0000")).WithFontSize(20),
		tpdf.Text("Bold text").InBold(),
		tpdf.Text("In italic").InItalic(),
		tpdf.Text("Underlined").Underlined(),
		tpdf.Text("Strike out").StrikeOut(),
		tpdf.Text(strings.Repeat("LeftAlign. ", 15)).Align(tpdf.TextAlignLeft),
		tpdf.Text(strings.Repeat("RightAlign. ", 15)).Align(tpdf.TextAlignRight),
		tpdf.Text(strings.Repeat("CenterAlign. ", 15)).Align(tpdf.TextAlignCenter),
	).WithSpacing(8)

	generator := tpdf.NewGenerator()
	generator.SetMargins(8, 8, 8, 8)
	generator.SetMainWidget(content)

	generator.GenerateToFile("text_customizations.pdf")
}
