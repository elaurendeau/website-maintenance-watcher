package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	fmt.Printf("token: '%s'\n", os.Args[1])
	fmt.Printf("user: '%s'\n", os.Args[2])
	websiteUrl := "http://www.irrigationdepot.ca/boutique/english/home.html"
	expected := "http://www.irrigationdepot.ca/messages/maintenance.html"

	notify("Maintenance watcher, starting", "The maintenance watcher is now active.")

	for {
			fmt.Println(time.Now().String() + " Querying " + websiteUrl)
			resp, _ := http.Get(websiteUrl)
			fmt.Println("Response " + resp.Request.URL.String())
			if resp.Request.URL.String() != expected {
				fmt.Println("Maintenance is over!!")
				notify("Maintenance watcher, website is up", "Head to " + websiteUrl + " and start shopping.")
				time.Sleep(10 * time.Minute)
			} else {
				fmt.Println("Still under maintenance, waiting longer..")
			}
			time.Sleep(2 * time.Minute)
		}

}

func notify(title string, message string) {
	formValues := url.Values{
		"token": {os.Args[1]},
		"user": {os.Args[2]},
		"device": {"galaxynote9"},
		"title": {title},
		"message": {message}}

	http.PostForm("https://api.pushover.net/1/messages.json", formValues)
}
