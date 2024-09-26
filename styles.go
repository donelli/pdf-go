package tpdf

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
