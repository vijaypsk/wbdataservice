package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Data struct {
	Year  int64   `json:"year"`
	Value float64 `json:"value"`
}
type YearlyData []Data

var alldata YearlyData

type Indicator struct {
	IndicatorName string     `json:iname`
	IndicatorCode string     `json:icode`
	CountryName   string     `json:cname`
	CountryCode   string     `json:ccode`
	YData         YearlyData `json:ydata`
}

type IndicatorDefiniton struct {
	IndicatorCode  string `json:icode`
	IndicatorName  string `json:iname`
	LongDefinition string `json:longdef`
	Source         string `json:source`
}

type Indicators []Indicator

var indicators Indicators

type IndicatorDefnitions []IndicatorDefiniton

var indicatordefs IndicatorDefnitions

// Give us some seed data
func init() {
	readData()
	readIndicatorDefinitions()
}

func readData() {

	file, err := os.Open("./indicators.csv")

	// Create a new reader.

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(bufio.NewReader(file))
	isHeader := true
	counter := 0
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		if isHeader {
			isHeader = false
		} else {
			rindex := 4
			var yr int64
			alldata := make(YearlyData, 12, 15)
			for yr = 2000; yr <= 2012; yr++ {
				strval := record[rindex]
				if strval == "NA" || strval == "" {
					strval = "0"
				}
				val, err := strconv.ParseFloat(strval, 64)
				if err != nil {
					log.Fatal(err)
				}
				alldata = append(alldata, Data{Year: yr, Value: val})
				rindex = rindex + 1
			}

			i := Indicator{IndicatorName: record[0],
				IndicatorCode: record[1],
				CountryName:   record[2],
				CountryCode:   record[3],
				YData:         alldata}
			indicators = append(indicators, i)
			counter = counter + 1

		}

	}

}

func readIndicatorDefinitions() {

	file, err := os.Open("./definitions.csv")

	// Create a new reader.

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(bufio.NewReader(file))
	isHeader := true
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		if isHeader {
			isHeader = false
		} else {

			i := IndicatorDefiniton{IndicatorName: record[1],
				IndicatorCode:  record[0],
				LongDefinition: record[2],
				Source:         record[3]}
			indicatordefs = append(indicatordefs, i)

		}

	}

}
