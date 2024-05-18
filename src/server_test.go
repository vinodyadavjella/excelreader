package main

import (
	"os"
	"testing"
)

func TestProcessExcelFile(t *testing.T) {

	excelFilePath := "C:/Users/vinod/source/repos/GO/excelreader/src/sample_test.xlsx"
	testData, err := os.ReadFile(excelFilePath)

	if err != nil {
		t.Errorf("failed to read test excel file with error: %v", err)
	}
	// []byte(`
	// 	Name, Age, City
	// 	John, 30, New York
	// 	Alice, 25, Los Angeles
	// `)

	jsonData, err := processExcelFile(testData)

	if err != nil {
		t.Errorf("processExcelFile() returned an unexpected error: %v", err)
	}

	expectedNumberOfRows := 4

	if len(jsonData) != expectedNumberOfRows {
		t.Errorf("processExcelFile() returned an incorrect number of JSON string: got %d, want %d", len(jsonData), expectedNumberOfRows)
	}

	expectedFirstRowJSON := `{"cell0":"John","cell1":"30","cell2":"New York"}`
	if jsonData[1] != expectedFirstRowJSON {
		t.Errorf("processExcelFile() returned an incorrect JSON for the first row: got %s, want %s", jsonData[1], expectedFirstRowJSON)
	}
	expectedSecondRowJSON := `{"cell0":"Alice","cell1":"25","cell2":"Los Angeles"}`
	if jsonData[2] != expectedSecondRowJSON {
		t.Errorf("processExcelFile() returned an incorrect JSON for the Second row: got %s, want %s", jsonData[2], expectedSecondRowJSON)
	}

	expectedThirdRowJSON := `{"cell0":"Mounkia","cell1":"10","cell2":"Aubrey"}`
	if jsonData[3] != expectedThirdRowJSON {
		t.Errorf("processExcelFile() returned an incorrect JSON for the Third row: got %s, want %s", jsonData[3], expectedThirdRowJSON)
	}

}
