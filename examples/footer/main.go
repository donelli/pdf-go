package main

import "tpdf"

func main() {
	content := tpdf.Column(
		tpdf.Text("A").WithFontSize(60),
		tpdf.Text("B").WithFontSize(60),
		tpdf.Text("C").WithFontSize(60),
		tpdf.Text("D").WithFontSize(60),
		tpdf.Text("E").WithFontSize(60),
		tpdf.Text("F").WithFontSize(60),
		tpdf.Text("G").WithFontSize(60),
		tpdf.Text("H").WithFontSize(60),
		tpdf.Text("I").WithFontSize(60),
		tpdf.Text("J").WithFontSize(60),
		tpdf.Text("K").WithFontSize(60),
	).WithSpacing(32)

	theme := tpdf.NewTheme()
	writer := tpdf.NewWriter(8, 8, 8, 8, theme)

	writer.SetMainWidget(content)

	writer.SetFooter(func(page int, totalPagesAlias string) tpdf.Widget {

		if page == 1 {
			return tpdf.Text("Footer")
		}

		return tpdf.Column(
			tpdf.Text("Footer - first line"),
			tpdf.Text("Footer - second line"),
		)
	})

	writer.GenerateToFile("footer.pdf")
}
