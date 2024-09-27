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
	strikeOut bool
	textAlign TextAlign
	link      string
	maxLines  int
}

func Text(value string) *text {
	return &text{
		value:     value,
		fontSize:  nil,
		color:     nil,
		bold:      false,
		italic:    false,
		underline: false,
		strikeOut: false,
		textAlign: TextAlignAuto,
		maxLines:  0,
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

func (t *text) StrikeOut() *text {
	t.strikeOut = true
	return t
}

func (t *text) Align(textAlign TextAlign) *text {
	t.textAlign = textAlign
	return t
}

func (t *text) Link(link string) *text {
	t.link = link
	return t
}

func (t *text) MaxLines(maxLines int) *text {
	t.maxLines = maxLines
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
		t.strikeOut,
		t.maxLines,
	)
}

func (t *text) Render(ctx *core.RenderContext) error {
	ctx.Writer.WriteMultiline(
		ctx.MaxWidth,
		t.value,
		t.calculatedFontSize(ctx),
		t.calculatedFontColor(ctx),
		t.bold,
		t.italic,
		t.underline,
		t.strikeOut,
		core.TextAlign(t.textAlign),
		t.link,
		t.maxLines,
	)
	return nil
}
