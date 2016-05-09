package config

import (
	"fmt"

	"github.com/evilspacehamster/tharsis/compiler"
	"github.com/evilspacehamster/tharsis/sourcebook"
)

type Config struct {
	Sourcebook sourcebook.SourceBook

	compilerMap map[string]*compiler.Compiler
}

func (c *Config) GetCompilerForFileExtension(extension string) (*compiler.Compiler, error) {

	if cmp, exists := compilerMap[extension]; exists {
		return cmp, nil
	}
	switch extension {
	case "md":
		compilerMap[extension] = compiler.NewMarkdownCompiler
		return compilerMap[extension], nil
	}
	return nil, fmt.Errorf("unable to find suitable compiler for extension: %s", extension)
}
