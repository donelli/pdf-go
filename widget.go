package tpdf

type Widget interface {
	Render(ctx *RenderContext) error
	CalculateSize(ctx *RenderContext) (width float64, height float64)
}
