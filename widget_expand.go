package tpdf

type expand struct {
	child Widget
}

func Expand(child Widget) *expand {
	return &expand{
		child: child,
	}
}

func (e *expand) CalculateSize(ctx *RenderContext) (float64, float64) {
	return e.child.CalculateSize(ctx)
}

func (e *expand) Render(ctx *RenderContext) error {
	return e.child.Render(ctx)
}
