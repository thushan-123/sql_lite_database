package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {

	databaseFilePath := os.Args[1]
	command := os.Args[2]

	switch command {
	case ".dbinfo":
		dbFile, err := os.Open(databaseFilePath)

		if err != nil {
			fmt.Println("[ERROR] : file Open fail " + err.Error())
		}

		header := make([]byte, 120)

		_, err = dbFile.Read(header)

		if err != nil {
			fmt.Println("[ERROR] : " + err.Error())
		}

		pageSize := binary.BigEndian.Uint16(header[16:18])
		numberOfTable := binary.BigEndian.Uint16(header[103:105])

		fmt.Println("Page Size : ", pageSize)
		fmt.Println("Number Of Table : ", numberOfTable)

	default:
		fmt.Println("[Error] : Unkonown command")
		os.Exit(1)
	}

}
