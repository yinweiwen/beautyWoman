package wallpaper

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func ClearExtDirs() {
	dirs, err := ioutil.ReadDir("download")
	if err != nil {
		return
	}
	var now = time.Now()
	for _, dir := range dirs {
		if dir.IsDir() {
			ph := path.Join("download", dir.Name())
			files, err := ioutil.ReadDir(ph)
			if err != nil || len(files) > 0 {
				continue
			}
			dn := dir.Name()
			idx := strings.Index(dn, "_")
			if idx > 0 {
				if t, err := time.Parse("2006_1_2", dn[idx+1:]); err == nil {
					if now.Sub(t).Hours() > 48 {
						os.RemoveAll(ph)
					}
				}
			}
		}
	}
}

// copy of CrawWallPaper
// use one colly instance
func CrawWallPaperInOneColly() {
	ClearExtDirs()

	today := time.Now()
	outputDir := fmt.Sprintf("download/wallpaper_%d_%d_%d/", today.Year(), today.Month(), today.Day())
	if _, oe := os.Stat(outputDir); os.IsNotExist(oe) {
		os.MkdirAll(outputDir, 0777)
	}

	c := colly.NewCollector(
		colly.AllowedDomains("w.wallhaven.cc", "wallhaven.cc"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36"),
	)
	c.SetRequestTimeout(120 * time.Second)
	c.SetProxy("http://127.0.0.1:7890")

	c2 := c.Clone()

	c.Limit(&colly.LimitRule{
		DomainRegexp: `wallhaven\.cc`,
		Parallelism:  2,
		Delay:        5 * time.Second,
	})

	c2.Async = true
	c2.Limit(&colly.LimitRule{
		DomainRegexp: `wallhaven\.cc`,
		Parallelism:  5,
		Delay:        5 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Set("if-modified-since", "Fri, 01 Apr 2022 13:56:34 GMT")
		r.Headers.Set("if-none-match", "\"62470492-f06a5\"")
		r.Headers.Set("upgrade-insecure-requests", "1")
		r.Headers.Set("Cache-Control", "max-age=0")
		r.Headers.Set("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"99\", \"Google Chrome\";v=\"99\"")
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", "\"Windows\"")
		r.Headers.Set("sec-fetch-dest", "document")
		r.Headers.Set("sec-fetch-mode", "navigate")
		r.Headers.Set("sec-fetch-site", "none")
		r.Headers.Set("sec-fetch-user", "?1")
	})

	c.OnHTML("section.thumb-listing-page", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
			a := element.Attr("href")
			if a == "#top" {
				return
			}
			log.Println("down visit ", a)
			c.Visit(a)
		})
	})

	c.OnHTML("#wallpaper", func(e *colly.HTMLElement) {
		detail := e.Attr("src")
		log.Println("downloading ", detail)
		c2.Visit(detail)
	})

	c2.OnResponse(func(r *colly.Response) {
		if strings.Index(r.Headers.Get("Content-Type"), "image") > -1 {
			r.Save(outputDir + r.FileName())
			return
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", err, r.Request.URL, string(r.Body))
	})

	index := 1
	for i := index; i < index+10; i++ {
		a := fmt.Sprintf(url, i)
		log.Println("visit ", a)
		c.Visit(a)
	}
	c.Wait()

	log.Println("done...")
	time.Sleep(5 * time.Minute)
	log.Println("done!!!")
}
