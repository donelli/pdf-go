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

type borderRadius interface {
	TopLeft() float64
	TopRight() float64
	BottomLeft() float64
	BottomRight() float64
}

type borderRadiusAll float64

func BorderRadiusAll(radius float64) *borderRadiusAll {
	return (*borderRadiusAll)(&radius)
}

func (b borderRadiusAll) TopLeft() float64 {
	return float64(b)
}

func (b borderRadiusAll) TopRight() float64 {
	return float64(b)
}

func (b borderRadiusAll) BottomLeft() float64 {
	return float64(b)
}

func (b borderRadiusAll) BottomRight() float64 {
	return float64(b)
}

type borderRadiusEach struct {
	topLeft     float64
	topRight    float64
	bottomLeft  float64
	bottomRight float64
}

func BorderRadiusEach(topLeft, topRight, bottomLeft, bottomRight float64) *borderRadiusEach {
	return &borderRadiusEach{
		topLeft:     topLeft,
		topRight:    topRight,
		bottomLeft:  bottomLeft,
		bottomRight: bottomRight,
	}
}

func (b *borderRadiusEach) TopLeft() float64 {
	return b.topLeft
}

func (b *borderRadiusEach) TopRight() float64 {
	return b.topRight
}

func (b *borderRadiusEach) BottomLeft() float64 {
	return b.bottomLeft
}

func (b *borderRadiusEach) BottomRight() float64 {
	return b.bottomRight
}
