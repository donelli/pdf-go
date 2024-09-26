package core

import (
	"bufio"
	"bytes"
	"fmt"
	"image/color"

	"github.com/go-pdf/fpdf"
)

const debug = true

type Writer struct {
	Pdf              *fpdf.Fpdf
	x                float64
	y                float64
	footerHeight     float64
	marginLeft       float64
	marginRight      float64
	marginTop        float64
	marginBottom     float64
	ignorePageBreak  bool
	defaultFontSize  float64
	defaultFontColor color.Color
}

func NewWriter(topMargin, rightMargin, bottomMargin, leftMargin float64) *Writer {
	pdf := fpdf.New("P", "pt", "A4", "")
	pdf.SetFont("Arial", "", 14)
	pdf.SetCellMargin(0)
	pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)

	pdf.AddPage()

	w := &Writer{
		Pdf:              pdf,
		x:                leftMargin,
		y:                topMargin,
		marginLeft:       leftMargin,
		marginRight:      rightMargin,
		marginTop:        topMargin,
		marginBottom:     bottomMargin,
		ignorePageBreak:  false,
		defaultFontSize:  14,
		defaultFontColor: color.RGBA{0, 0, 0, 255},
	}

	pdf.AliasNbPages(w.getNbAlias())

	if debug {
		pageWidth, pageHeight := pdf.GetPageSize()
		fmt.Println("[DEBUG] Created writer. page width", pageWidth, "page height", pageHeight, "max width", w.MaxWidth(), "max height", w.MaxHeight())
	}

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

	maxWidth := w.MaxWidth()
	maxHeight := w.MaxHeight()

	if debug {
		fmt.Println("[DEBUG] NewBuildContext: maxWidth:", maxWidth, "maxHeight:", maxHeight)
	}

	return &RenderContext{
		Writer:    w,
		MaxWidth:  maxWidth,
		MaxHeight: maxHeight,
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

func (w *Writer) GeneratePdfToFile(fileName string) error {
	return w.Pdf.OutputFileAndClose(fileName)
}

func (w *Writer) GeneratePdfToBuffer(fileName string) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	bufferWriter := bufio.NewWriter(&buffer)

	err := w.Pdf.Output(bufferWriter)

	if err != nil {
		return nil, err
	}

	return &buffer, nil
}

func (w *Writer) GetStringSize(
	text string,
	fontSize float64,
	maxWidth float64,
	color color.Color,
	bold bool,
	italic bool,
	underline bool,
	strikeOut bool,
) (float64, float64) {
	w.setFontStyles(fontSize, color, bold, italic, underline, strikeOut)
	lines := w.Pdf.SplitText(text, maxWidth)

	height := fontSize * float64(len(lines))
	width := 0.0

	for _, line := range lines {
		lineWidth := w.Pdf.GetStringWidth(line)
		if lineWidth > width {
			width = lineWidth
		}
	}

	if debug {
		fmt.Println("[DEBUG] GetStringSize: text:", text, "fontSize:", fontSize, "width:", width, "height:", height)
	}

	return width, height
}

func (w *Writer) WriteMultiline(
	width float64,
	text string,
	fontSize float64,
	color color.Color,
	bold bool,
	italic bool,
	underline bool,
	strikeOut bool,
	textAlign TextAlign,
	link string,
) {

	w.Pdf.SetXY(w.x, w.y)

	w.setFontStyles(fontSize, color, bold, italic, underline, strikeOut)

	textAlignStr := w.textAlignToPdfStr(textAlign)

	if debug {
		fmt.Println("[DEBUG] WriteMultiline: w:", width, "text:", text, "textAlign:", textAlignStr)
	}

	lines := w.Pdf.SplitText(text, width)

	x := w.x
	for _, line := range lines {
		w.Pdf.SetX(x)
		w.Pdf.CellFormat(width, fontSize, line, "", 0, textAlignStr, false, 0, "")
		w.Pdf.Ln(-1)
	}

	endY := w.Pdf.GetY()

	if link != "" {
		w.Pdf.LinkString(w.x, w.y, width, endY-w.y, link)
	}
}

func (w *Writer) textAlignToPdfStr(textAlign TextAlign) string {
	switch textAlign {
	case TextAlignLeft:
		return "L"
	case TextAlignCenter:
		return "C"
	case TextAlignRight:
		return "R"
	}

	return ""
}

func (w *Writer) setFontStyles(
	fontSize float64,
	color color.Color,
	bold bool,
	italic bool,
	underline bool,
	strikeOut bool,
) {
	w.Pdf.SetFontUnitSize(fontSize)

	r, g, b, _ := color.RGBA()

	w.Pdf.SetTextColor(int(r), int(g), int(b))

	styleString := ""

	if bold {
		styleString += "B"
	}

	if italic {
		styleString += "I"
	}

	if underline {
		styleString += "U"
	}

	if strikeOut {
		styleString += "S"
	}

	w.Pdf.SetFontStyle(styleString)

	if debug {
		fmt.Println("[DEBUG] setFontStyles: fontSize:", fontSize, "color:", color, "bold:", bold, "italic:", italic, "underline:", underline, "strikeOut:", strikeOut)
	}
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

func (w *Writer) SetDefaultFontSize(fontSize float64) {
	w.defaultFontSize = fontSize
}

func (w *Writer) DefaultFontSize() float64 {
	return w.defaultFontSize
}

func (w *Writer) SetDefaultFontColor(color color.Color) {
	w.defaultFontColor = color
}

func (w *Writer) DefaultFontColor() color.Color {
	return w.defaultFontColor
}
