//NewsAggregator Web App using Go

//Go has built-in concurrency which we are using to speedup the fetching time
//adding concurrency is to speed up the loading
//so what is basically happening is that we send a request and then wait for it to respond 
//go routines provide concurrency
//go routine is like a light weight thread
//we need to synchronize
//we can use channels, defer 
// := for assigning necessary for initialization, rest all places can work with =
// var grades ... can be written as grades:= make(.....) if inside a function

package main
import (
	"fmt"
	"net/http"
	"html/template"
	"encoding/xml"
	"io/ioutil"
	"sync"
)

//for synchronization
var wg sync.WaitGroup

//defining struct to store the keyword and location 
type NewsMap struct {
	Keyword string
	Location string
}

//this struct is to store title and a map to store news links
type NewsAggPage struct {
    Title string
    News map[string]NewsMap
}

//To access the loc tag inside sitemap tag
type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

//To access respective tags
type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

//indexHandler function for index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the test page, lets go to main page, add /agg to the link address</h1>")
}

//we were sending request to each xml site waiting to get response and then next, so that was slower
//We implement the below function to add concurrency by go routines  
//this function below is to add routines to the main loop
func newsRoutine (c chan News,Location string){

	defer wg.Done()
	var n News
	resp, _ := http.Get(Location)				//gets the response from that location from main sitemap
	bytes, _ := ioutil.ReadAll(resp.Body)		//read the response 
	xml.Unmarshal(bytes, &n)					//unmarshall or convert it into text format
	resp.Body.Close()

	c<- n 										//returning n to the channel
}

//main page under /agg path
func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s Sitemapindex
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)											//read the response 
	xml.Unmarshal(bytes, &s)														//unmarshall or convert it into text format
	news_map := make(map[string]NewsMap)
	resp.Body.Close()				//check this out
	//string_body := string(bytes)
	//fmt.Println(string_body)
	//resp.Body.Close()

	queue := make(chan News, 30)					//make a channel with buffer size of 30
	for _, Location := range s.Locations {
		wg.Add(1)									//Adding everytime a new routine is defined
		go newsRoutine(queue,Location)				//calling go routine

	}

	wg.Wait()										//waits for all routines to finish executing
	close(queue)									//closing the channel

	//the reason I removed for idx loop from the before loop is that we need to access bunch in form of channels
	for elem := range queue {
		for idx, _ := range elem.Keywords {
			news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
		}
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}	//passing the news_map to News variable
    t, _ := template.ParseFiles("newsaggtemplate.html")					//giving the template
    t.Execute(w, p)														//executing the function to display
    
}

//main function 
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil) 
}