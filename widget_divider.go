package tpdf

import "image/color"

type divider struct {
	lineHeight float64
	color      color.Color
	capStyle   DividerCapStyle
}

func Divider() *divider {
	return &divider{
		lineHeight: 1,
	}
}

func (d *divider) LineHeight(lineHeight float64) *divider {
	d.lineHeight = lineHeight
	return d
}

func (d *divider) Color(color color.Color) *divider {
	d.color = color
	return d
}

func (d *divider) CapStyle(capStyle DividerCapStyle) *divider {
	d.capStyle = capStyle
	return d
}

func (d *divider) Render(ctx *RenderContext) error {
	ctx.Writer.Line(ctx.MaxWidth, 0, d.color, d.lineHeight, d.capStyle)
	return nil
}

func (d *divider) CalculateSize(ctx *RenderContext) (width float64, height float64) {
	return ctx.MaxWidth, d.lineHeight
}
