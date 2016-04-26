package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var paths []string

var root string

func main() {
	fmt.Println("Tharsis Documentation Service")

	src := flag.String("src", "", "Source directory for markdown documentation")
	// out := flag.String("out", "", "HTML Output directory")
	flag.Parse()

	if *src == "" {
		fmt.Println("src directory is required")
		return
	}

	filepath.Walk(*src, visit)
}

func visit(path string, f os.FileInfo, err error) error {

	p := strings.TrimPrefix(path, root)
	if strings.HasSuffix(p, ".md") {
		fmt.Printf("Found: %s\n", p)
		paths = append(paths, p)
	}

	return nil
}
