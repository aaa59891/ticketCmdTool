package cmd

import (
	"encoding/json"
	"io/ioutil"

	"github.com/aaa59891/ticket/utils"

	"github.com/aaa59891/ticket/config"
	"gopkg.in/urfave/cli.v1"
)

const (
	account  = "account"
	password = "password"
)

var (
	Email = cli.Command{
		Name:  "email",
		Usage: "Set email's information.",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  account,
				Usage: "Email's account",
			},
			cli.StringFlag{
				Name:  password,
				Usage: "Email's password",
			},
		},
		Action: setEmailInfo,
	}
)

func setEmailInfo(c *cli.Context) error {
	if c.NumFlags() == 0 {
		return nil
	}
	cfg := config.GetConfig()
	newAccount := c.String(account)
	newPassword := c.String(password)

	if len(newAccount) > 0 {
		cfg.EmailConfig.Account = newAccount
	}
	if len(newPassword) > 0 {
		encryptd, err := utils.Encrypt(cfg.Security.EncryptKey, newPassword)
		if err != nil {
			return err
		}
		cfg.EmailConfig.Password = encryptd
	}

	b, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return err
	}

	ioutil.WriteFile(config.FilePath, b, 0644)

	return nil
}
