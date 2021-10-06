package gorms

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Forms struct {
	Document  goquery.Document
	Questions []Question
}
type Question struct {
	Title     string
	Selection goquery.Selection
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
	c.Questions = findQuestions(c.Document)
	return c
}

func findQuestions(d goquery.Document) []Question {
	var result []Question
	d.Find("div.office-form-question-content").Each(func(index int, s *goquery.Selection) {
		result = append(result, Question{
			Title:     s.Find("question-title-box").Text(),
			Selection: *s,
		})
	})
	return result
}
func GetQuestionType(s goquery.Selection) {

}
