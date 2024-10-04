package tpdf

func Padding(padding PaddingValue, child Widget) *container {
	return Container(child).
		Padding(
			padding.Left(), padding.Right(),
			padding.Top(), padding.Bottom(),
		)
}
