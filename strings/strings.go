package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func checkString(text string) {

	textRunes := []rune(text)
	start := unicode.IsUpper(textRunes[0])
	l := len(textRunes) - 2
	end := string(textRunes[l])
	if start && end == "." {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}

func palindrom(text string) {
	tr := []rune(text)
	l := len(tr) - 2

	for i := 0; i < l; i++ {
		if tr[i] == tr[l-i-1] {
		} else {
			fmt.Println("Нет")
			return
		}
	}
	fmt.Println("Палиндром")
}

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	checkString(text)
	palindrom(text)
}
