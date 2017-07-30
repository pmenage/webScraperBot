package scraper

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// ArticlesScraper gets the first articles of the page
func ArticlesScraper() ([]string, []string) {

	doc, err := goquery.NewDocument("http://www.aljazeera.com/")
	if err != nil {
		panic(err)
	}

	var msgs []string
	var images []string

	doc.Find(".top-section-lt").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h1").Text()
		link, _ := s.Find("a").Attr("href")
		img, _ := s.Find("img").Attr("src")
		description := s.Find(".top-sec-desc").Text()
		msgText := fmt.Sprintf("<b>%s</b>\nhttp://www.aljazeera.com%s\n%s\n\n", title, link, description)
		msgs = append(msgs, msgText)
		imgText := fmt.Sprintf("http://www.aljazeera.com%s", img)
		images = append(images, imgText)
	})

	doc.Find(".top-section-rt-s1, .top-section-rt-s2").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h2").Text()
		link, _ := s.Find("a").Attr("href")
		img, _ := s.Find("img").Attr("src")
		msgText := fmt.Sprintf("<b>%s</b>\nhttp://www.aljazeera.com%s\n\n", title, link)
		msgs = append(msgs, msgText)
		imgText := fmt.Sprintf("http://www.aljazeera.com%s", img)
		images = append(images, imgText)
	})

	doc.Find(".media").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h4").Text()
		link, _ := s.Find("a").Attr("href")
		img, _ := s.Find("img").Attr("src")
		msgText := fmt.Sprintf("<b>%s</b>\nhttp://www.aljazeera.com%s\n\n", title, link)
		msgs = append(msgs, msgText)
		imgText := fmt.Sprintf("http://www.aljazeera.com%s", img)
		images = append(images, imgText)
	})

	return msgs, images

}
