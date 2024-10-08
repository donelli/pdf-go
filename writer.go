package tpdf

import (
	"bufio"
	"bytes"
	"fmt"
	"image/color"

	"github.com/go-pdf/fpdf"
)

const debug = false
const debugDrawPageMargin = false
const debugDrawFooterArea = false
const debugDrawTextBounds = false
const debugDrawRectBounds = false

type Writer struct {
	Pdf             *fpdf.Fpdf
	internalX       float64
	internalY       float64
	footerHeight    float64
	marginLeft      float64
	marginRight     float64
	marginTop       float64
	marginBottom    float64
	ignorePageBreak bool
	footerRenderer  func(page int, totalPagesAlias string) Widget
	offsetX         float64
	offsetY         float64
	currentCapStyle DividerCapStyle
	theme           *Theme
	mainWidget      Widget
	metadata        *Metadata
}

func NewWriter(
	topMargin, rightMargin, bottomMargin, leftMargin float64,
	theme *Theme,
) *Writer {
	pdf := fpdf.New("P", "pt", "A4", "")
	pdf.SetFont("Arial", "", 14)
	pdf.SetCellMargin(0)
	pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)

	w := &Writer{
		Pdf:             pdf,
		internalX:       leftMargin,
		internalY:       topMargin,
		marginLeft:      leftMargin,
		marginRight:     rightMargin,
		marginTop:       topMargin,
		marginBottom:    bottomMargin,
		ignorePageBreak: false,
		theme:           theme,
	}

	w.SetLineCapStyle(w.currentCapStyle)

	pdf.AliasNbPages(w.getNbAlias())

	if debug {
		pageWidth, pageHeight := pdf.GetPageSize()
		fmt.Println("[DEBUG] Created writer. page width", pageWidth, "page height", pageHeight, "max width", w.MaxWidth(), "max height", w.MaxHeight())
	}

	return w
}

func (w *Writer) AddPage() {
	w.Pdf.AddPage()
	w.internalY = w.marginTop

	w.computeFooterHeight(w.Pdf.PageNo())

	if debug {
		fmt.Println("[DEBUG] AddPage")
	}

	if debugDrawPageMargin {
		w.ignorePageBreak = true

		pageWidth, pageHeight := w.Pdf.GetPageSize()

		w.setFillColor(color.RGBA{255, 140, 140, 255})

		w.Pdf.Rect(0, 0, pageWidth, w.marginTop, "F")
		w.Pdf.Rect(0, 0, w.marginLeft, pageHeight, "F")
		w.Pdf.Rect(pageWidth-w.marginRight, 0, w.marginRight, pageHeight, "F")
		w.Pdf.Rect(0, pageHeight-w.marginBottom, pageWidth, w.marginBottom, "F")

		w.ignorePageBreak = false
	}

	if debugDrawFooterArea {
		w.ignorePageBreak = true
		pageWidth, pageHeight := w.Pdf.GetPageSize()

		w.setFillColor(color.RGBA{176, 176, 255, 255})

		x := w.marginLeft
		y := pageHeight - w.marginBottom - w.footerHeight
		width := pageWidth - w.marginLeft - w.marginRight
		height := w.footerHeight

		w.Pdf.Rect(x, y, width, height, "F")

		w.ignorePageBreak = false
	}

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

	w.internalX = x
}

func (w *Writer) SetY(y float64) {
	if debug {
		fmt.Println("[DEBUG] SetY:", y)
	}

	w.internalY = y

	if w.Y() >= w.MaxHeight() {
		w.BreakPage()
	}
}

func (w *Writer) SetYWithoutPageBreakCheck(y float64) {
	if debug {
		fmt.Println("[DEBUG] SetYWithoutPageBreakCheck:", y)
	}

	w.internalY = y
}

func (w *Writer) BreakPage() {
	if w.ignorePageBreak {
		return
	}

	if debug {
		fmt.Println("[DEBUG] BreakPage")
	}

	w.AddPage()
}

func (w *Writer) SetMainWidget(widget Widget) {
	w.mainWidget = widget
}

func (w *Writer) getNbAlias() string {
	return "{nb}"
}

func (w *Writer) SetFooter(
	handler func(page int, totalPagesAlias string) Widget,
) {

	renderer := func() {
		page := w.Pdf.PageNo()
		w.ignorePageBreak = true

		footerWidget := handler(page, w.getNbAlias())
		context := w.NewBuildContext()

		w.SetYWithoutPageBreakCheck(w.MaxHeight() + w.marginTop)

		footerWidget.Render(context)

		w.ignorePageBreak = false
	}

	w.Pdf.SetFooterFunc(renderer)
	w.footerRenderer = handler
}

