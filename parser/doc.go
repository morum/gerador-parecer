package parser

import (
	"fmt"
	"log/slog"

	"github.com/nguyenthenguyen/docx"
)

type ReplacementParam string

func (r ReplacementParam) String() string {
	return string(r)
}

type Replacement struct {
	Old ReplacementParam
	New string
}

// ReplaceAndSave substitutes the text in the docx file and saves it.
func ReplaceAndSaveDoc(docxFile, newFile string, replacements ...Replacement) error {
	file, err := docx.ReadDocxFile(docxFile)
	if err != nil {
		return err
	}
	defer file.Close()

	doc := file.Editable()

	for _, replacement := range replacements {
		replacementFormated := fmt.Sprintf("{{%s}}", replacement.Old)
		err = doc.Replace(replacementFormated, replacement.New, -1)
		if err != nil {
			slog.Error("Erro ao substituir texto", "erro", err)
		}
	}

	err = doc.WriteToFile(newFile)
	if err != nil {
		return err
	}
	slog.Info("Arquivo gerado com sucesso", "arquivo", newFile)
	return nil
}
