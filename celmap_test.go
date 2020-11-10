package celmap

import (
	"testing"
)

func TestGetIndex(t *testing.T) {
	var (
		expectedFirstCol  = 2
		expectedFirstRow  = 21
		expectedSecondCol = 28
		expectedSecondRow = 1021
	)

	actualFirstCol, actualFirstRow := GetIndex("C22")
	actualSecondCol, actualSecondRow := GetIndex("AC1022")

	if actualFirstCol != expectedFirstCol {
		t.Fatalf("Expected %v for first column but got %v", expectedFirstCol, actualFirstCol)
	} else if actualFirstRow != expectedFirstRow {
		t.Fatalf("Expected %v for first row but got %v", expectedFirstRow, actualFirstRow)
	} else if actualSecondCol != expectedSecondCol {
		t.Fatalf("Expected %v for second column but got %v", expectedSecondCol, actualSecondCol)
	} else if actualSecondRow != expectedSecondRow {
		t.Fatalf("Expected %v for second row but got %v", expectedSecondRow, actualSecondRow)
	}
}

func TestGetRef(t *testing.T) {
	var (
		expectedFirstString  = "C22"
		expectedSecondString = "AC1022"
	)

	actualFirstString := GetRef(2, 21)
	actualSecondString := GetRef(28, 1021)

	if actualFirstString != expectedFirstString {
		t.Fatalf("Expected %s for first ref but got %s", expectedFirstString, actualFirstString)
	} else if actualSecondString != expectedSecondString {
		t.Fatalf("Expected %v for second ref but got %v", expectedSecondString, actualSecondString)
	}
}
