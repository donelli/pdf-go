package tpdf

import (
	"fmt"
	"tpdf/internal/core"
)

type row struct {
	children     []core.Widget
	mainAxisSize MainAxisSize
}

func Row(
	children ...core.Widget,
) *row {
	return &row{
		children:     children,
		mainAxisSize: MainAxisSizeMax,
	}
}

func (r *row) getWidthPerChild(ctx *core.RenderContext) []float64 {
	widthPerChild := make([]float64, len(r.children))

	if r.mainAxisSize == MainAxisSizeMax {
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

func sum(numbers []float64) float64 {
	var result float64 = 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func (r *row) Render(ctx *core.RenderContext) error {

	width, height := r.CalculateSize(ctx)
	ctx.Writer.WillWrite(width, height)

	widthPerChild := r.getWidthPerChild(ctx)

	fmt.Println(widthPerChild)
	fmt.Println("sum" + fmt.Sprint(sum(widthPerChild)))
	fmt.Println("max" + fmt.Sprint(ctx.MaxWidth))

	for index, child := range r.children {
		childWidth := widthPerChild[index]

		x := ctx.Writer.X()

		fmt.Println("current x: " + fmt.Sprint(x))

		if r.mainAxisSize == MainAxisSizeMin {

			maxX := ctx.MaxWidth + ctx.HorizontalMargin()
			nextX := x + childWidth

			if nextX > maxX {
				panic(fmt.Sprint("overflow by ", nextX-maxX))
			}
		}

		context := ctx.Copy()
		context.MaxWidth = childWidth

		if err := child.Render(context); err != nil {
			return err
		}

		ctx.Writer.SetX(x + childWidth)
	}

	return nil
}

func (r *row) WithMainAxisSize(mainAxisSize MainAxisSize) *row {
	r.mainAxisSize = mainAxisSize
	return r
}
