package main

import (
	"pdf_go_test/core"
	pdf "pdf_go_test/widgets"
)

func main() {
	content := pdf.Column(
		pdf.Text("Hello", pdf.FontSize(20), pdf.FontColor("#ff0000")),
	)

	writer := core.NewWriter(8, 8, 8, 8)

	writer.RenderWidget(content)
	writer.GeneratePdfToFile("text_customizations.pdf")
}
