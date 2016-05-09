package article

// Article represents an article and it's metadata
type Article struct {
	ID         []byte
	Metadata   MetaData
	Title      string
	Path       string
	URL        string
	Source     []byte
	Content    []byte
	IsCompiled bool
}
