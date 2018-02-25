package main

import (
	"log"
	"os"

	"github.com/aaa59891/ticket/cmd"
	"github.com/aaa59891/ticket/db"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	defer func() {
		if db.DB != nil {
			db.DB.Close()
		}
	}()
	app := cli.NewApp()
	app.Commands = []cli.Command{
		cmd.ImportInvitationCode,
		cmd.List,
		cmd.Email,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
