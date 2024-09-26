package tpdf

import (
	"image/color"
	"tpdf/internal/core"
)

type text struct {
	value     string
	fontSize  *float64
	color     *color.Color
	bold      bool
	italic    bool
	underline bool
}

func Text(value string) *text {
	return &text{
		value:     value,
		fontSize:  nil,
		color:     nil,
		bold:      false,
		italic:    false,
		underline: false,
	}
}

func (t *text) WithFontSize(fontSize float64) *text {
	t.fontSize = &fontSize
	return t
}

func (t *text) WithColor(color color.Color) *text {
	t.color = &color
	return t
}

func (t *text) InBold() *text {
	t.bold = true
	return t
}

func (t *text) InItalic() *text {
	t.italic = true
	return t
}

func (t *text) Underlined() *text {
	t.underline = true
	return t
}

func (t *text) calculatedFontSize(ctx *core.RenderContext) float64 {
	if t.fontSize != nil {
		return *t.fontSize
	}

	return ctx.DefaultFontSize()
}

func (t *text) calculatedFontColor(ctx *core.RenderContext) color.Color {
	if t.color != nil {
		return *t.color
	}

	return ctx.DefaultFontColor()
}

func (t *text) CalculateSize(ctx *core.RenderContext) (float64, float64) {
	return ctx.Writer.GetStringSize(
		t.value,
		t.calculatedFontSize(ctx),
		ctx.MaxWidth,
		t.calculatedFontColor(ctx),
		t.bold,
		t.italic,
		t.underline,
	)
}

func (t *text) Render(ctx *core.RenderContext) error {
	width, _ := t.CalculateSize(ctx)
	ctx.Writer.WriteMultiline(
		width,
		t.value,
		t.calculatedFontSize(ctx),
		t.calculatedFontColor(ctx),
		t.bold,
		t.italic,
		t.underline,
	)
	return nil
}
