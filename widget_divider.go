package tpdf

import "image/color"

type divider struct {
	lineHeight float64
	color      color.Color
	capStyle   DividerCapStyle
}

func Divider() *divider {
	return &divider{}
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
	color := ctx.Theme().DefaultDividerColor
	if d.color != nil {
		color = d.color
	}

	capStyle := ctx.Theme().DefaultDividerCapStyle
	if d.capStyle != DividerCapStyleButt {
		capStyle = d.capStyle
	}

	lineHeight := ctx.Theme().DefaultDividerLineHeight
	if d.lineHeight != 0 {
		lineHeight = d.lineHeight
	}

	ctx.Writer.Line(ctx.MaxWidth, 0, color, lineHeight, capStyle)
	return nil
}

func (d *divider) CalculateSize(ctx *RenderContext) (width float64, height float64) {
	return ctx.MaxWidth, d.lineHeight
}
