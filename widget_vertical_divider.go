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
	color := ctx.Theme().DefaultDividerColor
	if d.color != nil {
		color = d.color
	}

	capStyle := ctx.Theme().DefaultDividerCapStyle
	if d.capStyle != DividerCapStyleButt {
		capStyle = d.capStyle
	}

	lineWidth := ctx.Theme().DefaultDividerLineHeight
	if d.lineWidth != 0 {
		lineWidth = d.lineWidth
	}

	ctx.Writer.Line(0, ctx.MaxHeight, color, lineWidth, capStyle)
	return nil
}

func (d *verticalDivider) CalculateSize(ctx *RenderContext) (width float64, height float64) {
	return d.lineWidth, ctx.MaxHeight
}
