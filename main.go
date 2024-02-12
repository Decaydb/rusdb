package main

import (
	"fmt"

	"github.com/decaydb/rusdb/core/config"
)

func main() {
	var conf config.DBconfig
	connect, err := conf.Connect("config/config.json")
	fmt.Println(connect, err)
}
