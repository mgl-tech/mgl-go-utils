package mgl

import (
	"fmt"
	"github.com/mgl-tech/mgl-go-utils/sitemap"
)

func Test() {
	fmt.Print("--------------")

	o := sitemap.Options{
		MaxURLs:  2,
		Dir:      "t",
		Filename: "articles",
		BaseURL:  "https://example.com/",
	}

	g := sitemap.New(o)

	g.Open()

	g.Add(sitemap.URL{
		Loc: "ddddd",
	})

	g.Close()

}
