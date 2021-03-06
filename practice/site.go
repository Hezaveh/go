package main

import ("fmt"
		"net/http"
		"html/template"
		"io/ioutil"
		"encoding/xml")

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

func newsAggHandler(w http.ResponseWriter, r *http.Request){


	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)

	//resp, _ := http.Get("https://www.washingtonpost.com/arcio/news-sitemap/")
	//bytes, _ := ioutil.ReadAll(resp.Body)
	bytes := washPostXML
	//string_body := string(bytes)
	//fmt.Println(string_body)
	//resp.Body.Close()
	xml.Unmarshal(bytes, &s)

	//fmt.Println(s.Locations)

	for _, Location := range s.Locations{
		//fmt.Printf("%s\n", Location)
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.Keywords {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
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