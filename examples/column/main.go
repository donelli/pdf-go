package main

import (
	"pdf_go_test/core"
	pdf "pdf_go_test/widgets"
	"strings"
)

func main() {
	content := pdf.Column(
		pdf.Text(strings.Repeat("Hello ", 20), pdf.FontSize(14)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
		pdf.Text("Hello", pdf.FontSize(20)),
	)

	writer := core.NewWriter(0, 0, 0, 0)

	writer.RenderWidget(content)
	writer.GeneratePdfToFile("column.pdf")
}
