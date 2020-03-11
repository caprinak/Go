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
	//`xml:"url>news>titles"`
	LastMods []string `xml:"sitemap>lastmod"`
	Locations []string `xml:"sitemap>loc"`
}

func main() {
	var s SitemapIndex
	//var n News

	resp, _ := http.Get("https://vnexpress.net/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	/* stringbody := string(bytes)
	fmt.Println(stringbody) */
	resp.Body.Close()	
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)
	// range for iteration


	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	} 


}
