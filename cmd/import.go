package cmd

import (
	"bufio"
	"os"

	"github.com/aaa59891/ticket/models"

	"github.com/aaa59891/ticket/utils"

	cli "gopkg.in/urfave/cli.v1"
)

var (
	ImportInvitationCode = cli.Command{
		Name:        "importCode",
		Aliases:     []string{"ic"},
		Usage:       "Import invitation codes",
		Subcommands: []cli.Command{subIcFile},
	}
)

var subIcFile = cli.Command{
	Name:    "file",
	Aliases: []string{"f"},
	Usage:   "Import invitation codes from file",
	Action:  importCodeFromFile,
}

func importCodeFromFile(c *cli.Context) error {
	dir, err := utils.GetCurrentDir()
	if err != nil {
		return err
	}
	file, err := os.Open(dir + "/" + c.Args().Get(0))
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := models.Ticket{
			Code: scanner.Text(),
		}
		if err = utils.Transactional(t.Create); err != nil {
			return err
		}
	}
	return nil
}
