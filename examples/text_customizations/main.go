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
		tpdf.Text("With link").Link("https://google.com"),
		tpdf.Text(strings.Repeat("Multi line link. ", 10)).Link("https://google.com"),
		tpdf.Text("ABC ABC ABC ABC ABC ABC ABC ABC ABC ABC ABC ABC ABC ABC ABC ABC AA B C D E F G H").MaxLines(1),
	).WithSpacing(10)

	theme := tpdf.NewTheme()
	writer := tpdf.NewWriter(8, 8, 8, 8, theme)

	writer.SetMainWidget(content)

	writer.GenerateToFile("text_customizations.pdf")
}
