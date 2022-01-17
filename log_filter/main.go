package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//command line arguments
	path := flag.String("path", "app.log", "File path for log file location (default: app.log)")
	level := flag.String("level", "WARNING", "Severity of log message may be WARNING, TRACE, INFO (default: WARNING)")
	// go run . -help
	// go run . -level TRACE -path <filepath>

	flag.Parse()

	file, err := os.Open(*path)
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
		if strings.Contains(line, *level) {
			fmt.Println("*", line)
		}
	}
}
