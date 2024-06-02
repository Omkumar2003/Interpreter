package main

import (
	"fmt"
	"io/ioutil"

	// "io"
	"os"
	"strings"

	"github.com/om/interpreter/repl"
)

func main() {
	if len(os.Args) == 1 || (len(os.Args) == 2 && os.Args[1] == "repl") {
		// fmt.Println("hdjvdh")
		fmt.Print("░█████╗░███╗░░░███╗\n██╔══██╗████╗░████║\n██║░░██║██╔████╔██║\n██║░░██║██║╚██╔╝██║\n╚█████╔╝██║░╚═╝░██║\n░╚════╝░╚═╝░░░░░╚═╝\n")

		repl.Start(os.Stdin, os.Stdout, true)

	} else if len(os.Args) == 2 {
		filename := os.Args[1]
		if !strings.HasSuffix(filename, ".om") {
			fmt.Println("Error: Input file must have a .om extension")
			os.Exit(1)
		}

		content, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			os.Exit(1)
		}

		repl.Start(strings.NewReader(string(content)), os.Stdout, false)
	} else {
		fmt.Println("Usage: go run main.go [repl] [<filename.om>]")
		os.Exit(1)
	}
}
