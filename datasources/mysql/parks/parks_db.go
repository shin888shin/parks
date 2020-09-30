package parks

import (
	"database/sql"
	"io/ioutil"
	"log"

	"github.com/olebedev/config"
)

var (
	Client *sql.DB
	cfg    *config.Config
)

func init() {
	var err error

	file, err := ioutil.ReadFile("conf.json")
	if err != nil {
		panic(err)
	}
	jsonString := string(file)

	cfg, err = config.ParseJson(jsonString)
	dsnStr, dsnErr := cfg.String("mysql.local.dsn")
	if dsnErr != nil {
		panic(err)
	}
	Client, err = sql.Open("mysql", dsnStr)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("+++> database successfully configured")
	log.Println("+++> database successfully configured")
	log.Println("+++> database successfully configured")
	log.Println("+++> database successfully configured")
}
