package tpdf

import (
	"image/color"
	"tpdf/internal/core"
)

type Widget core.Widget

type Generator struct {
	writer                                           *core.Writer
	topMargin, rightMargin, bottomMargin, leftMargin float64
	mainWidget                                       core.Widget
	defaultFontSize                                  *float64
	defaultFontColor                                 *color.Color
	footerHandler                                    func(page int, totalPagesAlias string) Widget
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

func (g *Generator) SetDefaultFontSize(size float64) {
	g.defaultFontSize = &size
}

func (g *Generator) SetDefaultFontColor(color color.Color) {
	g.defaultFontColor = &color
}

func (g *Generator) SetFooter(handler func(page int, totalPagesAlias string) Widget) {
	g.footerHandler = handler
}

func (g *Generator) GenerateToFile(filename string) error {
	g.writer = core.NewWriter(g.topMargin, g.rightMargin, g.bottomMargin, g.leftMargin)

	if g.defaultFontSize != nil {
		g.writer.SetDefaultFontSize(*g.defaultFontSize)
	}

	if g.defaultFontColor != nil {
		g.writer.SetDefaultFontColor(*g.defaultFontColor)
	}

	if g.footerHandler != nil {
		g.writer.SetFooter(func(page int, totalPagesAlias string) core.Widget {
			return g.footerHandler(page, totalPagesAlias)
		})
	}

	g.writer.RenderWidget(g.mainWidget)

	return g.writer.GeneratePdfToFile(filename)
}
