package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Forms struct {
	Document  goquery.Document
	Questions []goquery.Selection
}

func NewForms(targetUrl string) *Forms {
	c := new(Forms)
	resp, err := http.Get(targetUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	c.Document = *doc
	c.Questions = getQuestionSelections(c.Document)
	return c
}

func getQuestionSelections(d goquery.Document) []goquery.Selection {
	var result []goquery.Selection
	d.Find("text-format-content").Each(func(index int, s *goquery.Selection) {
		result = append(result, *s)
	})
	return result
}
