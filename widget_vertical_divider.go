package tpdf

import "image/color"

type verticalDivider struct {
	lineWidth float64
	color     color.Color
	capStyle  DividerCapStyle
}

func VerticalDivider() *verticalDivider {
	return &verticalDivider{
		lineWidth: 1,
	}
}

func (d *verticalDivider) LineWidth(lineWidth float64) *verticalDivider {
	d.lineWidth = lineWidth
	return d
}

func (d *verticalDivider) Color(color color.Color) *verticalDivider {
	d.color = color
	return d
}

func (d *verticalDivider) CapStyle(capStyle DividerCapStyle) *verticalDivider {
	d.capStyle = capStyle
	return d
}

func (d *verticalDivider) Render(ctx *RenderContext) error {
	ctx.Writer.Line(0, ctx.MaxHeight, d.color, d.lineWidth, d.capStyle)
	return nil
}

func (d *verticalDivider) CalculateSize(ctx *RenderContext) (width float64, height float64) {
	return d.lineWidth, ctx.MaxHeight
}
