package core

import "image/color"

type RenderContext struct {
	Writer    *Writer
	MaxWidth  float64
	MaxHeight float64
}

func (b *RenderContext) Copy() *RenderContext {
	return &RenderContext{
		Writer:    b.Writer,
		MaxWidth:  b.MaxWidth,
		MaxHeight: b.MaxHeight,
	}
}

func (b *RenderContext) HorizontalMargin() float64 {
	return b.Writer.marginLeft + b.Writer.marginRight
}

func (b *RenderContext) VerticalMargin() float64 {
	return b.Writer.marginTop + b.Writer.marginBottom
}

func (b *RenderContext) DefaultFontSize() float64 {
	return b.Writer.defaultFontSize
}

func (b *RenderContext) DefaultFontColor() color.RGBA {
	return b.Writer.defaultFontColor
}
