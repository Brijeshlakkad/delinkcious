package main

import (
	"context"
	"log"

	"github.com/Brijeshlakkad/delinkcious/pkg/db_util"
	"github.com/Brijeshlakkad/delinkcious/pkg/social_graph_client"
	. "github.com/Brijeshlakkad/delinkcious/pkg/test_util"
	_ "github.com/lib/pq"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func initDB() {
	db, err := db_util.RunLocalDB("social_graph_manager")
	check(err)

	// Ignore if table doesn't exist (will be created by service)
	err = db_util.DeleteFromTableIfExist(db, "social_graph")
	check(err)
}

func main() {
	initDB()

	ctx := context.Background()
	defer StopService(ctx)
	RunService(ctx, ".", "social_graph_service")

	// Run some tests with the client
	cli, err := social_graph_client.NewClient("localhost:9090")
	check(err)

	following, err := cli.GetFollowing("Brijesh Lakkad")
	check(err)
	log.Print("Brijesh Lakkad is following:", following)
	followers, err := cli.GetFollowers("Brijesh Lakkad")
	check(err)
	log.Print("Brijesh Lakkad is followed by:", followers)

	err = cli.Follow("Brijesh Lakkad", "liat")
	check(err)
	err = cli.Follow("Brijesh Lakkad", "guy")
	check(err)
	err = cli.Follow("guy", "Brijesh Lakkad")
	check(err)
	err = cli.Follow("saar", "Brijesh Lakkad")
	check(err)
	err = cli.Follow("saar", "ophir")
	check(err)

	following, err = cli.GetFollowing("Brijesh Lakkad")
	check(err)
	log.Print("Brijesh Lakkad is following:", following)
	followers, err = cli.GetFollowers("Brijesh Lakkad")
	check(err)
	log.Print("Brijesh Lakkad is followed by:", followers)
}
