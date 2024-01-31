package main

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"slices"

	"github.com/eriklima/csv-comparator/internal/csv"
	"github.com/eriklima/csv-comparator/internal/utils"
)

const LeftCsvName = "esquerda.csv"
const RightCsvName = "direita.csv"
const CsvDelimiter = ','

var filesPath = "arquivos"

func init() {
	currentFilePath := getCurrentFilePath()
	filesPath = filepath.Join(currentFilePath, filesPath)
}

func getCurrentFilePath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get current frame")
	}
	return path.Dir(filename)
}

func main() {
	leftCsvPath := getLeftCSVPath()
	leftRecords := csv.ReadAllRecords(leftCsvPath, CsvDelimiter, true)

	rightCsvPath := getRightCSVPath()
	rightRecords := csv.ReadAllRecords(rightCsvPath, CsvDelimiter, true)

	compare(leftRecords, rightRecords)
}

func getLeftCSVPath() string {
	return filepath.Join(filesPath, LeftCsvName)
}

func getRightCSVPath() string {
	return filepath.Join(filesPath, RightCsvName)
}

func compare(leftRecords [][]string, rightRecords [][]string) {
	leftColumnIndex := 0
	rightColumnIndex := 0

	leftValues := csv.GetColumnValues(leftRecords, leftColumnIndex)
	rightValues := csv.GetColumnValues(rightRecords, rightColumnIndex)

	leftDiff := getDiffValues(leftValues, rightValues)
	rightDiff := getDiffValues(rightValues, leftValues)
	
	fmt.Printf("Valores da ESQUERDA: %s\n", leftValues)
	fmt.Printf("Valores da DIREITA : %s\n", rightValues)

	fmt.Printf("\nValores FALTANDO na ESQUERDA: %s\n", leftDiff)
	fmt.Printf("Valores FALTANDO na DIREITA : %s\n", rightDiff)
}

func getDiffValues(leftValues []string, rightValues []string) []string {
	var diff []string

	valuesAreEquals := slices.Compare[[]string](leftValues, rightValues) == 0

	if !valuesAreEquals {

		criteriaForFilter := func(value string) bool {
			return !slices.Contains[[]string](leftValues, value)
		}

		diff = utils.SliceFilter[string](rightValues, criteriaForFilter)
	}	

	return diff
}