package parser

import (
	"log/slog"
	"strings"

	"github.com/xuri/excelize/v2"
)

type ReplaceableField struct {
	VariableName string
	Value        string
}

type Data struct {
	Fields []ReplaceableField
}

func GetDataFromXlsx(file, sheet string) ([]Data, error) {
	f, err := excelize.OpenFile(file)
	if err != nil {
		return nil, err
	}

	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	vars := make([]string, 0)
	for _, cell := range rows[0] {
		vars = append(vars, cell)
	}

	data := make([]Data, 0)
	for _, row := range rows[1:] {
		d := Data{}
		for i, cell := range row {
			d.Fields = append(d.Fields, ReplaceableField{
				VariableName: formatString(vars[i]),
				Value:        formatString(cell),
			})
		}
		data = append(data, d)
	}

	slog.Info("Xlsx extraído com sucesso", "quantidade", len(data))

	return data, nil
}

func formatString(s string) string {
	return strings.TrimSpace(s)
}
