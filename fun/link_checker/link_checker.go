package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io/ioutil"

	"github.com/Brijeshlakkad/delinkcious/pkg/link_checker"
	"github.com/Brijeshlakkad/delinkcious/pkg/link_checker_events"
	om "github.com/Brijeshlakkad/delinkcious/pkg/object_model"
)

const natsUrl = "nats-cluster.default.svc.cluster.local:4222"

type msg struct {
	Content string `json:"content"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	var e om.CheckLinkRequest
	err = json.Unmarshal(reqBody, &e)

	response := msg{}

	respBody, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	username := e.Username
	url := e.Url
	if username == "" || url == "" {
		errorMessage := fmt.Sprintf("missing USERNAME ('%s') and/or URL ('%s')", username, url)

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorMessage))
	}

	status := om.LinkStatusValid
	err = link_checker.CheckLink(url)
	if err != nil {
		status = om.LinkStatusInvalid
	}

	sender, err := link_checker_events.NewEventSender(natsUrl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	sender.OnLinkChecked(username, url, status)
	w.Write([]byte(respBody))
}

func main() {

}
