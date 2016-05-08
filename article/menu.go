package article

// MenuEntry represents a single entry (and children) within the menu
type MenuEntry struct {
	Title    string
	Article  *Article
	Children []*MenuEntry
}

// Menu represents the menu structure
type Menu struct {
	Entries []*MenuEntry
}
