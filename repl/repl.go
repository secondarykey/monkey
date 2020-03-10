package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey/lexer"
	"monkey/token"
)

const Prompt = ">> "

func Start(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)

	for {
		fmt.Printf(Prompt)

		if !scanner.Scan() {
			return nil
		}

		line := scanner.Text()
		lx := lexer.New(line)

		for tok := lx.NextToken(); tok.Type != token.EOF; tok = lx.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

	}
	return nil
}
