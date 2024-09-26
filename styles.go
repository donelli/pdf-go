package tpdf

import "tpdf/internal/core"

type MainAxisAlignment int16

const (
	MainAxisAlignmentStart MainAxisAlignment = iota
	MainAxisAlignmentEnd
	MainAxisAlignmentCenter
	MainAxisAlignmentSpaceBetween
	MainAxisAlignmentSpaceAround
	MainAxisAlignmentSpaceEvenly
)

type MainAxisSize int8

const (
	MainAxisSizeMin MainAxisSize = iota
	MainAxisSizeMax
)

type TextAlign core.TextAlign

const (
	TextAlignLeft   TextAlign = TextAlign(core.TextAlignLeft)
	TextAlignCenter TextAlign = TextAlign(core.TextAlignCenter)
	TextAlignRight  TextAlign = TextAlign(core.TextAlignRight)
	TextAlignAuto   TextAlign = TextAlign(core.TextAlignAuto)
)