func (w *Writer) computeFooterHeight(pageNumber int) {
	if w.footerRenderer == nil {
		return
	}

	footerWidget := w.footerRenderer(pageNumber, w.getNbAlias())
	context := w.NewBuildContext()

	_, footerHeight := footerWidget.CalculateSize(context)

	w.footerHeight = footerHeight

	if debug {
		fmt.Println("[DEBUG] Set footer height of ", footerHeight, " for page ", pageNumber)
	}
}

func (w *Writer) SetMetadata(metadata Metadata) {
	w.metadata = &metadata
}

func (w *Writer) generate() error {
	if w.metadata != nil {
		w.metadata.addMetadataToPdf(w.Pdf)
	}

	w.AddPage()
	context := w.NewBuildContext()

	err := w.mainWidget.Render(context)
	if err != nil {
		return err
	}

	return nil
}

func (w *Writer) GenerateToFile(fileName string) error {
	err := w.generate()

	if err != nil {
		return err
	}

	return w.Pdf.OutputFileAndClose(fileName)
}

func (w *Writer) GenerateToBuffer(fileName string) (*bytes.Buffer, error) {
	err := w.generate()

	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	bufferWriter := bufio.NewWriter(&buffer)

	err = w.Pdf.Output(bufferWriter)

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
	maxLines int,
) (float64, float64) {
	w.setFontStyles(fontSize, color, bold, italic, underline, strikeOut)
	lines := w.Pdf.SplitText(text, maxWidth)

	height := fontSize * float64(len(lines))
	width := 0.0

	for lineIndex, line := range lines {
		lineWidth := w.Pdf.GetStringWidth(line)
		if lineWidth > width {
			width = lineWidth
		}

		if maxLines > 0 && lineIndex == maxLines-1 {
			stringWithEllipsisWidth := w.Pdf.GetStringWidth(line + "...")

			if stringWithEllipsisWidth < maxWidth {
				width = stringWithEllipsisWidth
			} else {
				width = maxWidth
			}

			break
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
	fontColor color.Color,
	bold bool,
	italic bool,
	underline bool,
	strikeOut bool,
	textAlign TextAlign,
	link string,
	maxLines int,
) {

	w.Pdf.SetXY(w.X(), w.Y())

	w.setFontStyles(fontSize, fontColor, bold, italic, underline, strikeOut)

	textAlignStr := w.textAlignToPdfStr(textAlign)

	if debug {
		fmt.Println("[DEBUG] WriteMultiline: w:", width, "text:", text, "textAlign:", textAlignStr)
	}

	lines := w.Pdf.SplitText(text, width)

	x := w.X()
	for lineIndex, line := range lines {

		shouldAddEllipsis := maxLines > 0 && lineIndex == maxLines-1 && len(lines) > maxLines

		if shouldAddEllipsis {
			for i := range 3 {
				textWithEllipsis := line[0:len(line)-i] + "..."

				textSize := w.Pdf.GetStringWidth(textWithEllipsis)

				if textSize <= width {
					line = textWithEllipsis
					break
				}
			}
		}

		if debug {
			fmt.Println("[DEBUG] CellFormat: text:", line, "width", width)
		}

		borderStr := ""
		if debugDrawTextBounds {
			w.setDrawColor(color.RGBA{0, 0, 255, 255})
			borderStr = "1"
		}

		w.Pdf.SetX(x)
		w.Pdf.CellFormat(width, fontSize, line, borderStr, 0, textAlignStr, false, 0, "")
		w.Pdf.Ln(-1)

		if maxLines > 0 && lineIndex == maxLines-1 {
			break
		}
	}

	endY := w.Pdf.GetY()

	if link != "" {
		w.Pdf.LinkString(w.X(), w.Y(), width, endY-w.Y(), link)
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
	w.SetFontUnitSize(fontSize)

	w.setTextColor(color)

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

	w.SetFontStyle(styleString)

	if debug {
		fmt.Println("[DEBUG] setFontStyles: fontSize:", fontSize, "color:", color, "bold:", bold, "italic:", italic, "underline:", underline, "strikeOut:", strikeOut)
	}
}

func (w *Writer) X() float64 {
	return w.internalX + w.offsetX
}

func (w *Writer) Y() float64 {
	return w.internalY + w.offsetY
}

func (w *Writer) WillWrite(width, height float64) {

	if debug {
		fmt.Println("[DEBUG] WillWrite: w:", width, "h:", height, "y:", w.Y(), "maxHeight:", w.MaxHeight())
	}

	if w.Y()+height > w.MaxHeight() {
		w.BreakPage()
	}

}

func (w *Writer) Line(
	width float64,
	height float64,
	color color.Color,
	lineWidth float64,
	capStyle DividerCapStyle,
) {
	w.setDrawColor(color)
	w.setLineWidth(lineWidth)
	w.SetLineCapStyle(capStyle)

	w.Pdf.Line(w.X(), w.Y(), w.X()+width, w.Y()+height)
}

func (w *Writer) Rect(
	width, height float64,
	backgroundColor color.Color,
	borderColor color.Color, borderWidth float64,
) {
	styleStr := ""

	if backgroundColor != nil {
		w.setFillColor(backgroundColor)
		styleStr += "F"
	}

	if borderColor != nil {
		w.setDrawColor(borderColor)
		w.setLineWidth(borderWidth)
		styleStr += "D"
	}

	if styleStr != "" {
		w.Pdf.Rect(w.X(), w.Y(), width, height, styleStr)
	}

	if debugDrawRectBounds {
		w.setDrawColor(color.RGBA{0, 0, 255, 255})
		w.Pdf.Rect(w.X(), w.Y(), width, height, "D")
	}
}

func (w *Writer) RoundedRect(
	width, height float64,
	backgroundColor color.Color,
	borderRadius borderRadius,
	borderColor color.Color,
	borderWidth float64,
) {
	styleStr := ""

	if backgroundColor != nil {
		w.setFillColor(backgroundColor)
		styleStr += "F"
	}

	if borderColor != nil {
		w.setLineWidth(borderWidth)
		w.setDrawColor(borderColor)
		styleStr += "D"
	}

	if styleStr != "" {
		w.Pdf.RoundedRectExt(w.X(), w.Y(),
			width, height,
			borderRadius.TopLeft(), borderRadius.TopRight(),
			borderRadius.BottomRight(), borderRadius.BottomLeft(),
			styleStr,
		)
	}

	if debugDrawRectBounds {
		w.setDrawColor(color.RGBA{0, 0, 255, 255})
		w.Pdf.Rect(w.X(), w.Y(), width, height, "D")
	}
}

func goColorToRGBA(color color.Color) (r, g, b, a int) {
	multipliedR, multipliedG, multipliedB, _ := color.RGBA()
	return int(multipliedR / 255), int(multipliedG / 255), int(multipliedB / 255), 255
}

func (w *Writer) setFillColor(color color.Color) {
	r, g, b, _ := goColorToRGBA(color)

	currentR, currentG, currentB := w.Pdf.GetFillColor()
	if currentR == r && currentG == g && currentB == b {
		return
	}

	w.Pdf.SetFillColor(r, g, b)
}

func (w *Writer) setDrawColor(color color.Color) {
	r, g, b, _ := goColorToRGBA(color)

	currentR, currentG, currentB := w.Pdf.GetDrawColor()
	if currentR == r && currentG == g && currentB == b {
		return
	}

	w.Pdf.SetDrawColor(r, g, b)
}

func (w *Writer) setLineWidth(lineWidth float64) {
	if w.Pdf.GetLineWidth() == lineWidth {
		return
	}

	w.Pdf.SetLineWidth(lineWidth)
}

func (w *Writer) SetFontUnitSize(fontSize float64) {
	currentPtSize, _ := w.Pdf.GetFontSize()

	if currentPtSize == fontSize {
		return
	}

	w.Pdf.SetFontUnitSize(fontSize)
}

func (w *Writer) SetFontStyle(style string) {
	w.Pdf.SetFontStyle(style)
}

func (w *Writer) SetLineCapStyle(capStyle DividerCapStyle) {
	if w.currentCapStyle == capStyle {
		return
	}

	w.currentCapStyle = capStyle
	w.Pdf.SetLineCapStyle(capStyle.fpdfValue())
}

func (w *Writer) setTextColor(color color.Color) {
	r, g, b, _ := color.RGBA()

	w.Pdf.SetTextColor(int(r/255), int(g/255), int(b/255))
}

func (w *Writer) AddOffsets(offsetX, offsetY float64) {
	w.offsetX += offsetX
	w.offsetY += offsetY
}

func (w *Writer) SubtractOffsets(offsetX, offsetY float64) {
	w.offsetX -= offsetX
	w.offsetY -= offsetY
}

func (w *Writer) PageNumber() int {
	return w.Pdf.PageNo()
}
