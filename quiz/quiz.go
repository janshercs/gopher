package quiz

import (
	"encoding/csv"
	"fmt"
	"os"
)

type InvalidInputError struct {
	message string
}

func (e InvalidInputError) Error() string {
	return e.message
}

func NewInvalidInputError(msg string) InvalidInputError {
	return InvalidInputError{msg}
}

func OpenAndReadCSVFile(f string) ([][]string, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)
	return ReadCSVFile(csvReader)
}

func ReadCSVFile(r *csv.Reader) ([][]string, error) {
	content, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return content, nil
}

func GetQuestionsFromStrings(input [][]string) ([]Question, error) {
	var questions = make([]Question, 0, len(input))
	var err error
	for i, QA := range input {
		if len(QA[0]) == 0 || len(QA[1]) == 0 {
			err = NewInvalidInputError(fmt.Sprintf("skipping entry #%d in file due to empty input", i))
			continue
		}
		questions = append(questions, Question{QA[0], QA[1]})
	}
	return questions, err
}

func Start(questions []Question) {
	score := 0
	for i, question := range questions {
		fmt.Printf("Question #%d: %s = ", i+1, question.Q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == question.A {
			score++
		}
	}
	fmt.Printf("You got %d answers correct \n", score)
}

type Question struct {
	Q, A string
}
