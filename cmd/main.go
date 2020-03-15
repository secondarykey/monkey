package main

import (
	"fmt"
	"os"
	"os/user"

	"monkey/repl"
)

func main() {

	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf(`Hello %s!`+"\n", u.Username)
	fmt.Println(`  This interpreter was created by "Writing An Interpreter In Go".`)
	fmt.Println(`    https://interpreterbook.com/`)
	fmt.Println(`    https://www.oreilly.co.jp/books/9784873118222/`)
	fmt.Println(`  Thx Thorsten Ball and Hirotsugu Shitara.`)

	repl.Start(os.Stdin, os.Stdout)
}
