package link_checker_events

import (
	"log"

	om "github.com/Brijeshlakkad/delinkcious/pkg/object_model"
	"github.com/nats-io/nats.go"
)

type eventSender struct {
	hostname string
	nats     *nats.EncodedConn
}

func (s *eventSender) OnLinkChecked(username string, url string, status om.LinkStatus) {
	err := s.nats.Publish(subject, Event{username, url, status})
	if err != nil {
		log.Fatal(err)
	}
}

func NewEventSender(url string) (om.LinkCheckerEvents, error) {
	ec, err := connect(url)
	if err != nil {
		return nil, err
	}
	return &eventSender{hostname: url, nats: ec}, nil
}
