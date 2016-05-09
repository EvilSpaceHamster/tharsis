package compiler

import (
	"errors"

	"github.com/evilspacehamster/tharsis/article"
	"github.com/russross/blackfriday"
)

type MarkdownCompiler struct{}

func (c *MarkdownCompiler) CompileContent(source []byte) ([]byte, error) {
	content := blackfriday.MarkdownCommon(source)
	return content, nil
}

func (c *MarkdownCompiler) ExtractMetadata(source []byte) (article.MetaData, error) {
	return nil, errors.New("Not implemented")
}

func NewMarkdownCompiler() *MarkdownCompiler {
	return &MarkdownCompiler{}
}
