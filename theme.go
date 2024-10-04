package tpdf

import "image/color"

type Theme struct {
	DefaultFontSize          float64
	DefaultFontColor         color.Color
	DefaultDividerColor      color.Color
	DefaultDividerCapStyle   DividerCapStyle
	DefaultDividerLineHeight float64
}

func NewTheme() *Theme {
	return &Theme{
		DefaultFontSize:          14,
		DefaultFontColor:         color.RGBA{0, 0, 0, 255},
		DefaultDividerColor:      color.RGBA{220, 220, 220, 255},
		DefaultDividerCapStyle:   DividerCapStyleButt,
		DefaultDividerLineHeight: 1,
	}
}
