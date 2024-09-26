package core

type RenderContext struct {
	Writer       *Writer
	MaxWidth     float64
	MaxHeight    float64
	MarginLeft   float64
	MarginRight  float64
	MarginTop    float64
	MarginBottom float64
}

func (b *RenderContext) Copy() *RenderContext {
	return &RenderContext{
		Writer:       b.Writer,
		MaxWidth:     b.MaxWidth,
		MaxHeight:    b.MaxHeight,
		MarginLeft:   b.MarginLeft,
		MarginRight:  b.MarginRight,
		MarginTop:    b.MarginTop,
		MarginBottom: b.MarginBottom,
	}
}

func (b *RenderContext) HorizontalMargin() float64 {
	return b.MarginLeft + b.MarginRight
}

func (b *RenderContext) VerticalMargin() float64 {
	return b.MarginTop + b.MarginBottom
}
