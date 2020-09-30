package parks

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/olebedev/config"
)

var (
	Client *sql.DB
	cfg    *config.Config
)

func init() {
	var err error
	var dsnStr string
	var dsnErr error

	if os.Getenv("DEBUG") == "1" {
		log.Println("+++> running locally")
		file, err := ioutil.ReadFile("conf.json")
		if err != nil {
			panic(err)
		}
		jsonString := string(file)

		cfg, err = config.ParseJson(jsonString)
		dsnStr, dsnErr = cfg.String("mysql.local.dsn")
		if dsnErr != nil {
			panic(err)
		}
	} else {
		log.Println("+++> running aws")
		// dsnStr = "root:root@tcp(mysql_parks:3306)/parks?parseTime=true"
		dsnStr = "root:root@<aws-rds-endpoint>/parks?parseTime=true"
	}

	Client, err = sql.Open("mysql", dsnStr)
	if err != nil {
		panic(err)
	}

	log.Println("+++> database successfully configured")
	log.Println("+++> database successfully configured")
	log.Println("+++> database successfully configured")
	log.Println("+++> database successfully configured")
}
