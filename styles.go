package tpdf

type MainAxisSize int8

const (
	MainAxisSizeMin MainAxisSize = iota
	MainAxisSizeMax
)

type TextAlign int8

const (
	TextAlignAuto TextAlign = iota
	TextAlignLeft
	TextAlignCenter
	TextAlignRight
)
