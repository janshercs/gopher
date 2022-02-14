package main

import (
	"flag"
	"log"
	"quiz"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "csv file containing the problems and answers")
	flag.Parse()
	lines, err := quiz.OpenAndReadCSVFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	questions, err := quiz.GetQuestionsFromStrings(lines)
	if err != nil {
		log.Fatal(err)
	}
	quiz.Start(questions)
}
