package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type NewsAggPage struct {
    Title string
    News string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}		//giving the some hardcoded data
    t, _ := template.ParseFiles("basictemplating.html")							//giving the template
    t.Execute(w, p)																//executing the function to display
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil) 
}