package widgets

import (
	"image/color"
	"pdf_go_test/core"
)

type fontColor struct {
	color.RGBA
}

func FontColor(hexColor string) *fontColor {
	rgba := core.HexToRGBA(hexColor)

	return &fontColor{
		rgba,
	}
}
