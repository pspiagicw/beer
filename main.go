package main

import (
	"fmt"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/osy/lexer"
	"github.com/pspiagicw/osy/token"
	"github.com/pspiagicw/regolith"
)

// import "fmt"

func main() {
	rg, err := regolith.New(&regolith.Config{})
	defer rg.Close()

	if err != nil {
		goreland.LogFatal("Error creating regolith instance: %v", err)
	}

	// fmt.Println("Hello, World")

	for {
		input, err := rg.Input()

		if err != nil {
			goreland.LogFatal("Error reading input: %v", err)
		}

		if len(input) == 0 {
			continue
		}

		l := lexer.New(input)

		for {
			tok := l.Next()
			if tok.Type == token.EOF {
				break
			}
			fmt.Println(tok)
		}
	}
}
