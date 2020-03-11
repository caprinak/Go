package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	LastMods []string `xml:"url>lastmod"`
	Locations []string `xml:"url>loc"`
}
type NewsMap struct {
	LastMod string
	Location string
}

func main() {
	//var s SitemapIndex
	var n News
	news_map := make(map[string]string)
	resp, _ := http.Get("https://vnexpress.net/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	
	resp.Body.Close()	
	xml.Unmarshal(bytes, &n)

	for idx, _ := range n.Locations {
		news_map[NewsMap{n.Locations[idx]}] = NewsMap{n.LastMods[idx]}
		//fmt.Printf("\n%s", Location)
	} 
	fmt. Println (news_map)

}