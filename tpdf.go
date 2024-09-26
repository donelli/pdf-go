package tpdf

import "tpdf/internal/core"

type Generator struct {
	writer                                           *core.Writer
	topMargin, rightMargin, bottomMargin, leftMargin float64
	mainWidget                                       core.Widget
}

func NewGenerator() *Generator {
	return &Generator{
		topMargin:    8,
		rightMargin:  8,
		bottomMargin: 8,
		leftMargin:   8,
	}
}

func (g *Generator) SetMargins(top, right, bottom, left float64) {
	g.topMargin = top
	g.rightMargin = right
	g.bottomMargin = bottom
	g.leftMargin = left
}

func (g *Generator) SetMainWidget(widget core.Widget) {
	g.mainWidget = widget
}

func (g *Generator) GenerateToFile(filename string) error {
	g.writer = core.NewWriter(g.topMargin, g.rightMargin, g.bottomMargin, g.leftMargin)
	g.writer.RenderWidget(g.mainWidget)
	return g.writer.GeneratePdfToFile(filename)
}
