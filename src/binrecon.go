/*
 ▄▄▄▄    ██▓ ███▄    █  ██▀███  ▓█████  ▄████▄   ▒█████   ███▄    █
▓█████▄ ▓██▒ ██ ▀█   █ ▓██ ▒ ██▒▓█   ▀ ▒██▀ ▀█  ▒██▒  ██▒ ██ ▀█   █
▒██▒ ▄██▒██▒▓██  ▀█ ██▒▓██ ░▄█ ▒▒███   ▒▓█    ▄ ▒██░  ██▒▓██  ▀█ ██▒
▒██░█▀  ░██░▓██▒  ▐▌██▒▒██▀▀█▄  ▒▓█  ▄ ▒▓▓▄ ▄██▒▒██   ██░▓██▒  ▐▌██▒
░▓█  ▀█▓░██░▒██░   ▓██░░██▓ ▒██▒░▒████▒▒ ▓███▀ ░░ ████▓▒░▒██░   ▓██░
░▒▓███▀▒░▓  ░ ▒░   ▒ ▒ ░ ▒▓ ░▒▓░░░ ▒░ ░░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░   ▒ ▒
▒░▒   ░  ▒ ░░ ░░   ░ ▒░  ░▒ ░ ▒░ ░ ░  ░  ░  ▒     ░ ▒ ▒░ ░ ░░   ░ ▒░
 ░    ░  ▒ ░   ░   ░ ░   ░░   ░    ░   ░        ░ ░ ░ ▒     ░   ░ ░
 ░       ░           ░    ░        ░  ░░ ░          ░ ░           ░
      ░                                ░
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type paste struct {
	ScrapeURL string `json:"scrape_url"`
	FullURL   string `json:"full_url"`
	Date      string `json:"date"`
	Key       string `json:"key"`
	Size      string `json:"size"`
	Expire    string `json:"expire"`
	Title     string `json:"title"`
	Syntax    string `json:"text"`
	User      string `json:"user"`
	Data      string `json:"data"`
}

func pullLatest() []paste {
	url := "https://scrape.pastebin.com/api_scraping.php?limit=250"

	reconClient := http.Client{
		Timeout: time.Second * 15, //Maximum of 15 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "binrecon")

	res, getErr := reconClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	pastes := make([]paste, 0)
	jsonErr := json.Unmarshal([]byte(body), &pastes)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return pastes
}

func pullPaste(pasteKey string) string {
	url := "https://scrape.pastebin.com/api_scrape_item.php?i=" + pasteKey

	reconClient := http.Client{
		Timeout: time.Second * 2, //Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "binrecon")

	res, getErr := reconClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return string(body)

}

//https://gist.github.com/ryanfitz/4191392
func doEvery(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

func recon() {
	latest := pullLatest()
	for i := range latest {
		latest[i].Data = pullPaste(latest[i].Key)
		//latest[i].Data = "data here"
		fmt.Println("[*] Saving paste " + latest[i].Title + " {" + latest[i].Key + "}")
		pasteJSON, _ := json.Marshal(latest[i])

		file, err := os.OpenFile("/collections/"+"recon-"+latest[i].Key+".json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0744)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		fmt.Fprintf(file, string(pasteJSON)+"\n")
	}
}

func main() {
	recon()
	doEvery(10*time.Minute, recon)
}
