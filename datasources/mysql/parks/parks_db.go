package parks

import (
	"database/sql"
	"log"

	"github.com/federicoleon/bookstore_utils-go/logger"
	"github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
	// cfg    *config.Config
)

func init() {
	var err error
	// log.Println("+++> cfg 1")
	// log.Println("+++> cfg 1")
	// log.Println("+++> cfg 1")
	// ggg, _ := config.ParseJsonFile("conf.json")
	// cfg, err = config.ParseJsonFile("conf.json")
	// log.Println("+++> cfg 2")
	// log.Printf("+++> cfg 3 %+v\n", ggg.UString("mysql.dsn"))
	// log.Println("+++> cfg 4")
	// log.Println("+++> cfg 4")
	// log.Println("+++> cfg 4")

	// util.LogFatal(err)

	// db, err = sqlx.Connect("mysql", cfg.UString("mysql.dsn"))
	// util.LogFatal(err)

	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
	// 	"root", "root", "127.0.0.1:3306", "parks",
	// )
	// Client, err = sql.Open("mysql", dataSourceName)
	//          "root:@tcp(127.0.0.1:3306)/buildtemplate0818_jwt"
	// Client, err = sql.Open("mysql", "test:root@/parks")
	// Client, err = sql.Open("mysql", "test:root@tcp(127.0.0.1:3306)/parks")
	// Client, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/parks")
	// Client, err = sql.Open("mysql", "root:@/parks")
	// DB, err = sql.Open("mysql", "root:pasaribu@tcp(database.dev:3306)/shopee")
	Client, err = sql.Open("mysql", "root:root@tcp(mysql_parks:3306)/parks?parseTime=true")
	// Client, err = sql.Open("mysql", cfg.UString("mysql.dsn"))
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	mysql.SetLogger(logger.GetLogger())
	log.Println("+++> database successfully configured")
	log.Println("+++> database successfully configured")
	log.Println("+++> database successfully configured")
	log.Println("+++> database successfully configured")
}

// func init() {
// 	fmt.Println("+++> init 2m2")
// 	if skipForTest() {
// 		return
// 	}
// 	// connString := "host=db port=5432 user=postgres password=secret dbname=bird_encyclopedia sslmode=disable"
// 	// connString := "user=postgres password=secret dbname=bird_encyclopedia sslmode=disable"
// 	connString := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		"db", 5432, "postgres", "secret", "bird_encyclopedia")

// 	db, err = sqlx.Connect("postgres", connString)
// 	// db, err := sql.Open("postgres", connString)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Println("+++> Connection failure!")
// 		panic(err)
// 	}

// 	fmt.Println("+++> Successfully connected!")
// }
