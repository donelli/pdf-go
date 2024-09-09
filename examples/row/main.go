package main

import (
	"pdf_go_test/core"
	pdf "pdf_go_test/widgets"
)

func main() {
	writer := core.NewWriter(0, 0, 0, 0)

	content := pdf.Column(
		pdf.Row(
			pdf.RowConfig{
				MainAxisSize: pdf.MainAxisSizeMin,
			},
			pdf.Text("Side "),
			pdf.Text("by "),
			pdf.Text("Side"),
			pdf.Expand(
				pdf.Row(
					pdf.RowConfig{
						MainAxisSize: pdf.MainAxisSizeMax,
					},
					pdf.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
					pdf.Text("B"),
				),
			),
		),
		pdf.Text("Bottom Bottom Bottom"),
	)

	writer.RenderWidget(content)
	writer.GeneratePdf("row.pdf")
}
