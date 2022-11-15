package sitemap

import (
	"os"
	"path/filepath"
	"testing"
)

const tmpDir = `testdata`

func init() {
	t, err := filepath.Abs(tmpDir)
	if err != nil {
		panic(err)
	}
	err = os.RemoveAll(t)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(t, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TestGenerator(t *testing.T) {
	o := Options{
		MaxURLs:  2,
		Dir:      tmpDir,
		Filename: "articles",
		BaseURL:  "https://example.com/",
	}

	g := New(o)
	g.Open()
	g.Add(Url{Loc: o.BaseURL, Priority: "1.0", LastMod: "12121", ChangeFreq: ChangeFreqDaily, Language: "zh,en,ko,ja"})
	g.Close()

	o.Filename = "b"
	g = New(o)
	g.Open()
	g.Add(Url{Loc: "test1"})
	g.Add(Url{Loc: "test2"})
	g.Add(Url{Loc: "test3"})
	g.Add(Url{Loc: "test4"})
	g.Add(Url{Loc: "test5"})
	g.Close()

	/*	files, _ := os.ReadDir(g.opt.Dir)
		files[0].Name()
		files[1].Name()
		files[2].Name()
		files[3].Name()
		files[4].Name()*/
}

func TestParamChecks(t *testing.T) {
	/*	var (
			g   *Generator
			err error
		)
		opt := Options{
			MaxFileSize: 2,
			MaxURLs:     -1,
			BaseURL:     "/",
			Dir:         tmpDir,
		}
		g = New(opt)
		err = g.Open()
		require.Error(t, err)

		opt.MaxFileSize = len(header) + len(footer) + 10
		g = New(opt)
		err = g.Open()
		require.Error(t, err)

		opt.MaxURLs = 2
		g = New(opt)
		err = g.Open()
		 err)*/
}

func TestInternals(t *testing.T) {
	/*	var (
			g   *Generator
			err error
		)
		opt := Options{
			MaxFileSize: len(header) + len(footer) + 10,
			MaxURLs:     2,
			BaseURL:     "/",
			Dir:         tmpDir,
		}
		g = New(opt)
		err = g.Open()
		 err)
		require.True(t, g.canFit(10))
		require.False(t, g.canFit(11))

		n1 := g.formatURLNode(URL{Loc: "test1"})
		Url `<url><loc>test1</loc></url>`, n1)

		n2 := g.formatURLNode(URL{Loc: "test2", ChangeFreq: ChangeFreqDaily})
		Url `<url><loc>test2</loc><changefreq>daily</changefreq></url>`, n2)*/
}
