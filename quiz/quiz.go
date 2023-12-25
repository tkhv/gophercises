package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func errCheck (err error) {
	if err != nil {
		panic(err)
	}
}

func main(){
	file := flag.String("file", "problems.csv", "Filename of the csv containing the quiz.")
	flag.Parse()
	
	quizFile, err := os.Open(*file)
	errCheck(err)
	reader := csv.NewReader(quizFile)
	
	var answer int
	counter, correct := 0, 0

	for {
		question, err := reader.Read()
		if err == io.EOF {
			break
		} else {
			errCheck(err)
		}

		ans, err := strconv.Atoi(question[1])
		errCheck(err)

		fmt.Println(question[0])
		fmt.Scan(&answer)
		counter++
		if answer == ans {
			correct++
		}
	}
	
	fmt.Printf("You got %d/%d correct.", correct, counter)
}