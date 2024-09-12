package main

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)


const count_rows = 5;


func main() {
	file, err := os.Open("letters.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	var letters map[string][]interface{}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&letters)
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your text: ")

	var text string

	if scanner.Scan() {
		text = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < count_rows; i++ {
		for _, symbol := range text {
			s := string(symbol)
			if value, exists := letters[s]; exists {
				fmt.Printf("%s ", value[i])
			} else {
				fmt.Printf("Symbol %c does not exist\n")
			}
		}
		fmt.Println()
	}
}
