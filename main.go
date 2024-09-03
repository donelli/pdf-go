package main

import (
	"pdf_go_test/core"
	pdf "pdf_go_test/widgets"
	"strings"
)

func main() {
	content := pdf.Column(
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
		pdf.Text(strings.Repeat("Hello ", 1), pdf.FontSize(20)),
	)

	writer := core.NewWriter()
	writer.RenderWidget(content)
	writer.GeneratePdf("test.pdf")
}
