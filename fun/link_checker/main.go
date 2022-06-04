package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Brijeshlakkad/delinkcious/pkg/link_checker"
	"github.com/Brijeshlakkad/delinkcious/pkg/link_checker_events"
	om "github.com/Brijeshlakkad/delinkcious/pkg/object_model"
)

type Log struct {
}

const natsUrl = "nats-cluster.default.svc.cluster.local:4222"

type Response struct {
	Status string `json:"content"`
	Err    string `json:"err"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var e om.CheckLinkRequest
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := e.Username
	url := e.Url
	if username == "" || url == "" {
		errorMessage := fmt.Sprintf("missing USERNAME ('%s') and/or URL ('%s')", username, url)

		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	status := om.LinkStatusValid
	err = link_checker.CheckLink(url)
	if err != nil {
		status = om.LinkStatusInvalid
	}

	response := Response{Status: status}

	sender, err := link_checker_events.NewEventSender(natsUrl)
	if err != nil {
		response.Err = err.Error()
	} else {
		sender.OnLinkChecked(username, url, status)
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func main() {
	l := &Log{}
	l.Info("Hello world sample started.")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("ListenAndServe error:%s ", err.Error())
	}
}

func (log *Log) Infof(format string, a ...interface{}) {
	log.log("INFO", format, a...)
}

func (log *Log) Info(msg string) {
	log.log("INFO", "%s", msg)
}

func (log *Log) Errorf(format string, a ...interface{}) {
	log.log("ERROR", format, a...)
}

func (log *Log) Error(msg string) {
	log.log("ERROR", "%s", msg)
}

func (log *Log) Fatalf(format string, a ...interface{}) {
	log.log("FATAL", format, a...)
}

func (log *Log) Fatal(msg string) {
	log.log("FATAL", "%s", msg)
}

func (log *Log) log(level, format string, a ...interface{}) {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	ft := fmt.Sprintf("%s %s %s\n", time.Now().In(cstSh).Format("2006-01-02 15:04:05"), level, format)
	fmt.Printf(ft, a...)
}
