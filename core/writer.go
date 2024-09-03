package core

import (
	"fmt"

	"github.com/go-pdf/fpdf"
)

const debug = true

type Writer struct {
	Pdf *fpdf.Fpdf
	x   float64
	y   float64
}

func NewWriter() *Writer {
	pdf := fpdf.New("P", "pt", "A4", "")
	pdf.SetFont("Arial", "", 14)
	pdf.SetCellMargin(0)
	pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)

	pdf.AddPage()

	return &Writer{
		Pdf: pdf,
		x:   4,
		y:   4,
	}
}

func (w *Writer) MaxWidth() float64 {
	pageWidth, _ := w.Pdf.GetPageSize()
	return pageWidth
}

func (w *Writer) MaxHeight() float64 {
	_, pageHeight := w.Pdf.GetPageSize()

	return pageHeight
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
func (w *Writer) BreakPage() {
	if debug {
		fmt.Println("[DEBUG] BreakPage")
	}

	w.Pdf.AddPage()
	w.y = 4
}

func (w *Writer) RenderWidget(widget Widget) error {
	context := w.NewBuildContext()

	err := widget.Render(context)
	if err != nil {
		return err
	}

	return nil
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

func (w *Writer) WriteMultiline(width float64, text string, fontSize float64) {

	if debug {
		fmt.Println("[DEBUG] WriteMultiline: w:", width, "text:", text, "fontSize:", fontSize)
	}

	w.Pdf.SetFontUnitSize(fontSize)
	w.Pdf.SetXY(w.x, w.y)
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
