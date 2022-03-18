package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	correct := ""
	if string(text[0]) == "9" {
		correct = "251" + text

	} else if strings.HasPrefix(text, "+251") {
		correct = "251" + text[4:]
	} else if strings.HasPrefix(text, "09") {
		correct = "251" + text[1:]
	} else if strings.HasPrefix(text, "251") {
		correct = "251" + text[3:]
	} else {
		fmt.Println("incorrect input")
	}

	fmt.Println(correct)

}
