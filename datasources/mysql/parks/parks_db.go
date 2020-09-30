package parks

import (
	"database/sql"
	"io/ioutil"
	"log"

	"github.com/federicoleon/bookstore_utils-go/logger"
	"github.com/go-sql-driver/mysql"
	"github.com/olebedev/config"
)

var (
	Client *sql.DB
	cfg    *config.Config
)

func init() {
	var err error
	// var color string
	log.Println("+++> cfg 1")
	log.Println("+++> cfg 1")
	log.Println("+++> cfg 1")

	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("+++> cfg 1.01 dir: %+v\n", dir)

	// data, err := ioutil.ReadFile("hello.txt")
	// if err != nil {
	// 	fmt.Println("File reading error", err)
	// 	return
	// }
	// log.Println("+++> cfg 1.1")
	// log.Printf("+++> cfg 1.2 hello: %+v\n", data)
	// ggg, _ := config.ParseJsonFile("conf.json")

	file, err := ioutil.ReadFile("conf.json")
	if err != nil {
		panic(err)
	}
	log.Println("+++> cfg 2.1")
	jsonString := string(file)
	log.Println("+++> cfg 2.2")

	cfg, err = config.ParseJson(jsonString)
	log.Println("+++> cfg 2.3")
	dsnStr, dsnErr := cfg.String("mysql.local.dsn")
	if dsnErr != nil {
		panic(err)
	}
	log.Printf("+++> cfg 2.4 dsnStr %+v\n", dsnStr)
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

	// Client, err = sql.Open("mysql", "root:root@tcp(mysql_parks:3306)/parks?parseTime=true")
	Client, err = sql.Open("mysql", dsnStr)
	log.Println("+++> cfg 2.5 yay")

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
