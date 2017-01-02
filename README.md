# hugo
Hürriyet API Wrapper For Golang

## Install

`$ go get github.com/melihmucuk/hugo`

## Example
```go
package main

import (
	"fmt"
	"github.com/melihmucuk/hugo"
)

func main() {
	h := hugo.NewHurriyetAPI("YOUR_API_KEY")

	query := &hugo.Query{Top: 10}
	articles, _ := h.ListArticles(query)
	fmt.Println("articles count => ", len(articles))

	article, _ := h.SingleArticle(articles[0].ID, nil)
	fmt.Println("article => ", article)

	searchQuery := &hugo.Query{S: "1", Skip: 50, Top: 50} // default: -1, -1 => new to old , 1 => old to new
	searchResults, _ := h.Search("reina", searchQuery)
	fmt.Println("results found => ", searchResults.Count)

	writers, _ := h.ListWriters(nil)
	fmt.Println("writers count => ", len(writers))

	writer, _ := h.SingleWriter(writers[2].ID, nil)
	fmt.Println("writer => ", writer)

	galleryQuery := &hugo.Query{Filter: "Path eq '/teknoloji/'", Top: 5}
	galleries, _ := h.ListPhotoGalleries(galleryQuery)
	fmt.Println("photo galleries about 'teknoloji' => ", len(galleries))

}
```

## TODO
- [X] Article
- [X] Search
- [X] Writer
- [X] Path
- [X] Page
- [X] News Photo Gallery
- [ ] Date
- [X] Column
- [ ] Test
