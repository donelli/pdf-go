package tpdf

import "tpdf/internal/core"

type column struct {
	children []core.Widget
	spacing  float64
}

func Column(children ...core.Widget) *column {
	return &column{
		children: children,
		spacing:  0,
	}
}

func (t *column) CalculateSize(ctx *core.RenderContext) (float64, float64) {
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

func (t *column) Render(ctx *core.RenderContext) error {
	for _, child := range t.children {

		width, height := child.CalculateSize(ctx)
		ctx.Writer.WillWrite(width, height)

		x, y := ctx.Writer.X(), ctx.Writer.Y()

		err := child.Render(ctx)

		if err != nil {
			return err
		}

		ctx.Writer.SetX(x)
		ctx.Writer.SetY(y + height + t.spacing)
	}

	return nil
}

func (t *column) WithSpacing(spacing float64) *column {
	t.spacing = spacing
	return t
}
