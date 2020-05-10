package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gobuffalo/packr"
)

func list() {
	box := packr.NewBox("./templates")
	templates := box.List()

	for _, t := range templates {
		fmt.Println(strings.Split(t, ".")[0])
	}
}

func copyIgnore(language string) {
	box := packr.NewBox("./templates")
	template, err := box.Open(language + ".gitignore")
	if err != nil {
		fmt.Fprintf(os.Stderr, "no template found for: %s\n", language)
	}

	to, err := os.OpenFile(".gitignore", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to write .gitignore to current directory")
		os.Exit(1)
	}
	defer to.Close()

	_, err = io.Copy(to, template)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to write contents to .gitignore")
		os.Exit(1)
	}
}

func main() {
	shouldList := flag.Bool("l", false, "list available gitignores")
	language := flag.String("g", "", "language to generate gitignore for")
	flag.Parse()

	if *shouldList {
		list()
		os.Exit(0)
	}

	if *language == "" {
		fmt.Fprintln(os.Stderr, "language is required, use the -h flag for help")
		os.Exit(1)
	}

	copyIgnore(*language)
}
