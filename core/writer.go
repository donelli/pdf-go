package core

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/go-pdf/fpdf"
)

const debug = true

type Writer struct {
	Pdf             *fpdf.Fpdf
	x               float64
	y               float64
	footerHeight    float64
	marginLeft      float64
	marginRight     float64
	marginTop       float64
	marginBottom    float64
	ignorePageBreak bool
}

func NewWriter(topMargin, rightMargin, bottomMargin, leftMargin float64) *Writer {
	pdf := fpdf.New("P", "pt", "A4", "")
	pdf.SetFont("Arial", "", 14)
	pdf.SetCellMargin(0)
	pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)

	pdf.AddPage()

	w := &Writer{
		Pdf:             pdf,
		x:               leftMargin,
		y:               topMargin,
		marginLeft:      leftMargin,
		marginRight:     rightMargin,
		marginTop:       topMargin,
		marginBottom:    bottomMargin,
		ignorePageBreak: false,
	}

	pdf.AliasNbPages(w.getNbAlias())

	return w
}

func (w *Writer) MaxWidth() float64 {
	pageWidth, _ := w.Pdf.GetPageSize()
	return pageWidth - w.marginLeft - w.marginRight
}

func (w *Writer) MaxHeight() float64 {
	_, pageHeight := w.Pdf.GetPageSize()

	return pageHeight - w.footerHeight - w.marginTop - w.marginBottom
}

func (w *Writer) NewBuildContext() *RenderContext {
	return &RenderContext{
		Writer:    w,
		MaxWidth:  w.MaxWidth(),
		MaxHeight: w.MaxHeight(),
	}
}

func (w *Writer) SetX(x float64) {
	if debug {
		fmt.Println("[DEBUG] SetX:", x)
	}

	w.x = x
}

func (w *Writer) SetY(y float64) {
	if debug {
		fmt.Println("[DEBUG] SetY:", y)
	}

	w.y = y

	if w.y >= w.MaxHeight() {
		w.BreakPage()
	}
}

func (w *Writer) SetYWithoutPageBreakCheck(y float64) {
	if debug {
		fmt.Println("[DEBUG] SetYWithoutPageBreakCheck:", y)
	}

	w.y = y
}

func (w *Writer) BreakPage() {
	if w.ignorePageBreak {
		return
	}

	if debug {
		fmt.Println("[DEBUG] BreakPage")
	}

	w.Pdf.AddPage()
	w.y = w.marginTop
}

func (w *Writer) RenderWidget(widget Widget) error {
	context := w.NewBuildContext()

	err := widget.Render(context)
	if err != nil {
		return err
	}

	return nil
}

func (w *Writer) getNbAlias() string {
	return "{nb}"
}

func (w *Writer) SetFooter(
	handler func(page int, totalPagesAlias string) Widget,
) {
	w.computeAndSetFooterHeight(handler)

	w.Pdf.SetFooterFunc(func() {
		page := w.Pdf.PageNo()
		w.ignorePageBreak = true

		footerWidget := handler(page, w.getNbAlias())
		context := w.NewBuildContext()

		w.SetYWithoutPageBreakCheck(w.MaxHeight() + w.marginTop)

		footerWidget.Render(context)

		w.ignorePageBreak = false
	})
}

func (w *Writer) computeAndSetFooterHeight(
	handler func(page int, totalPagesAlias string) Widget,
) {
	footerWidget := handler(0, w.getNbAlias())
	context := w.NewBuildContext()

	_, actualHeight := footerWidget.CalculateSize(context)

	w.footerHeight = actualHeight

	if debug {
		fmt.Println("[DEBUG] Footer height:", w.footerHeight)
	}
}

func (w *Writer) GeneratePdf(fileName string) error {
	return w.Pdf.OutputFileAndClose(fileName)
}

func (w *Writer) GetStringSize(text string, fontSize float64, maxWidth float64) (float64, float64) {
	w.Pdf.SetFontSize(fontSize)
	lines := w.Pdf.SplitText(text, maxWidth)

	height := fontSize * float64(len(lines))
	width := 0.0

	for _, line := range lines {
		lineWidth := w.Pdf.GetStringWidth(line)
		if lineWidth > width {
			width = lineWidth
		}
	}

	return width, height
}

func (w *Writer) WriteMultiline(
	width float64,
	text string,
	fontSize float64,
	color color.RGBA,
) {

	if debug {
		fmt.Println("[DEBUG] WriteMultiline: w:", width, "text:", text, "fontSize:", fontSize)
	}

	w.Pdf.SetFontUnitSize(fontSize)
	w.Pdf.SetXY(w.x, w.y)
	w.Pdf.SetTextColor(int(color.R), int(color.G), int(color.B))
	w.Pdf.MultiCell(width, fontSize, text, "", "", false)
}

func (w *Writer) X() float64 {
	return w.x
}

func (w *Writer) Y() float64 {
	return w.y
}

func (w *Writer) WillWrite(width, height float64) {

	if debug {
		fmt.Println("[DEBUG] WillWrite: w:", width, "h:", height, "y:", w.y, "maxHeight:", w.MaxHeight())
	}

	if w.y+height > w.MaxHeight() {
		w.BreakPage()
	}

}

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
