package main

import (
	"fmt"
	"os"
)

func main() {

	databaseFilePath := os.Args[1]
	command := os.Args[2]

	switch command {
	case ".dbinfo":
		_ , err := os.Open(databaseFilePath)

		if err != nil {
			fmt.Println("[ERROR] : file Open fail " + err.Error())
		}
	}

}
