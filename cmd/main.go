package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/vanc0uv3r/go-concurrency/cmd/lexer"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	lex := lexer.NewLex()
	for true {
		line, _,err := reader.ReadLine()
		if err != nil {
			fmt.Println("Occured error with reading stdin: ", err.Error())
		}
		
		lex.Analyze(line)
	}
	
	}

