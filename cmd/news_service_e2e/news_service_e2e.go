package main

import (
	"context"
	"fmt"
	"os"
	"time"

	lme "github.com/Brijeshlakkad/delinkcious/pkg/link_manager_events"
	nmc "github.com/Brijeshlakkad/delinkcious/pkg/news_manager_client"
	om "github.com/Brijeshlakkad/delinkcious/pkg/object_model"
	. "github.com/Brijeshlakkad/delinkcious/pkg/test_util"
)

func runNewsService(ctx context.Context) {
	// Set environment
	err := os.Setenv("PORT", "6060")
	Check(err)

	err = os.Setenv("NATS_CLUSTER_SERVICE_HOST", "nats://localhost")
	Check(err)

	err = os.Setenv("NATS_CLUSTER_SERVICE_PORT", "4222")
	Check(err)

	err = os.Setenv("NEWS_MANAGER_REDIS_SERVICE_HOST", "localhost")
	Check(err)

	err = os.Setenv("NEWS_MANAGER_REDIS_SERVICE_PORT", "6379")
	Check(err)

	if os.Getenv("RUN_NEWS_SERVICE") != "false" {
		RunService(ctx, ".", "news_service")
	}
}

func main() {
	ctx := context.Background()
	defer KillServer(ctx)

	fmt.Println("Launching local NATS server...")
	err := RunLocalNatsServer()
	Check(err)
	time.Sleep(time.Second * 1)

	fmt.Println("Launching local Redis server...")
	err = RunLocalRedisServer()
	Check(err)
	time.Sleep(time.Second * 1)

	if os.Getenv("RUN_NEWS_SERVICE") != "false" {
		fmt.Println("Launching local news service...")
		runNewsService(ctx)
		time.Sleep(time.Second * 1)
	}

	fmt.Println("creating news event sender...")
	cli, disconnectFunc, err := nmc.NewClient(":6060")
	defer disconnectFunc()
	Check(err)

	sender, err := lme.NewEventSender("nats://localhost:4222")
	Check(err)

	time.Sleep(time.Second * 1)
	fmt.Println("Sending OnLinkAdded event...")
	sender.OnLinkAdded("brijesh", &om.Link{Url: "http://example.org", Title: "Example"})
	time.Sleep(time.Second * 1)

	fmt.Println("Fetching news...")
	res, err := cli.GetNews(om.GetNewsRequest{Username: "brijesh"})
	Check(err)

	for _, e := range res.Events {
		fmt.Println(*e)
	}
}
