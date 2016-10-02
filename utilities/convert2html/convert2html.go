package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/humrs/ouchapp"
)

var htmlstringbegin = `<!DOCTYPE html>

<html lang="en" xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta charset="UTF-8" />
    <!--[if IE]><meta http-equiv="X-UA-Compatible" content="IE=edge"><![endif]-->
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="generator" content="Asciidoctor 1.5.2">
</head>
<body>
    <div id="content">`

var htmlstringend = `
    </div>
</body>
</html>`

func main() {
	fmt.Println(htmlstringbegin)
	fmt.Println(htmlstringend)

	afilename := flag.String("afile", "", "Input Article Filename")
	pfilename := flag.String("pfile", "", "Input Page Filename")
	//sfilename := flag.String("sfile", "", "Input Stylesheet Filename")
	//tfilename := flag.String("tfile", "", "Input template filename")
	flag.Parse()

	fp, err := os.Open(*afilename)
	if err != nil {
		log.Fatalf("Unable to open input filename\n")
	}

	jsdecoder := json.NewDecoder(fp)

	var articles []ouchapp.Article

	err = jsdecoder.Decode(&articles)
	if err != nil {
		log.Fatal(err)
	}

	fp.Close()

	fp, err = os.Open(*pfilename)
	if err != nil {
		log.Fatalf("unable to open pfile")
	}

	jsdecoder = json.NewDecoder(fp)

	var pages []ouchapp.Page

	err = jsdecoder.Decode(&pages)
	if err != nil {
		log.Fatal(err)
	}

	fp.Close()

	fmt.Printf("Pages: %v\n", len(pages))

	fmt.Printf("Count: %v\n", len(articles))

	articlesDict := make(map[string]ouchapp.Article)
	for _, v := range articles {
		articlesDict[v.EmbeddedID] = v
	}

	pagesDict := make(map[string]ouchapp.Page)
	for _, v := range pages {
		pagesDict[v.EmbeddedID] = v
	}
	fmt.Printf("Pages Dict: %v\n", len(pagesDict))
	fmt.Printf("Articles Dict: %v\n", len(articlesDict))

	for _, v := range pages {
		if strings.Contains(v.EmbeddedID, "weblink_") {
			continue
		}

		var items []string
		items = append(items, htmlstringbegin)

		var middle string
		if v.ArticleID == "" {
			middle = v.LinkTitle
		} else {
			article, ok := articlesDict[v.ArticleID]
			if !ok {
				fmt.Printf("Problem with finding weblink article: %v, %v\n", v.EmbeddedID, v.ArticleID)
			}
			middle = article.HTMLContent
		}
		items = append(items, middle)

		var ids []string
		err := json.Unmarshal([]byte(v.LinkIDs), &ids)
		if err != nil {
			fmt.Printf("Error: %v, %v\n", v.LinkIDs, v.EmbeddedID)
		}

		for _, w := range ids {

			var link string

			a, ok := articlesDict[w]
			if !ok {
				fmt.Printf("Problem: %v: %v\n", v.EmbeddedID, w)
			}

			link = fmt.Sprintf(`<div class="clickableLink"><p class="clickableLink"><a class="clickableLink" href="%v">%v</a></p></div>`, a.EmbeddedID, a.Title)

			items = append(items, link)
		}

		items = append(items, htmlstringend)

		htmlcontent := strings.Join(items, "")

		filename := fmt.Sprintf("%v.html", v.EmbeddedID)
		outfp, err := os.Create(filename)
		if err != nil {
			fmt.Println("Problem opening the html file")
		}
		outfp.WriteString(htmlcontent)
		outfp.Close()

	}

	for _, v := range articles {
		if strings.Contains(v.EmbeddedID, "weblink") {
			doWebLink(v)
		}
	}
	fmt.Println("Pages")

	for _, v := range pages {
		fmt.Println(v.EmbeddedID)
	}

}

func doWebLink(article ouchapp.Article) {
	filename := fmt.Sprintf("%v.json", article.EmbeddedID)
	outfp, err := os.Create(filename)
	if err != nil {
		fmt.Println("Problem opening weblink file")
	}

	var payload []byte
	payload, err = json.Marshal(article)
	outfp.WriteString(string(payload))
	outfp.Close()

}
