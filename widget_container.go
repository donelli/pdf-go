package tpdf

import (
	"image/color"
	"math"
)

type container struct {
	width           *float64
	height          *float64
	backgroundColor color.Color
	child           Widget
	paddingLeft     float64
	paddingRight    float64
	paddingTop      float64
	paddingBottom   float64
	borderRadius    borderRadius
	borderColor     color.Color
	borderWidth     float64
}

func Container(child Widget) *container {
	return &container{
		child:       child,
		borderWidth: 1,
	}
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
	c.backgroundColor = backgroundColor
	return c
}

func (c *container) PaddingAll(padding float64) *container {
	c.paddingLeft = padding
	c.paddingRight = padding
	c.paddingTop = padding
	c.paddingBottom = padding
	return c
}

func (c *container) Padding(paddingLeft, paddingRight, paddingTop, paddingBottom float64) *container {
	c.paddingLeft = paddingLeft
	c.paddingRight = paddingRight
	c.paddingTop = paddingTop
	c.paddingBottom = paddingBottom
	return c
}

func (c *container) PaddingHorizontal(padding float64) *container {
	c.paddingLeft = padding
	c.paddingRight = padding
	return c
}

func (c *container) PaddingVertical(padding float64) *container {
	c.paddingTop = padding
	c.paddingBottom = padding
	return c
}

func (c *container) BorderRadius(borderRadius borderRadius) *container {
	c.borderRadius = borderRadius
	return c
}

func (c *container) Bordered(borderColor color.Color, borderWidth float64) *container {
	c.borderColor = borderColor
	c.borderWidth = borderWidth
	return c
}

func (c *container) CalculateSize(ctx *RenderContext) (float64, float64) {
	width := 0.0
	if c.width != nil {
		width = *c.width
	}

	height := 0.0
	if c.height != nil {
		height = *c.height
	}

	if c.child != nil && (width == 0 || height == 0) {

		updatedCtx := ctx.Copy()

		if width != 0 {
			updatedCtx.MaxWidth = width
		}

		if height != 0 {
			updatedCtx.MaxHeight = height
		}

		childWidth, childHeight := c.child.CalculateSize(updatedCtx)

		if width == 0 {
			width = childWidth
		}

		if height == 0 {
			height = childHeight
		}
	}

	if width == 0 || width == math.MaxFloat64 {
		width = ctx.MaxWidth
	} else {
		width += c.paddingLeft + c.paddingRight
	}

	if height == 0 || height == math.MaxFloat64 {
		height = ctx.MaxHeight
	} else {
		height += c.paddingTop + c.paddingBottom
	}

	return float64(width), float64(height)
}

func (c *container) Render(ctx *RenderContext) error {

	width, height := c.CalculateSize(ctx)

	if c.borderRadius != nil {
		ctx.Writer.RoundedRect(
			width, height,
			c.backgroundColor,
			c.borderRadius, c.borderColor, c.borderWidth,
		)
	} else {
		ctx.Writer.Rect(
			width, height,
			c.backgroundColor,
			c.borderColor, c.borderWidth,
		)
	}

	if c.child != nil {
		ctx.Writer.AddOffsets(c.paddingLeft, c.paddingTop)

		updatedCtx := ctx.Copy()
		updatedCtx.MaxWidth = width - c.paddingLeft - c.paddingRight
		updatedCtx.MaxHeight = height - c.paddingTop - c.paddingBottom

		c.child.Render(updatedCtx)

		ctx.Writer.SubtractOffsets(c.paddingLeft, c.paddingTop)
	}

	return nil
}
