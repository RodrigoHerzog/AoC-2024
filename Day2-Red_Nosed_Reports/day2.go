package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports_list := ReadInput("./input-day2-red_nosed_reports.txt")

	fmt.Printf("Safe reports - No dampening: %v\n", getSafeReports(reports_list, false))
	fmt.Printf("Safe reports - With dampening: %v\n", getSafeReports(reports_list, true))
}

func getSafeReports(reports_list [][]int, damper_toggler bool) int {
	safe_reports := 0
	for _, report := range reports_list {

		report_is_safe := CheckLevelListSafety(report)

		if !report_is_safe {
			if damper_toggler {
				for idx := 0; idx < len(report); idx++ {
					report_damped := []int{}
					report_damped = append(report_damped, report[:idx]...)
					report_damped = append(report_damped, report[idx+1:]...)
					report_is_safe = CheckLevelListSafety(report_damped)
					if report_is_safe {
						break
					}
				}
			}
		}
		if report_is_safe {
			safe_reports++
		}
	}

	return safe_reports
}

func CheckLevelListSafety(report []int) bool {
	latest_increment := 0
	for idx := range report {

		if idx == len(report)-1 {
			continue
		}

		level_change := report[idx] - report[idx+1]
		if level_change == 0 ||
			level_change < -3 ||
			level_change > 3 ||
			(latest_increment == -1 && level_change > 0) ||
			(latest_increment == 1 && level_change < 0) {
			return false
		}

		if level_change < 0 {
			latest_increment = -1
		} else if level_change > 0 {
			latest_increment = 1
		}
	}
	return true
}

func ReadInput(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var list_input [][]int
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		list_report_levels := strings.Fields(string(line))
		list_report_int, err := ConvStringsToIntegers(list_report_levels)
		if err != nil {
			log.Fatal(err)
		}
		list_input = append(list_input, list_report_int)
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return list_input
}

func ConvStringsToIntegers(list_string []string) ([]int, error) {
	list_int := make([]int, len(list_string))
	for i, str_to_conv := range list_string {
		int_converted, err := strconv.Atoi(str_to_conv)
		if err != nil {
			return nil, err
		}
		list_int[i] = int_converted
	}
	return list_int, nil
}
