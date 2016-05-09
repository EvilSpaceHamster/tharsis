package sourcebook

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/evilspacehamster/tharsis/article"
)

type NativeSourceBook struct {
	RootFilePath string
}

func (b *NativeSourceBook) getAllArticles() ([]*article.Article, error) {

	var paths []string
  func processFile(path string, f os.FileInfo, err error) error {
    paths = append(paths,path)
  }
  filepath.Walk(b.RootFilePath, processFile)
	var articles []*article.Article

	for _, p := range paths {
		a, err := getArticleFromFilePath(p)
		if err == nil {
			articles = append(articles, a)
		}
	}

	return articles, nil
}

func getArticleFromFilePath(path string) (*article.Article, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var title string
	filename := strings.Split(filepath.Base(path), ".")
	title = strings.Replace(filename[0], "-", " ", 0)

	title = strings.Title(title)

	article := &Article{
		Path:       path,
		Source:     bytes,
		Title:      title,
		IsCompiled: false,
	}
	return article, nil
}
