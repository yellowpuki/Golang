package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func phraseCounter(fileName string, phrase string) int {
	var count int
	count = 0

	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		log.Fatalf("error when opening file: %s\n", fileName)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()

		if word == phrase {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file ")
	}
	return count
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two parameters: file path and phrase!")
		return
	}

	filePath := os.Args[1]
	phrase := os.Args[2]

	count := phraseCounter(filePath, phrase)

	fmt.Println(count)

}
