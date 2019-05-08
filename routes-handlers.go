package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type trayTable struct {
	ID               sql.NullInt64 `json:"Id"`
	Timestmp         string        `json:"Timestmp"`
	UserID           string        `json:"User_id"`
	IsLoggedIn       string        `json:"Is_logged_in"`
	MissingTrayTitle string        `json:"Missing_tray_title"`
	AddedTrayTitle   string        `json:"Added_tray_title"`
}

func renderHome(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/index.html")
}

func getID(response http.ResponseWriter, request *http.Request) {
	var table trayTable

	id := mux.Vars(request)["id"]

	rows, err := db.Query("select * from trayTable where id = ?", id)
	//rows, err := db.Query("select * from trayTable where id = " + id)
	if err != nil {
		fmt.Println(err)
		returnErrorResponse(response, request)
	}

	for rows.Next() {
		err = rows.Scan(&table.ID, &table.Timestmp, &table.UserID, &table.IsLoggedIn, &table.MissingTrayTitle, &table.AddedTrayTitle)
		if err != nil {
			returnErrorResponse(response, request)
		}
	}

	defer rows.Close()

	if table.Timestmp == "" {

		jsonResponse, err := json.Marshal("No rows available")
		if err != nil {
			panic(err)
		}
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(jsonResponse)

	} else {

		jsonResponse, jsonError := json.Marshal(table)
		if jsonError != nil {
			fmt.Println(jsonError)
			returnErrorResponse(response, request)
		}

		if jsonResponse == nil {
			returnErrorResponse(response, request)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(jsonResponse)
		}
	}
}

func insertRow(response http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	var table trayTable
	err := decoder.Decode(&table)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO trayTable VALUES(?,?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		returnErrorResponse(response, request)
	}

	res, err := stmt.Exec(sql.NullInt64{}, time.Now().UTC(), table.UserID, table.IsLoggedIn, table.MissingTrayTitle, table.AddedTrayTitle)
	if err != nil {
		fmt.Println(err)
		returnErrorResponse(response, request)
	}

	jsonResponse, jsonError := json.Marshal(res)
	if jsonError != nil {
		fmt.Println(jsonError)
		returnErrorResponse(response, request)
	}

	if jsonResponse == nil {
		returnErrorResponse(response, request)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request) {
	jsonResponse, err := json.Marshal("There's no data available")
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusInternalServerError)
	response.Write(jsonResponse)
}
