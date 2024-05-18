package main

import (
	"encoding/json"
	"fmt"

	"github.com/tealeg/xlsx"
)

func processExcelFile(fileData []byte) ([]string, error) {
	file, err := xlsx.OpenBinary(fileData)

	if err != nil {
		return nil, err
	}

	var jsonStrings []string
	for _, sheet := range file.Sheets {
		for _, row := range sheet.Rows {
			cells := make(map[string]string)
			for i, cell := range row.Cells {
				cellValue := cell.String()
				cells[fmt.Sprintf("cell%d", i)] = cellValue
			}

			jsonStr, _ := json.Marshal(cells)
			jsonStrings = append(jsonStrings, string(jsonStr))
		}
	}

	return jsonStrings, nil
}
