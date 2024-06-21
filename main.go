package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/bhagyaraj1208117/url/URLName"
)

func main() {
	fmt.Println("hell ")
	url := flag.String("url", "https://www.youtube.com/", "Enter url")
	flag.Parse()
	fmt.Println("Entered url", *url)
	urlValue := strings.Split(*url, "/")
	length := len(urlValue) - 1
	fmt.Println("url path", urlValue[length])
	path := URLName.NewUrl(urlValue[length])

	length--
	domainName := URLName.NewUrl(urlValue[length])

	fmt.Println("Domain name", domainName.Name)
	fmt.Println("path ", path.Name)
}
