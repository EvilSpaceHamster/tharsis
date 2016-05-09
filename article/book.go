package article

import (
	"path/filepath"

	"github.com/evilspacehamster/tharsis/config"
)

type Book struct {
	Menu       *Menu
	Articles   []*Article
	IsCompiled bool
}

func (b *Book) Compile(c *config.Config) {
	for _, a := range b.Articles {
		if !a.IsCompiled {
			ext := filepath.Ext(a.Path)
			compiler, err := c.GetCompilerForFileExtension(ext)
			if err != nil {
				continue
			}
			output, err := compiler.CompileContent(a.Source)
			if err != nil {
				continue
			}
			a.Content = output
			a.IsCompiled = true
		}

	}
	b.IsCompiled = true
}
