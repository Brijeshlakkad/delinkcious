package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	net_url "net/url"
	"os"
	"os/exec"
	"time"

	om "github.com/Brijeshlakkad/delinkcious/pkg/object_model"
	. "github.com/Brijeshlakkad/delinkcious/pkg/test_util"
)

var (
	delinkciousUrl   string
	delinkciousToken = os.Getenv("DELINKCIOUS_TOKEN")
	httpClient       = http.Client{}
)

func getLinks() {
	req, err := http.NewRequest("GET", string(delinkciousUrl)+"/links", nil)
	Check(err)

	req.Header.Add("Access-Token", delinkciousToken)
	r, err := httpClient.Do(req)
	Check(err)

	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		Check(errors.New(r.Status))
	}

	var glr om.GetLinksResult
	body, err := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(body, &glr)
	Check(err)

	log.Println("======= Links =======")
	for _, link := range glr.Links {
		log.Println(fmt.Sprintf("title: '%s', url: '%s', status: '%s', description: '%s'", link.Title,
			link.Url,
			link.Status,
			link.Description))
	}
}

func addLink(url string, title string) {
	params := net_url.Values{}
	params.Add("url", url)
	params.Add("title", title)
	qs := params.Encode()

	log.Println("===== Add Link ======")
	log.Println(fmt.Sprintf("Adding new link - title: '%s', url: '%s'", title, url))

	url = fmt.Sprintf("%s/links?%s", delinkciousUrl, qs)
	req, err := http.NewRequest("POST", url, nil)
	Check(err)

	req.Header.Add("Access-Token", delinkciousToken)
	r, err := httpClient.Do(req)
	Check(err)
	if r.StatusCode != http.StatusOK {
		defer r.Body.Close()
		bodyBytes, err := ioutil.ReadAll(r.Body)
		Check(err)
		message := r.Status + " " + string(bodyBytes)
		Check(errors.New(message))
	}
}

func deleteLink(url string) {
	params := net_url.Values{}
	params.Add("url", url)
	qs := params.Encode()

	url = fmt.Sprintf("%s/links?%s", delinkciousUrl, qs)
	req, err := http.NewRequest("DELETE", url, nil)
	Check(err)

	req.Header.Add("Access-Token", delinkciousToken)
	r, err := httpClient.Do(req)
	Check(err)
	if r.StatusCode != http.StatusOK {
		defer r.Body.Close()
		bodyBytes, err := ioutil.ReadAll(r.Body)
		Check(err)
		message := r.Status + " " + string(bodyBytes)
		Check(errors.New(message))
	}
}

func main() {
	tempUrl, err := exec.Command("minikube", "service", "api-gateway", "--url").CombinedOutput()
	delinkciousUrl = string(tempUrl[:len(tempUrl)-1]) + "/v1"
	Check(err)

	// Delete link
	deleteLink("https://github.com/Brijeshlakkad")

	// Get links
	getLinks()

	// Add a new link
	addLink("https://github.com/Brijeshlakkad", "Brijesh Lakkad on Github")

	// Get links again
	getLinks()

	// Wait a little and get links again
	time.Sleep(time.Second * 3)
	getLinks()
}
