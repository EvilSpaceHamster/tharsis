package article

import "github.com/russross/blackfriday"

func RenderMarkdownArticle(source string) string {
	output := blackfriday.MarkdownCommon([]byte(source))

	return string(output)
}
