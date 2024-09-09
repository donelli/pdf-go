package widgets

import "pdf_go_test/core"

type expand struct {
	child core.Widget
}

func Expand(child core.Widget) *expand {
	return &expand{
		child: child,
	}
}

func (e *expand) CalculateSize(ctx *core.RenderContext) (float64, float64) {
	return e.child.CalculateSize(ctx)
}

func (e *expand) Render(ctx *core.RenderContext) error {
	return e.child.Render(ctx)
}
