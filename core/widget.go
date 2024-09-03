package core

type Widget interface {
	Render(ctx *RenderContext) error
	CalculateSize(ctx *RenderContext) (float64, float64)
}
