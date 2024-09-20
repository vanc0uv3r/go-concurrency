package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/vanc0uv3r/go-concurrency/cmd/lexer"
	"github.com/vanc0uv3r/go-concurrency/cmd/engine"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	lex := lexer.NewLex()
	engine := engine.NewEngine()
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Occured error with reading stdin: ", err.Error())
		}

		lex.Analyze(line)
		engine.SetLexemes(lex.GetLexemes())
		res, err := engine.Execute()

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("result of %s operation: %s\n", engine.GetCommandName(), res)
		}
		lex.ClearLexer()
	}

}
