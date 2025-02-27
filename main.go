package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	proxyUrl, err := url.Parse("http://spp8xnfb0l:EP7qbckL9thF=4inx4@pl.smartproxy.com:20001")
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
	}
	req, err := http.NewRequest("GET", "https://info-car.pl/new/", nil)
	if err != nil {
		log.Println(err)
	}

	res, err := client.Do(req)
	log.Println(res)

	if err != nil {
		log.Println(err)
	}
}
