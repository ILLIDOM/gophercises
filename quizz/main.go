package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

type quizzEntry struct {
	question string
	answer   string
}

func readCSVFile(fileName string) *[]quizzEntry {
	var entryList []quizzEntry

	//open file
	f, err := os.Open("questions.csv")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
	}
	// defer = execute when surrounding function returns
	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		entry, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading entry: %v", err)
		}
		entryList = append(entryList, quizzEntry{question: entry[0], answer: entry[1]})
	}
	return &entryList // pointer to list so no copy needed
}

func main() {
	entryList := readCSVFile("questions.csv")
	timeLimit := flag.Int("time", 10, "time limit to complete quizz, default 10 seconds")
	flag.Parse()
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct_answers := 0

	fmt.Print("Press Enter to start the quizz...\n")
	fmt.Scanln()

	// label for a loop to break anywhere
quizzloop:
	for index, entry := range *entryList {
		fmt.Printf("Problem %v: %v \n", index+1, entry.question)
		answerChannel := make(chan string)

		// go routine which gets directly executed and reads from stdin; writes result into answerChannel
		// go func() {
		// 	var answer string
		// 	fmt.Scanf("%s\n", &answer)
		// 	answerChannel <- answer
		// }()
		go readFromStdin(answerChannel)

		select {
		// try reading from time channel (duration is only written once - when duration is over)
		case <-timer.C:
			fmt.Println("Time Limit Exceeded")
			break quizzloop
		// try reading from answer channel (if result is present check if it is correct)
		case answer := <-answerChannel:
			if answer == entry.answer {
				correct_answers++
			}
		}
	}

	fmt.Printf("You got %v from %v correct answers", correct_answers, len(*entryList))
}

// function get passes a write only channel (because channel are created using make they are always passed as reference)
func readFromStdin(channel chan<- string) {
	var answer string
	fmt.Scanf("%s\n", &answer)
	channel <- answer
}
