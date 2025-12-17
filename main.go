package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	databaseFilePath := os.Args[1]
	command := os.Args[2]

	switch command {
	case ".dbinfo":
		dbFile , err := os.Open(databaseFilePath)

		if err != nil {
			fmt.Println("[ERROR] : file Open fail " + err.Error())
		}

		header := make([]byte, 120)
	default: 
		fmt.Println("[Error] : Unkonown command")
		os.Exit(1)
	}
	

}
