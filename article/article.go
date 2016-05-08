package article

import (
	"io/ioutil"

	"github.com/russross/blackfriday"
)

// Article represents an article and it's metadata
type Article struct {
	Title   string
	Path    string
	Content []byte
}

func NewArticleFromMarkdownFile(filepath string) (*Article, error) {

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	content := blackfriday.MarkdownCommon(bytes)
	article := &Article{
		Path:    filepath,
		Content: content,
	}
	return article, nil
}
