package main

import (
	"bufio"
	"embed"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

//go:embed data/wordlist.txt
var fsys embed.FS

func run(args []string, stdout io.Writer) error {
	fn := flag.String("l", "", "optional wordlist to use")
	state := flag.String("s", "", "state of the Wordle puzzle (A-Z=green,a-z=yellow,_=blank")
	deny := flag.String("d", "", "deny list of unavailable letters")

	flag.Parse()
	if *state == "" {
		flag.PrintDefaults()
		return fmt.Errorf("must specify -s flag")
	}
	var f fs.File
	var err error
	if *fn == "" {
		f, err = fsys.Open("data/wordlist.txt")
	} else {
		f, err = os.Open(*fn)
	}
	if err != nil {
		return fmt.Errorf("error opening wordlist: %w", err)
	}
	defer f.Close()

	words := make([]string, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}

	for _, w := range words {
		m, err := IsMatch(*state, w, *deny)
		if err != nil {
			return fmt.Errorf("ismatch: %v", err)
		}
		if m {
			fmt.Printf("%s\n", w)
		}
	}

	return nil
}

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		log.Fatalf("run: %v", err)
	}
}
