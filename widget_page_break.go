package tpdf

type pageBreak struct{}

func PageBreak() *pageBreak {
	return &pageBreak{}
}

func (p *pageBreak) Render(ctx *RenderContext) error {
	ctx.Writer.AddPage()
	return nil
}

func (p *pageBreak) CalculateSize(ctx *RenderContext) (width float64, height float64) {
	return 0, 0
}
