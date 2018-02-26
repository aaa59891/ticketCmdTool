package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aaa59891/ticket/db"
	"github.com/aaa59891/ticket/models"

	"github.com/aaa59891/ticket/utils"

	"github.com/aaa59891/ticket/config"
	"gopkg.in/urfave/cli.v1"
)

const (
	flagAccount  = "account"
	flagPassword = "password"
	flagFile     = "file"
	flagEmali    = "email"

	subject = "My journey in nano bigdata 邀請碼（上午場）"
	body    = `
感謝您的參與
以下是您的邀請碼： %s
	`
)

var (
	Email = cli.Command{
		Name:  "email",
		Usage: "Email's operations",
		Subcommands: []cli.Command{
			subEmailConfig,
			subSendEmail,
		},
	}
)

var subEmailConfig = cli.Command{
	Name:  "config",
	Usage: "Set email's informatino. --account=account --password=password",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  flagAccount,
			Usage: "Email's account",
		},
		cli.StringFlag{
			Name:  flagPassword,
			Usage: "Email's password",
		},
	},
	Action: setEmailInfo,
}

func setEmailInfo(c *cli.Context) error {
	if c.NumFlags() == 0 {
		return nil
	}
	cfg := config.GetConfig()
	newAccount := c.String(flagAccount)
	newPassword := c.String(flagPassword)

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

var subSendEmail = cli.Command{
	Name:  "send",
	Usage: "Send email: --file=file or --email=test@email.com",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "file",
		},
		cli.StringFlag{
			Name: "email",
		},
	},
	Action: sendEmail,
}

func sendEmail(c *cli.Context) error {
	if c.NumFlags() == 0 {
		return nil
	}
	file := c.String(flagFile)
	if len(file) > 0 {
		return sendEmailFromFile(file)
	}
	email := c.String(flagEmali)
	return sendEmailSingle(email)
}

func sendEmailFromFile(file string) error {
	dir, err := utils.GetCurrentDir()
	if err != nil {
		return err
	}
	filePath := dir + "/" + file
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err := sendEmailSingle(scanner.Text()); err != nil {
			return err
		}
		utils.RemoveFirstLine(filePath)
	}

	return nil
}

func sendEmailSingle(dest string) error {
	cfg := config.GetConfig()
	decrypted, err := utils.Decrypt(cfg.Security.EncryptKey, cfg.EmailConfig.Password)
	if err != nil {
		return err
	}
	ticket := models.Ticket{}
	if err := db.DB.First(&ticket, "email = ?", "").Error; err != nil {
		return err
	}
	mail := utils.Email{
		Address:  cfg.EmailConfig.Account,
		Password: decrypted,
		Subject:  subject,
		Body:     fmt.Sprintf(body, ticket.Code),
		Type:     utils.EmailTypeGmail,
	}
	if err := utils.SendEmail(mail, dest); err != nil {
		return err
	}
	ticket.Email = dest
	if err := db.DB.Save(&ticket).Error; err != nil {
		return err
	}
	return nil
}
