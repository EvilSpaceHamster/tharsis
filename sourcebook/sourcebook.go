package sourcebook

import "github.com/evilspacehamster/tharsis/article"

// SourceBook represents an system for getting the article source from
type SourceBook interface {
	GetAllArticles() ([]*article.Article, error)
	GetMenu() (*article.Menu, error)
}
