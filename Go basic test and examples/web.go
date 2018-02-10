package main
// := for assigning necessary for initialization, rest all places can work with =
// var grades ... can be written as grades:= make(.....) if inside a function

import ("fmt"
"io/ioutil"
"net/http"
"encoding/xml")

/*We can do it this way or shorten it as below this set comment
//struct for getting sitemap tag
type Sitemapindex struct {
  Locations []Location `xml:"sitemap"`
}

//struct for getting loc tag
type Location struct {
  Loc string `xml:"loc"`
}

//This function makes that links as strings it removes {} 
func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}
*/
type Sitemapindex struct {
  Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

func main() {
	//resp, _ := http.Get("https://www.nasa.gov/sitemap.xml") //This also works check this		//getting response from the site

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)						//reading response 
	//string_body := string(bytes)
	//fmt.Println(string_body)
	//resp.Body.Close()

	var s Sitemapindex
	var n News
	news_map :=make(map[string]NewsMap)
  	xml.Unmarshal(bytes, &s)
  	//fmt.Println(s.Locations)				//prints that as a sought of dictionary

  /*	for _,Location := range s.Locations {	//iteraring through the s Location struct wise
  		fmt.Printf("\n%s",Location)
  	}
*/

  	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Keywords {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

		for idx, data := range news_map {
		fmt.Println("\n\n\n\n\n",idx)
		fmt.Println("\n",data.Keyword)
		fmt.Println("\n",data.Location)
	}

	resp.Body.Close()											//closing the resources
}