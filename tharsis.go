package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var paths []string

var root string

func main() {
	fmt.Println("Tharsis Documentation Service")

	root = "/Users/ali/mockumentation/src"
	filepath.Walk(root, visit)
}

func visit(path string, f os.FileInfo, err error) error {

	p := strings.TrimPrefix(path, root)
	if strings.HasSuffix(p, ".md") {
		fmt.Printf("Found: %s\n", p)
		paths = append(paths, p)
	}

	return nil
}
