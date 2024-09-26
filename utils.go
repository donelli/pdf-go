package tpdf

import (
	"image/color"
	"strconv"
)

func HexToRGBA(hex string) color.RGBA {
	values, err := strconv.ParseUint(string(hex[1:]), 16, 32)

	if err != nil {
		panic(err)
	}

	return color.RGBA{
		R: uint8(values >> 16),
		G: uint8((values >> 8) & 0xFF),
		B: uint8(values & 0xFF),
		A: 255,
	}
}
