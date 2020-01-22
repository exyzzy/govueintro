package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//For configlocaldb.json
type ConfigurationDB struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Name    string
	SSLMode string
}

var configdb ConfigurationDB
var Db *sql.DB

//load local db config
func loadConfigDB() {
	file, err := os.Open("data/configlocaldb.json")
	if err != nil {
		log.Fatalln("Cannot open configlocaldb file", err)
	}
	decoder := json.NewDecoder(file)
	configdb = ConfigurationDB{}
	err = decoder.Decode(&configdb)
	if err != nil {
		log.Fatalln("Cannot get local configurationdb from file", err)
	}
}

//init db
func init() {
	var err error

	psqlInfo := os.Getenv("DATABASE_URL")
	if psqlInfo == "" {
		loadConfigDB()
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s "+
			"dbname=%s sslmode=%s", configdb.Host, configdb.Port, configdb.User, configdb.Pass, configdb.Name, configdb.SSLMode)
	}

	fmt.Println(psqlInfo)

	Db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalln("DB Fail to open", err)
	} else {
		err = Db.Ping() //check for silent fail
		if err != nil {
			log.Fatalln("DB Silent Fail after open", err)
		}
	}

	return
}
