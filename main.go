package main

import (
	"fmt"
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

	writer := core.NewWriter(8, 8, 8, 8)

	writer.SetFooter(func(page int, totalPagesAlias string) core.Widget {
		return pdf.Column(
			pdf.Text(fmt.Sprintf("Page %d of %s", page, totalPagesAlias)),
			pdf.Text(fmt.Sprintf("Page %d of %s", page, totalPagesAlias)),
		)
	})

	writer.RenderWidget(content)
	writer.GeneratePdf("test.pdf")
}
