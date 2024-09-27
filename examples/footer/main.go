package main

import "tpdf"

func main() {
	generator := tpdf.NewGenerator()

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

	generator.SetFooter(func(page int, totalPagesAlias string) tpdf.Widget {

		if page == 1 {
			return tpdf.Text("Footer")
		}

		return tpdf.Column(
			tpdf.Text("Footer - first line"),
			tpdf.Text("Footer - second line"),
		)
	})

	generator.SetMainWidget(content)
	generator.GenerateToFile("footer.pdf")
}