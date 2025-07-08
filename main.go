package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command (add/exit): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			AddBook(input)
		}
	}
}

func AddBook(isbn string) {
	//init input
	fmt.Println("Adding book with ISBN:", isbn)

	//send isbn api req

	//process response

	//add book to excel
}