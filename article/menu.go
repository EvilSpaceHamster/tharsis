package article

const (
	//MenuTypeStandard is a menu where the highest level is the main root menu.
	MenuTypeStandard = iota

	//MenuTypeGrouped is a menu where each top level entry is a separate root menu.
	//An Example usage of this is versioned document books where each top level
	//menu entry is a version.
	MenuTypeGrouped
)

// MenuEntry represents a single entry (and children) within the menu
type MenuEntry struct {
	Title    string       `json:"title"`
	URL      string       `json:"URL"`
	Children []*MenuEntry `json:"children"`
}

// Menu represents the menu structure
type Menu struct {
	Type    int
	Entries []*MenuEntry
}
