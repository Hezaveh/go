package main

import ("fmt"
		"encoding/xml"
		"net/http"
		"io/ioutil")




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

/*type SitemapIndex struct{
	Locations []Location `xml:"sitemap"`
	//Locations []Location `xml:"url"`
}

type Location struct{
	Loc string `xml:"loc"`
}*/

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

/*func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}*/

func main(){

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

	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}