package main

import (
	"encoding/json"
	"time"
	tpdf "tpdf"
)

func main() {
	metadata := tpdf.Metadata{
		Title:            "Title",
		Author:           "Author",
		CreationDate:     time.Now(),
		ModificationDate: time.Now(),
		Creator:          "Creator",
		Keywords:         []string{"keyword1", "keyword2"},
		Subject:          "Subject",
		Language:         "en",
		Producer:         "Producer",
	}

	metadataJSON, _ := json.MarshalIndent(metadata, "", "  ")

	content := tpdf.Column(
		tpdf.Text(string(metadataJSON)),
	).WithSpacing(8)

	theme := tpdf.NewTheme()
	writer := tpdf.NewWriter(8, 8, 8, 8, theme)

	writer.SetMetadata(metadata)
	writer.SetMainWidget(content)

	writer.GenerateToFile("metadata.pdf")
}
