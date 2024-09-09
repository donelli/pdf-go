package core

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
