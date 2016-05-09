package compiler

import "github.com/evilspacehamster/tharsis/article"

// Compiler takes a raw article source, and compiles to HTML
type Compiler interface {

	// CompileContent takes a raw article source and returns the HTML output
	CompileContent([]byte) ([]byte, error)

	// ExtractMetaData takes a raw article source and extracts any metadata encoded
	//within the source code
	ExtractMetaData([]byte) (article.MetaData, error)
}
