package csv_handler

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type CSVRow struct {
	Row    []string
	Header bool
}

type ColumnsData struct {
	ColumnsMap  map[string]int
	ColumnsList []string
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type routes string

type ClientStruct struct {
	Client          HTTPClient
	RequestEndpoint string
}

type RequestData struct {
	Payload interface{}
	Path    string
	Method  string
	Client  *ClientStruct
}

func newRequestData(reqPayload interface{}, path, method string, client *ClientStruct) RequestData {
	return RequestData{
		Payload: reqPayload,
		Path:    path,
		Method:  method,
		Client:  client,
	}
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

func (cd *ColumnsData) processCsv(lineChannel chan CSVRow, languageCodeMap map[string]string) {

}

func makeColData(expectedColumns []string) ColumnsData {
	return ColumnsData{
		ColumnsList: expectedColumns,
		ColumnsMap:  make(map[string]int),
	}
}

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

const apiAddress = "http://www.7timer.info/bin/api.pl"

func createRequestStruct(requestEndpoint string) *ClientStruct {
	timeout := 10 * time.Second
	transport := http.Transport{
		DisableKeepAlives: true,
		MaxIdleConns:      20,
		MaxConnsPerHost:   100,
		IdleConnTimeout:   20,
	}
	client := http.Client{
		Timeout:   timeout,
		Transport: &transport,
	}
	return &ClientStruct{
		Client:          &client,
		RequestEndpoint: requestEndpoint,
	}
}

func (rd *RequestData) makeRequest(i interface{}) error {
	body, err := json.Marshal(rd.Payload)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(rd.Method, fmt.Sprintf("%v/%v", rd.Client.RequestEndpoint, rd.Path), bytes.NewBuffer(body))
	q := request.URL.Query()
	q.Add("lat", "51.5074")
	q.Add("lon", "0.1278")
	q.Add("unit", "metric")
	q.Add("output", "json")

	response, err := rd.Client.Client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	/*
		Alternative using json package instead of ioutil
		var data interface{}
		json.NewDecoder(response.Body).Decode(&data)
	*/
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if responseBody != nil {
		err = json.Unmarshal(responseBody, i)
	}
	return err
}

func useCSV() {

	// read csv file name
	csvFileName := "input.csv"
	clientStruct := createRequestStruct(apiAddress)

	weatherRequestData := newRequestData(nil, "", "GET", clientStruct)
	var data interface{}
	err := weatherRequestData.makeRequest(&data)
	if err != nil {
		log.Fatalf("failed retrieving weather data: %v", err)
	}

	// CSV
	rowChannel := make(chan CSVRow, 10)
	colData := makeColData([]string{"name", "language"})
	csvFile, err := os.Open(csvFileName)
	if err != nil {
		log.Fatalf("couldn't open the csv file: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		csvRdr := csv.NewReader(csvFile)
		err := colData.ReadCSV(csvRdr, rowChannel)
		if err != nil {
			log.Printf("error while reading lines: %v", err)
		}
	}()

}
