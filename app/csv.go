package app

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func (app *GrammarTest) LoadQuestionsFromCSV(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	firstLine := true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV record:", err)
			continue
		}
		if firstLine {
			firstLine = false
			continue
		}
		id := atoi(record[0])
		question := record[1]
		options := []string{record[2], record[3], record[4], record[5]}
		answer := record[6]
		feedback := record[7]

		q := Question{ID: id, Question: question, Options: options, Answer: answer, Feedback: feedback}
		app.questions = append(app.questions, q)
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	return i
}
