package main

import (
	"context"
	"log"

	"github.com/Brijeshlakkad/delinkcious/pkg/db_util"
	om "github.com/Brijeshlakkad/delinkcious/pkg/object_model"
	. "github.com/Brijeshlakkad/delinkcious/pkg/test_util"
	"github.com/Brijeshlakkad/delinkcious/pkg/user_client"
	_ "github.com/lib/pq"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func initDB() {
	db, err := db_util.RunLocalDB("user_manager")
	if err != nil {
		return
	}

	tables := []string{"sessions", "users"}
	for _, table := range tables {
		err = db_util.DeleteFromTableIfExist(db, table)
		check(err)
	}
}

func main() {
	initDB()

	ctx := context.Background()
	defer StopService(ctx)
	RunService(ctx, ".", "user_service")

	// Run some tests with the client
	cli, err := user_client.NewClient("localhost:7070")
	check(err)

	err = cli.Register(om.User{"gg@gg.com", "Brijesh Lakkad"})
	check(err)
	log.Print("Brijesh Lakkad has registered successfully")

	session, err := cli.Login("Brijesh Lakkad", "secret")
	check(err)
	log.Print("Brijesh Lakkad has logged in successfully. the session is: ", session)

	err = cli.Logout("Brijesh Lakkad", session)
	check(err)
	log.Print("Brijesh Lakkad has logged out successfully.")

}
