package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("app.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content := bufio.NewReader(file)
	for {
		line, err := content.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(line, "WARNING") {
			fmt.Println(line)
		}
	}
}
