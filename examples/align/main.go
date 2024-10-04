package main

import "tpdf"

func main() {
	content := tpdf.Column(
		buildContainer("center", tpdf.AlignmentCenter),
		buildContainer("bottom center", tpdf.AlignmentBottomCenter),
		buildContainer("top center", tpdf.AlignmentTopCenter),
		buildContainer("left center", tpdf.AlignmentLeftCenter),
		buildContainer("right center", tpdf.AlignmentRightCenter),
		buildContainer("top left", tpdf.AlignmentTopLeft),
		buildContainer("top right", tpdf.AlignmentTopRight),
		buildContainer("bottom left", tpdf.AlignmentBottomLeft),
		buildContainer("bottom right", tpdf.AlignmentBottomRight),
	).WithSpacing(6)

	theme := tpdf.NewTheme()
	writer := tpdf.NewWriter(8, 8, 8, 8, theme)

	writer.SetMainWidget(content)

	writer.GenerateToFile("align.pdf")
}

func buildContainer(label string, align tpdf.Alignment) tpdf.Widget {
	return tpdf.Container(
		tpdf.Align(
			align,
			tpdf.Text(label),
		),
	).WithSize(120, 120).Bordered(tpdf.HexToRGBA("#000000"), 1)
}
