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

// sort sorts question into index 0 and answer into index 1
func sort(records [][]string) ([]Quiz, error) {
	var data []Quiz

	for _, v := range records {

		if len(v) != 2 {
			return nil, fmt.Errorf("out of index range")

		}
		q := Quiz{
			question: strings.TrimSpace(v[0]),
			answer:   strings.TrimSpace(v[1]),
		}
		data = append(data, q)
	}

	return data, nil
}

// askQuestions displays the questions at index 0
func askQuestion(q []Quiz) {

	for _, v := range q {

		fmt.Printf("question: %s\n", v.question)

		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSuffix(input, "\n")

		//send the answer to an answer
		answerCh := make(chan string)

		go readInputAndCompare(answerCh, q)
		result(answerCh)

		os.Exit(1)

	}
}

// getInputAndCompare reads user input
// checks if user answer matches the correct answer
func readInputAndCompare(answerCh chan<- string, quiz []Quiz) {
	correct := 0

	for _, v := range quiz {
		fmt.Printf("question: %s\n", v.question)

		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("error reading user input")
		}
		answer = strings.TrimSpace(answer)

		if answer == v.answer {
			correct++
		}
	}

	quizScore := fmt.Sprintf("You scored %d out of %d\n", correct, len(quiz))

	answerCh <- quizScore

}

// result recieves the quiz score
func result(answerCh <-chan string) {
	fmt.Println(<-answerCh)
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "questions,answers")

	flag.Parse()

	f, err := os.Open(*csvFileName)
	if err != nil {

		log.Fatal(err)

	}

	//read csv file
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//sort records into indexes
	s, err := sort(records)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("Answer the following question\n")
	fmt.Println("See the RESULT at the end of your quiz\n")
	askQuestion(s)

}
