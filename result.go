package main

import (
	"fmt"
	"sort"
)

type Result struct {
	Value float64 `json:"value"`
	IName string  `json:"iname"`
	CName string  `json:"cname"`
}
type Results []Result

var res Results

func (slice Results) Len() int {
	return len(slice)
}

func (slice Results) Less(i, j int) bool {
	return slice[i].Value > slice[j].Value
}

func (slice Results) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func GetIndicatorsFor(year int64) Results {
	res := make(Results, 0)

	for value := range indicators {

		for i := range indicators[value].YData {
			if year == indicators[value].YData[i].Year {
				r := Result{IName: indicators[value].IndicatorName, Value: indicators[value].YData[i].Value, CName: indicators[value].CountryName}
				fmt.Println(r)
				res = append(res, r)
			}
		}

	}
	sort.Sort(res)
	return res
}

func GetIndicatorDataFor(year int64, icode string) Results {
	res := make(Results, 0)
	fmt.Println(year, icode)
	for value := range indicators {

		if indicators[value].IndicatorCode == icode {
			fmt.Println(indicators[value].IndicatorCode)

			for i := range indicators[value].YData {
				if year == indicators[value].YData[i].Year {
					r := Result{IName: indicators[value].IndicatorName, Value: indicators[value].YData[i].Value, CName: indicators[value].CountryName}
					fmt.Println(r)
					res = append(res, r)
				}
			}
		}

	}
	sort.Sort(res)
	return res
}
