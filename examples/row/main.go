package main

import "tpdf"

func main() {
	generator := tpdf.NewGenerator()

	content := tpdf.Column(
		tpdf.Row(
			tpdf.Text("Side "),
			tpdf.Text("by "),
			tpdf.Text("Side"),
			tpdf.Expand(
				tpdf.Row(
					tpdf.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
					tpdf.Text("B"),
				),
			),
		).WithMainAxisSize(tpdf.MainAxisSizeMin),
		tpdf.Text("Bottom Bottom Bottom"),
	)

	generator.SetMainWidget(content)
	generator.GenerateToFile("row.pdf")
}
