package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/evilspacehamster/tharsis/article"
)

var paths []string

var root string

var docs map[string]string

var docMutex *sync.RWMutex

func main() {
	fmt.Println("Tharsis Documentation Service")

	docMutex = &sync.RWMutex{}
	src := flag.String("src", "", "Source directory for markdown documentation")
	// out := flag.String("out", "", "HTML Output directory")
	flag.Parse()

	if *src == "" {
		fmt.Println("src directory is required")
		return
	}

	filepath.Walk(*src, processFile)
}

func processFile(path string, f os.FileInfo, err error) error {

	p := strings.TrimPrefix(path, root)
	if strings.HasSuffix(p, ".md") {
		fmt.Printf("Found: %s\n", p)
		paths = append(paths, p)
		go parseMarkdownFile(path, p)
	}

	return nil
}

func parseMarkdownFile(filename, relativeName string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading file '%s': %s\n", filename, err)
	}
	docMutex.Lock()
	defer docMutex.Unlock()
	docs[relativeName] = article.RenderMarkdownArticle(string(bytes))
}
