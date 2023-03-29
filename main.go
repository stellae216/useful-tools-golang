package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"useful-tools-golang/application"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("请输入chatGPT token: ")
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
		reply, err := gptClient.SendMessagesWithContext(message)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Printf("[%d] reply: %s\n", i, reply)
	}
	fmt.Println("finish...")
}
