package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	word_search_text := ReadInput("./input-day4-ceres_search.txt")
	fmt.Printf("Word XMAS occurs: %v\n", GetWordXMASCount(word_search_text))
	fmt.Printf("Word X-MAS occurs: %v\n", GetWordCrossMASCount(word_search_text))
}

func GetWordCrossMASCount(word_search_text [][]string) int {
	word_occurrence := 0

	text_rows := len(word_search_text)
	text_cols := len(word_search_text[0])

	for row := 0; row < text_rows; row++ {
		for col := 0; col < text_cols; col++ {
			if word_search_text[row][col] == "A" {
				if row+1 > text_rows-1 || row-1 < 0 || col+1 > text_cols-1 || col-1 < 0 {
					continue
				}
				if word_search_text[row-1][col-1] == "M" && word_search_text[row+1][col+1] == "S" ||
					word_search_text[row-1][col-1] == "S" && word_search_text[row+1][col+1] == "M" {
					if word_search_text[row-1][col+1] == "M" && word_search_text[row+1][col-1] == "S" ||
						word_search_text[row-1][col+1] == "S" && word_search_text[row+1][col-1] == "M" {
						word_occurrence++
					}
				}
			}
		}
	}

	return word_occurrence
}

func GetWordXMASCount(word_search_text [][]string) int {
	word_occurrence := 0

	text_rows := len(word_search_text)
	text_cols := len(word_search_text[0])
	diagonal_text_rows := text_rows + text_cols - 1

	horizontal_lines_to_search := GetHorizontalLines(text_rows, word_search_text)

	vertical_lines_to_search := GetVerticalLines(text_cols, text_rows, word_search_text)

	diagonal_left_lines_to_search := GetDiagonalLeftLines(diagonal_text_rows, text_rows, text_cols, word_search_text)

	diagonal_right_lines_to_search := GetDiagonalRightLines(horizontal_lines_to_search, diagonal_text_rows, text_rows, text_cols)

	regexp_instructions := "XMAS"
	word_occurrence += RegexpSearch(horizontal_lines_to_search, regexp_instructions)
	word_occurrence += RegexpSearch(vertical_lines_to_search, regexp_instructions)
	word_occurrence += RegexpSearch(diagonal_left_lines_to_search, regexp_instructions)
	word_occurrence += RegexpSearch(diagonal_right_lines_to_search, regexp_instructions)

	regexp_instructions = "SAMX"
	word_occurrence += RegexpSearch(horizontal_lines_to_search, regexp_instructions)
	word_occurrence += RegexpSearch(vertical_lines_to_search, regexp_instructions)
	word_occurrence += RegexpSearch(diagonal_left_lines_to_search, regexp_instructions)
	word_occurrence += RegexpSearch(diagonal_right_lines_to_search, regexp_instructions)

	return word_occurrence
}

func GetDiagonalRightLines(horizontal_lines_to_search []string, diagonal_text_rows int, text_rows int, text_cols int) []string {
	reverted_original_matrix := [][]string{}
	for _, line := range horizontal_lines_to_search {
		reverted_line := RevertString(line)
		char := strings.Split(reverted_line, "")
		reverted_original_matrix = append(reverted_original_matrix, char)
	}

	diagonal_right_word_search_text := make([][]string, diagonal_text_rows)
	for x := 0; x < text_rows; x++ {
		for y := 0; y < text_cols; y++ {
			diagonal_row := x + y
			diagonal_right_word_search_text[diagonal_row] = append(diagonal_right_word_search_text[diagonal_row], reverted_original_matrix[x][y])
		}
	}

	diagonal_right_lines_to_search := make([]string, diagonal_text_rows)
	for idx, line := range diagonal_right_word_search_text {
		diagonal_right_lines_to_search[idx] = strings.Join(line, "")
	}
	return diagonal_right_lines_to_search
}

func GetDiagonalLeftLines(diagonal_text_rows int, text_rows int, text_cols int, word_search_text [][]string) []string {
	diagonal_left_word_search_text := make([][]string, diagonal_text_rows)
	for x := 0; x < text_rows; x++ {
		for y := 0; y < text_cols; y++ {
			diagonal_row := x + y
			diagonal_left_word_search_text[diagonal_row] = append(diagonal_left_word_search_text[diagonal_row], word_search_text[x][y])
		}
	}

	diagonal_left_lines_to_search := make([]string, diagonal_text_rows)
	for idx, line := range diagonal_left_word_search_text {
		diagonal_left_lines_to_search[idx] = strings.Join(line, "")
	}
	return diagonal_left_lines_to_search
}

func GetVerticalLines(text_cols int, text_rows int, word_search_text [][]string) []string {
	vertical_word_search_text := make([][]string, text_cols)
	for i := range vertical_word_search_text {
		vertical_word_search_text[i] = make([]string, text_rows)
	}

	for x := 0; x < text_rows; x++ {
		for y := 0; y < text_cols; y++ {
			vertical_word_search_text[x][y] = word_search_text[y][x]
		}
	}

	vertical_lines_to_search := make([]string, text_cols)
	for idx, line := range vertical_word_search_text {
		vertical_lines_to_search[idx] = strings.Join(line, "")
	}
	return vertical_lines_to_search
}

func GetHorizontalLines(text_rows int, word_search_text [][]string) []string {
	horizontal_lines_to_search := make([]string, text_rows)
	for idx, line := range word_search_text {
		horizontal_lines_to_search[idx] = strings.Join(line, "")
	}
	return horizontal_lines_to_search
}

func RegexpSearch(lines_to_search []string, instructions string) int {
	total_occurrences := 0

	re := regexp.MustCompile(instructions)

	for _, line := range lines_to_search {
		occurrences := re.FindAllString(line, -1)
		total_occurrences += len(occurrences)
	}

	return total_occurrences
}

func RevertString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ReadInput(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input_list := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		char := strings.Split(line, "")
		input_list = append(input_list, char)
	}

	return input_list
}
