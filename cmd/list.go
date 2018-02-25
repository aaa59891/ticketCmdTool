package cmd

import (
	"fmt"

	"github.com/aaa59891/ticket/db"
	"github.com/aaa59891/ticket/models"
	"gopkg.in/urfave/cli.v1"
)

var (
	List = cli.Command{
		Name:  "list",
		Usage: "List tickets.",
		Subcommands: []cli.Command{
			subCmdListAllTickets,
			subCmdListAvaibleTickets,
		},
	}
)

var subCmdListAllTickets = cli.Command{
	Name:   "all",
	Usage:  "List all of the tickets.",
	Action: listAllTickets,
}

func listAllTickets(c *cli.Context) error {
	tickets := make([]models.Ticket, 0)
	if err := db.DB.Find(&tickets).Error; err != nil {
		return err
	}

	for _, ticket := range tickets {
		fmt.Println(ticket)
	}
	return nil
}

var subCmdListAvaibleTickets = cli.Command{
	Name:   "avaible",
	Usage:  "List the avaible tickets",
	Action: listAvaibleTickets,
}

func listAvaibleTickets(c *cli.Context) error {
	tickets := make([]models.Ticket, 0)
	if err := db.DB.Find(&tickets, "email = ?", "").Error; err != nil {
		return err
	}
	for _, ticket := range tickets {
		fmt.Println(ticket.StringWithoutEmail())
	}
	return nil
}
