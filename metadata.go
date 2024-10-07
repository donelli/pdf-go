package tpdf

import (
	"strings"
	"time"

	"github.com/go-pdf/fpdf"
)

type Metadata struct {
	Title            string
	Author           string
	CreationDate     time.Time
	ModificationDate time.Time
	Creator          string
	Keywords         []string
	Subject          string
	Language         string
	Producer         string
}

func (m *Metadata) addMetadataToPdf(pdf *fpdf.Fpdf) {

	if m.Title != "" {
		pdf.SetTitle(m.Title, true)
	}

	if m.Author != "" {
		pdf.SetAuthor(m.Author, true)
	}

	emptyTime := time.Time{}

	if m.CreationDate != emptyTime {
		pdf.SetCreationDate(m.CreationDate)
	}

	if m.ModificationDate != emptyTime {
		pdf.SetModificationDate(m.ModificationDate)
	}

	if m.Creator != "" {
		pdf.SetCreator(m.Creator, true)
	}

	if len(m.Keywords) > 0 {
		pdf.SetKeywords(strings.Join(m.Keywords, " "), true)
	}

	if m.Subject != "" {
		pdf.SetSubject(m.Subject, true)
	}

	if m.Language != "" {
		pdf.SetLang(m.Language)
	}

	if m.Producer != "" {
		pdf.SetProducer(m.Producer, true)
	}
}
