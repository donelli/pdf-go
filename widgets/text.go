package widgets

import (
	"image/color"
	"pdf_go_test/core"
)

type text struct {
	value    string
	fontSize float64
	color    color.RGBA
}

type fontSize struct {
	size float64
}

func FontSize(size float64) *fontSize {
	return &fontSize{
		size: size,
	}
}

func Text(configs ...any) *text {
	var value string
	var textFontSize float64 = 14
	color := color.RGBA{0, 0, 0, 255}

	for _, config := range configs {
		switch v := config.(type) {
		case string:
			value = v
		case *fontSize:
			textFontSize = v.size
		case *fontColor:
			color = v.RGBA
		}
	}

	return &text{
		value:    value,
		fontSize: textFontSize,
		color:    color,
	}
}

func (t *text) CalculateSize(ctx *core.RenderContext) (float64, float64) {
	return ctx.Writer.GetStringSize(t.value, t.fontSize, ctx.MaxWidth)

}

func (t *text) Render(ctx *core.RenderContext) error {
	width, _ := t.CalculateSize(ctx)
	ctx.Writer.WriteMultiline(width, t.value, t.fontSize, t.color)
	return nil
}
