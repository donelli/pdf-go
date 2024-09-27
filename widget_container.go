package tpdf

import (
	"image/color"
	"tpdf/internal/core"
)

type container struct {
	width           *float64
	height          *float64
	backgroundColor *color.Color
}

func Container() *container {
	return &container{}
}

func (c *container) WithWidth(width float64) *container {
	c.width = &width
	return c
}

func (c *container) WithHeight(height float64) *container {
	c.height = &height
	return c
}

func (c *container) WithSize(width, height float64) *container {
	c.width = &width
	c.height = &height
	return c
}

func (c *container) WithBackgroundColor(backgroundColor color.Color) *container {
	c.backgroundColor = &backgroundColor
	return c
}

func (c *container) CalculateSize(ctx *core.RenderContext) (float64, float64) {
	width := 0.0
	if c.width != nil {
		width = *c.width
	}

	height := 0.0
	if c.height != nil {
		height = *c.height
	}

	return float64(width), float64(height)
}

func (c *container) Render(ctx *core.RenderContext) error {

	width, height := c.CalculateSize(ctx)

	if width == 0 && height == 0 {
		return nil
	}

	if c.backgroundColor != nil {
		ctx.Writer.Rect(width, height, *c.backgroundColor)
	}

	return nil
}
