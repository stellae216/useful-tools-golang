package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/stellae216/useful-tools-golang/application"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("please input chatGPT token: ")
	scanner.Scan()
	token := strings.Trim(scanner.Text(), " ")
	gptClient := application.GenClient(token)
	fmt.Println("========================================")
	fmt.Println("Please start chatting, to exit please enter:【exit】!")
	for i := 1; true; i = i + 1 {
		fmt.Println("----------------------------------------")
		fmt.Printf("[%d] message: ", i)
		scanner.Scan()
		message := strings.Trim(scanner.Text(), " ")
		if message == "exit" {
			break
		}
		fmt.Printf("[%d] reply -> \n", i)
		_, err := gptClient.SendMessageStream(message)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
	}
	fmt.Println("finish...")
}
