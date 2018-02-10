//adding concurrency is to speed up the loading
//so what is basically happening is that we send a request and then wait for it to respond 
//go routines provide concurrency
//go routine is like a light weight thread

package main
// := for assigning necessary for initialization, rest all places can work with =
// var grades ... can be written as grades:= make(.....) if inside a function

import (
	"fmt"
	"net/http"
	"html/template"
	"encoding/xml"
	"io/ioutil"
)

type NewsMap struct {
	Keyword string
	Location string
}

type NewsAggPage struct {
    Title string
    News map[string]NewsMap
}

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s Sitemapindex
	var n News
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	news_map := make(map[string]NewsMap)
	//string_body := string(bytes)
	//fmt.Println(string_body)
	//resp.Body.Close()

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)					//this loop is really slow because its going to the site getting sitemap going to next and next

		for idx, _ := range n.Keywords {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}	//passing the news_map to News variable
    t, _ := template.ParseFiles("newsaggtemplate.html")					//giving the template
    t.Execute(w, p)														//executing the function to display
    resp.Body.Close()				//check this out
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil) 
}