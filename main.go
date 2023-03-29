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
	fmt.Print("please input chatGPT token: ")
	scanner.Scan()
	token := strings.Trim(scanner.Text(), " ")
	var gptClient = application.GenClientWithProxy(token, application.ProxyGPT{
		Protocol: "http", Addr: "127.0.0.1", Port: "7890",
	})
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
		fmt.Printf("[%d] reply -> \n%s\n", i, reply)
	}
	fmt.Println("finish...")
}
