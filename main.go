package main

import (
	"context"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/chromedp/chromedp"
	"github.com/graniticio/inifile"
)

var (
	dir string
)

func main() {
	flag.StringVar(&dir, "dir", "", "Absolute path for target directory")
	flag.Parse()
	if len(dir) < 1 {
		log.Fatal("No --dir is given")
	}

	urlFiles := Scan(dir, ".url")
	fmt.Printf("There are %d url files\n", len(urlFiles))
	for _, s := range urlFiles {
		ic, err := inifile.NewIniConfigFromPath(s)
		errExit(err)
		url, err := ic.Value("InternetShortcut", "URL")
		errExit(err)
		fmt.Printf("- [%s](%s)\n", getTitle(url), url)
	}
}

func errExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Scan(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func getTitle(urlstr string) string {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var title string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(urlstr),
		chromedp.Title(&title),
	); err != nil {
		log.Fatal(err)
	}
	return title
}
