package tpdf

import (
	"image/color"
	"tpdf/internal/core"
)

type Generator struct {
	writer                                           *core.Writer
	topMargin, rightMargin, bottomMargin, leftMargin float64
	mainWidget                                       core.Widget
	defaultFontSize                                  *float64
	defaultFontColor                                 *color.Color
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

func (g *Generator) GenerateToFile(filename string) error {
	g.writer = core.NewWriter(g.topMargin, g.rightMargin, g.bottomMargin, g.leftMargin)

	if g.defaultFontSize != nil {
		g.writer.SetDefaultFontSize(*g.defaultFontSize)
	}

	if g.defaultFontColor != nil {
		g.writer.SetDefaultFontColor(*g.defaultFontColor)
	}

	g.writer.RenderWidget(g.mainWidget)

	return g.writer.GeneratePdfToFile(filename)
}
