package storage

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search/query"
	"github.com/satori/go.uuid"
	"log"
	"os"
	"os/user"
	"path"
)

type Storage struct {
	index bleve.Index
}

type Article struct {
	Id   string
	Name string
	Body string
}

func (self *Article) Type() string {
	return "article"
}

func (self Storage) Remember(text string) (err error) {
	article := Article{
		Id:   uuid.NewV4().String(),
		Body: text,
	}

	return self.index.Index(article.Id, article)
}

func (self Storage) Recall(needle string) (err error) {
	q := bleve.NewMatchQuery(needle)
	q.SetOperator(query.MatchQueryOperatorAnd)
	searchRequest := bleve.NewSearchRequest(q)

	searchResult, err := self.index.Search(searchRequest)
	if err != nil {
		log.Fatal("unable to search for text " + needle)
	}

	for _, match := range searchResult.Hits {
		log.Println(match.String())
	}

	return nil
}

func GetStorage() *Storage {
	workDir := getWorkDirPath()
	if err := os.MkdirAll(workDir, 0700); err != nil {
		log.Fatal("unable to make work dir at " + workDir)
	}

	indexPath := getIndexPath()

	index, err := bleve.Open(indexPath)
	if err == bleve.ErrorIndexPathDoesNotExist {
		index, err = createIndex(indexPath)
		if err != nil {
			log.Fatalf("unable to create index at %s: %s",
				indexPath, err)
		}
	} else if err != nil {
		log.Fatal("unable to open index at " + indexPath)
	}

	return &Storage{index: index}
}

func getIndexPath() string {
	return path.Join(getWorkDirPath(), "index.bleve")
}

func getWorkDirPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("unable to get current user")
	}

	return path.Join(usr.HomeDir, ".eden")
}

func createIndex(indexPath string) (bleve.Index, error) {
	bodyFieldMapping := bleve.NewTextFieldMapping()

	articleMapping := bleve.NewDocumentMapping()
	articleMapping.AddFieldMappingsAt("body", bodyFieldMapping)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.DefaultAnalyzer = "standard"
	indexMapping.AddDocumentMapping("article", articleMapping)

	return bleve.New(indexPath, indexMapping)
}
