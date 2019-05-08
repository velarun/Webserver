package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type sqldb struct {
	Username string `json:"sql_user"`
	Password string `json:"sql_pass"`
	Dbserver string `json:"sql_db"`
}

func connectDatabse() {

	pwd, _ := os.Getwd()
	jsonFile, err := os.Open(pwd + "/config.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Read config.json file successfully")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var sqldb sqldb
	json.Unmarshal(byteValue, &sqldb)

	connect := fmt.Sprintf("%s:%s@tcp(%s)/", sqldb.Username, sqldb.Password, sqldb.Dbserver)
	db, err = sql.Open("mysql", connect)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected successfully")

	query, err := ioutil.ReadFile(pwd + "/hotstar.sql")
	if err != nil {
		log.Fatal(err)
	}

	requests := strings.Split(string(query), ";")

	for _, request := range requests {
		if len(request) > 0 {
			_, err := db.Exec(request)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Println("Default Database Loaded successfully")
}
