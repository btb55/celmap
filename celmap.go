// CelMap (Excel-Map) is used for mapping Excel file cells
// to and from their string form, and to and from their
// index form.
package celmap

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/plandem/ooxml"
)

// GetIndex will get the column and row number for a given Excel cell string.
// C22 returns 2, 21
func GetIndex(ref string) (col, row int) {
	colPart := strings.Map(ooxml.GetLettersFn, ref)
	rowPart := strings.Map(ooxml.GetNumbersFn, ref)

	var colIndex, rowIndex int
	for i, j := len(colPart)-1, 0; i >= 0; i, j = i-1, j+1 {
		colIndex += (int(colPart[i]) - int('A') + 1) * int(math.Pow(26, float64(j)))
	}

	rowIndex, _ = strconv.Atoi(rowPart)

	rowIndex--
	colIndex--

	return colIndex, rowIndex
}

// GetRef will get a cell reference for 0-based indexes.
// 2, 21 returns C22
func GetRef(col, row int) (ref string) {
	if col < 0 || row < 0 {
		return ""
	}

	var colName string
	i := col + 1
	for i > 0 {
		colName = string(rune(((i-1)%26 + 65))) + colName
		i = (i - 1) / 26
	}

	return fmt.Sprintf("%s%d", colName, row+1)
}
