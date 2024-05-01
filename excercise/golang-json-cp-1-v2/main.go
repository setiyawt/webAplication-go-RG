package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Study struct {
	StudyName   string `json:"study_name"`
	StudyCredit int    `json:"study_credit"`
	Grade       string `json:"grade"`
}

type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	// TODO: answer here
	response, err := ioutil.ReadFile(filename)
	if err != nil {
		return Report{}, err
	}

	var report Report
	err = json.Unmarshal(response, &report)
	if err != nil {
		return Report{}, err
	}

	if report.Studies == nil {
		report.Studies = make([]Study, 0)
	}

	return report, nil
}

func GradePoint(report Report) float64 {
	gradeMap := map[string]float64{
		"A":  4.0,
		"AB": 3.5,
		"B":  3.0,
		"BC": 2.5,
		"C":  2.0,
		"CD": 1.5,
		"D":  1.0,
		"DE": 0.5,
		"E":  0.0,
	}

	var totalScore float64
	var totalCredits int

	for _, study := range report.Studies {
		score := gradeMap[study.Grade] * float64(study.StudyCredit)
		totalScore += score
		totalCredits += study.StudyCredit
	}

	if totalCredits == 0 {
		return 0.0
	}

	return totalScore / float64(totalCredits) // TODO: replace this
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
