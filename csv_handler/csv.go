package csv_handler

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strings"
)

type CSVRow struct {
	Row    []string
	Header bool
}

type ColumnsData struct {
	ColumnsMap  map[string]int
	ColumnsList []string
}

var CSVParseErr = errors.New("csv parse error")

func CSVParseError(v interface{}) error {
	return fmt.Errorf("%v: %w", v, CSVParseErr)
}

func (cd *ColumnsData) ReadCSV(csvRdr *csv.Reader, lineChan chan CSVRow) error {
	defer close(lineChan)

	// reader header
	header, err := csvRdr.Read()
	if err != nil {
		return CSVParseError(err)
	}

	if len(header) == 1 {
		csvRdr.Comma = '\t'
		header = strings.Split(header[0], "\t")
		csvRdr.FieldsPerRecord = len(header)
	}

	// read header and validate all required columns are in place
	missingCol := cd.makeColMap(header)
	if len(missingCol) > 0 {
		return CSVParseError(fmt.Sprintf("missing columns: %v", strings.Join(missingCol, ", ")))
	}

	var lineErrors []error
	// read the rest of the lines
	for {
		line, err := csvRdr.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			lineErrors = append(lineErrors, err)
			continue
		}

		// send only rows that haven't errored through
		lineChan <- CSVRow{
			Row:    line,
			Header: false,
		}
	}
	if len(lineErrors) > 0 {
		return CSVParseError(lineErrors)
	}
	return nil
}

func (cd *ColumnsData) processCsv(lineChannel chan CSVRow, languageCodeMap map[string]string)

func (cd *ColumnsData) makeColMap(row []string) []string {
	for number, column := range row {
		cd.ColumnsMap[column] = number
	}
	var missingCol []string
	for _, col := range cd.ColumnsList {
		if _, ok := cd.ColumnsMap[col]; !ok {
			missingCol = append(missingCol, col)
		}
	}
	return missingCol
}
