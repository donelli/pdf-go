package tpdf

type column struct {
	children []Widget
	spacing  float64
}

func Column(children ...Widget) *column {
	return &column{
		children: children,
		spacing:  0,
	}
}

func (t *column) CalculateSize(ctx *RenderContext) (float64, float64) {
	width := 0.0
	height := 0.0

	for _, child := range t.children {
		childWidth, childHeight := child.CalculateSize(ctx)

		if childWidth > width {
			width = childWidth
		}

		height += childHeight
	}

	height += float64(len(t.children)-1) * t.spacing

	return width, height
}

func (t *column) Render(ctx *RenderContext) error {
	for _, child := range t.children {

		width, height := child.CalculateSize(ctx)
		ctx.Writer.WillWrite(width, height)

		x, y := ctx.Writer.X(), ctx.Writer.Y()
		initialPageNumber := ctx.Writer.PageNumber()

		err := child.Render(ctx)

		if err != nil {
			return err
		}

		if ctx.Writer.PageNumber() == initialPageNumber {
			ctx.Writer.SetY(y + height + t.spacing)
		}

		ctx.Writer.SetX(x)
	}

	return nil
}

func (t *column) WithSpacing(spacing float64) *column {
	t.spacing = spacing
	return t
}
