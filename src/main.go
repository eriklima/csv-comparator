package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/eriklima/csv-comparator/internal/csv"
	"github.com/eriklima/csv-comparator/internal/utils"
)

const LeftCsvName = "esquerda.csv"
const RightCsvName = "direita.csv"
const CsvDelimiter = ','

var CurrentPath string
var FilesPath = "../arquivos"

func init() {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	CurrentPath = dir
	FilesPath = filepath.Join(CurrentPath, FilesPath)
}

func main() {
	leftCsvPath := getLeftCSVPath()
	leftRecords := csv.ReadAllRecords(leftCsvPath, CsvDelimiter, true)

	rightCsvPath := getRightCSVPath()
	rightRecords := csv.ReadAllRecords(rightCsvPath, CsvDelimiter, true)

	compare(leftRecords, rightRecords)
}

func getLeftCSVPath() string {
	return filepath.Join(FilesPath, LeftCsvName)
}

func getRightCSVPath() string {
	return filepath.Join(FilesPath, RightCsvName)
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