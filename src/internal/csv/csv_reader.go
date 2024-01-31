package csv

import (
	"encoding/csv"
	"os"
)

func ReadAllRecords(filePath string, delimiter rune, ignoreFirstLine bool) [][]string {
	reader, file := getCSVReader(filePath, delimiter)
	defer file.Close()

	if ignoreFirstLine {
		if _, err := reader.Read(); err != nil {
			panic(err)
		}
	}

	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	return records
}

func getCSVReader(filePath string, demiliter rune) (*csv.Reader, *os.File) {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = demiliter

	return reader, file
}

func GetColumnValues(records [][]string, columnIndex int) []string {
	left := make([]string, len(records))

	for i, record := range records {		
		left[i] = record[columnIndex]
	}

	return left
}