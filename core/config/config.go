package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DBconfig struct {
	DName string `json:"dname"`
	DUser string `json:"duser"`
	DPass string `json:"dpass"`
	DPort string `json:"dport"`
	Dhost string `json:"dhost"`
}

func (dbc DBconfig) Connect(fileName string) (*sql.DB, error) {
	file, errj := os.ReadFile(fileName)
	if errj != nil {
		return nil, errj
	}

	errj = json.Unmarshal(file, &dbc)
	if errj != nil {
		return nil, errj
	}

	connInf := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbc.Dhost, dbc.DPort, dbc.DUser, dbc.DPass, dbc.DName)

	db, err := sql.Open("postgres", connInf)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db, nil
}
