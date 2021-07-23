package main

import ("fmt"
		"net/http"
		"html/template"
		"io/ioutil"
		"encoding/xml"
		"sync")

var washPostXML = []byte(`
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>
`)

var wg sync.WaitGroup

type SitemapIndex struct{
	Locations []string `xml:"sitemap>loc"`
}

type News struct{
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct{
	Keyword string
	Location string
}

type NewsAggPage struct {
	Title string
	News map[string]NewsMap
}

func newsRoutine(c chan News, Location string){
	defer wg.Done()
	var n News
	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()
	c <- n
}

func newsAggHandler(w http.ResponseWriter, r *http.Request){


	var s SitemapIndex
	
	news_map := make(map[string]NewsMap)

	//resp, _ := http.Get("https://www.washingtonpost.com/arcio/news-sitemap/")
	//bytes, _ := ioutil.ReadAll(resp.Body)
	bytes := washPostXML
	//string_body := string(bytes)
	//fmt.Println(string_body)
	//resp.Body.Close()
	xml.Unmarshal(bytes, &s)

	queue := make(chan News, 30)

	for _, Location := range s.Locations{
		wg.Add(1)
		go newsRoutine(queue, Location)
	}

	wg.Wait()
	close(queue)

	for elem := range queue {
		for idx, _ := range elem.Keywords {
			news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
		}
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
	t, err := template.ParseFiles("basictemplating.html")
	fmt.Println(err)
	//fmt.Println(t.Execute(w, p))
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "<h1>Whoa, go is neat!</h1>")
}

func main() {
	http.HandleFunc ("/", indexHandler)
	http.HandleFunc ("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}