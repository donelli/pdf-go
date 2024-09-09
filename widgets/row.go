package widgets

import (
	"fmt"
	"pdf_go_test/core"
)

type MainAxisSize int8

const (
	MainAxisSizeMin MainAxisSize = iota
	MainAxisSizeMax
)

type RowConfig struct {
	MainAxisSize MainAxisSize
}

type row struct {
	children []core.Widget
	config   RowConfig
}

func Row(
	config RowConfig,
	children ...core.Widget,
) *row {
	return &row{
		children: children,
		config:   config,
	}
}

func (r *row) getWidthPerChild(ctx *core.RenderContext) []float64 {
	widthPerChild := make([]float64, len(r.children))

	if r.config.MainAxisSize == MainAxisSizeMax {
		for i := range r.children {
			widthPerChild[i] = ctx.MaxWidth / float64(len(r.children))
		}

		return widthPerChild
	}

	originalWidths := make([]float64, len(r.children))
	availableWidth := ctx.MaxWidth
	numberOfExpands := 0.0

	for i, child := range r.children {
		switch child.(type) {
		case *expand:
			originalWidths[i] = 0
			numberOfExpands += 1
		default:
			originalWidths[i], _ = child.CalculateSize(ctx)
			availableWidth -= originalWidths[i]
		}
	}

	for i, originalWidth := range originalWidths {
		switch r.children[i].(type) {
		case *expand:
			widthPerChild[i] = availableWidth / numberOfExpands
		default:
			widthPerChild[i] = originalWidth
		}
	}

	return widthPerChild
}

func (r *row) CalculateSize(ctx *core.RenderContext) (float64, float64) {
	width := 0.0
	height := 0.0

	widthPerChild := r.getWidthPerChild(ctx)

	for i, child := range r.children {
		context := ctx.Copy()
		context.MaxWidth = widthPerChild[i]

		_, childHeight := child.CalculateSize(context)

		if childHeight > height {
			height = childHeight
		}
	}

	for _, childWidth := range widthPerChild {
		width += childWidth
	}

	return width, height
}

func (r *row) Render(ctx *core.RenderContext) error {

	width, height := r.CalculateSize(ctx)
	ctx.Writer.WillWrite(width, height)

	x := ctx.Writer.X()

	widthPerChild := r.getWidthPerChild(ctx)

	for index, child := range r.children {
		childWidth := widthPerChild[index]

		if r.config.MainAxisSize == MainAxisSizeMin {
			if x+childWidth > ctx.MaxWidth {
				panic(fmt.Sprint("overflow by ", x+childWidth-ctx.MaxWidth, " childWidth ", childWidth, " maxWidth ", ctx.MaxWidth))
			}
		}

		context := ctx.Copy()
		context.MaxWidth = childWidth

		if err := child.Render(context); err != nil {
			return err
		}

		x += childWidth
		ctx.Writer.SetX(x)
	}

	return nil
}
