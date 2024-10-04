package main

import (
	tpdf "tpdf"
)

func main() {
	content := tpdf.Column(
		tpdf.Text("Two containers with fixed width and height and no child"),
		tpdf.Row(
			tpdf.Container(nil).
				WithSize(100, 100).
				WithBackgroundColor(tpdf.HexToRGBA("#ff0000")),

			tpdf.Container(nil).
				WithSize(50, 50).
				WithBackgroundColor(tpdf.HexToRGBA("#0000ff")),
		).
			WithMainAxisSize(tpdf.MainAxisSizeMin).
			WithSpacing(8),

		tpdf.Divider(),

		tpdf.Text("Container with fixed width and height and no child"),
		tpdf.Container(nil).
			WithSize(100, 20).
			WithBackgroundColor(tpdf.HexToRGBA("#00ff00")),

		tpdf.Divider(),

		tpdf.Text("Container with padding, child and no fixed size"),
		tpdf.Container(
			tpdf.Text("Child"),
		).
			WithBackgroundColor(tpdf.HexToRGBA("#bbffbb")).
			PaddingAll(8),

		tpdf.Divider(),

		tpdf.Text("Container with padding, child, fixed width and no height"),
		tpdf.Container(
			tpdf.Text("Child"),
		).
			WithWidth(100).
			WithBackgroundColor(tpdf.HexToRGBA("#dddddd")).
			PaddingAll(8),

		tpdf.Divider(),

		tpdf.Text("Container with fixed size (40x40) and border radius all 4"),
		tpdf.Container(
			tpdf.Text("A"),
		).
			WithSize(40, 40).
			WithBackgroundColor(tpdf.HexToRGBA("#dddddd")).
			BorderRadius(tpdf.BorderRadiusAll(4)),

		tpdf.Divider(),

		tpdf.Text("Container with fixed size (40x40) and border radius all 4"),
		tpdf.Container(
			tpdf.Text("Border radius in specific corders"),
		).
			WithBackgroundColor(tpdf.HexToRGBA("#ff00ff")).
			BorderRadius(tpdf.BorderRadiusEach(4, 0, 0, 4)).
			PaddingAll(2),

		tpdf.Divider(),

		tpdf.Text("Container with no fixed size, no border radius and green border with 1 px size"),
		tpdf.Container(
			tpdf.Text("Child"),
		).
			PaddingAll(4).
			Bordered(tpdf.HexToRGBA("#00ff00"), 1),

		tpdf.Divider(),

		tpdf.Text("Container with no fixed size, with border radius and green border with 2 px size"),
		tpdf.Container(
			tpdf.Text("Child"),
		).
			PaddingAll(4).
			BorderRadius(tpdf.BorderRadiusAll(4)).
			Bordered(tpdf.HexToRGBA("#00ff00"), 2),
	).WithSpacing(8)

	theme := tpdf.NewTheme()
	writer := tpdf.NewWriter(8, 8, 8, 8, theme)

	writer.SetMainWidget(content)

	writer.GenerateToFile("container.pdf")
}
