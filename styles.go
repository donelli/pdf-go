package tpdf

import (
	"image/color"
	"tpdf/internal/core"
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

type MainAxisAlignment int16

const (
	MainAxisAlignmentStart MainAxisAlignment = iota
	MainAxisAlignmentEnd
	MainAxisAlignmentCenter
	MainAxisAlignmentSpaceBetween
	MainAxisAlignmentSpaceAround
	MainAxisAlignmentSpaceEvenly
)
