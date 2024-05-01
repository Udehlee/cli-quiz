package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Quiz struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "CSV file containing questions and answers")
	flag.Parse()

	// Open the CSV file
	f, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read CSV records
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Sort records into Quiz structs
	quizzes, err := sort(records)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("See the RESULT at the end of your quiz")

	// Ask questions and check answers
	result := checkAnswers(quizzes)
	fmt.Println(result)
}

// sort sorts CSV records into Quiz structs
func sort(records [][]string) ([]Quiz, error) {
	var data []Quiz

	for _, record := range records {
		if len(record) != 2 {
			return nil, fmt.Errorf("invalid record format: %v", record)
		}
		quiz := Quiz{
			question: strings.TrimSpace(record[0]),
			answer:   strings.TrimSpace(record[1]),
		}
		data = append(data, quiz)
	}

	return data, nil
}

// askQuestion asks a single question and returns user input
func askQuestion(question string) string {
	fmt.Printf("question: %s\n", question)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	return input
}

// checkAnswers compares user input and correct answers
// returns the quiz result
func checkAnswers(quiz []Quiz) string {
	correct := 0

	for _, quiz := range quiz {
		input := askQuestion(quiz.question)

		if input == quiz.answer {
			correct++
		}
	}

	result := fmt.Sprintf("You scored %d out of %d\n", correct, len(quiz))
	return result
}
