package tpdf

type align struct {
	alignment Alignment
	child     Widget
}

type Alignment uint8

const (
	AlignmentCenter Alignment = iota
	AlignmentTopCenter
	AlignmentBottomCenter
	AlignmentLeftCenter
	AlignmentRightCenter
	AlignmentTopLeft
	AlignmentTopRight
	AlignmentBottomLeft
	AlignmentBottomRight
)

func Align(
	alignment Alignment,
	child Widget,
) *align {
	return &align{
		child:     child,
		alignment: alignment,
	}
}

func (c *align) CalculateSize(ctx *RenderContext) (float64, float64) {
	return ctx.MaxWidth, ctx.MaxHeight
}

func (c *align) Render(ctx *RenderContext) error {
	childWidth, childHeight := c.child.CalculateSize(ctx)

	xOffset := 0.0
	yOffset := 0.0

	if c.alignment == AlignmentTopCenter || c.alignment == AlignmentTopLeft || c.alignment == AlignmentTopRight {
		yOffset = 0
	} else if c.alignment == AlignmentBottomCenter || c.alignment == AlignmentBottomLeft || c.alignment == AlignmentBottomRight {
		yOffset = ctx.MaxHeight - childHeight
	} else {
		yOffset = (ctx.MaxHeight - childHeight) / 2
	}

	if c.alignment == AlignmentLeftCenter || c.alignment == AlignmentTopLeft || c.alignment == AlignmentBottomLeft {
	} else if c.alignment == AlignmentRightCenter || c.alignment == AlignmentTopRight || c.alignment == AlignmentBottomRight {
		xOffset = ctx.MaxWidth - childWidth
	} else {
		xOffset = (ctx.MaxWidth - childWidth) / 2
	}

	ctx.Writer.AddOffsets(xOffset, yOffset)

	err := c.child.Render(ctx)

	ctx.Writer.SubtractOffsets(xOffset, yOffset)

	return err
}
