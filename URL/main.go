package main

import (
	"fmt"
	"net/url"
)

func main() {
	myurl := "https://api.example.com:8080/v1/users/12345?api_key=abcd1234"

	parseUrl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println(" ", err)
		return

	}
	fmt.Println("", parseUrl)
	parseUrl.Path = "/NEWPATH"
	parseUrl.RawQuery = "username=HarshJain"

	newUrl := parseUrl.String()
	fmt.Println(" ", newUrl)

}
