package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
)

type ticketConfig struct {
	DBConfig struct {
		Host         string
		Port         int
		Account      string
		Password     string
		DatabaseName string
	}
	EmailConfig struct {
		Account  string
		Password string
	}
	Security struct {
		EncryptKey string
	}
}

var (
	FilePath string
	config   *ticketConfig
)

func (ticketConfig ticketConfig) GetDBConnectString() string {
	dbConfig := ticketConfig.DBConfig
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.Account, dbConfig.DatabaseName, dbConfig.Password)
}

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	FilePath = usr.HomeDir + "/.ticket_config.json"
	b, err := ioutil.ReadFile(FilePath)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(b, &config); err != nil {
		log.Fatal(err)
	}
}

func GetConfig() ticketConfig {
	return *config
}
