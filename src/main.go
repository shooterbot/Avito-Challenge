package main

import (
	"Avito-Challenge/src/server"
	"fmt"
)

func MakeConStr(host string, port int, dbName string, user string, password string) string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?", user, password, host, port, dbName)
}

func main() {
	var err error = nil

	address := "127.0.0.1:31337"
	conString := MakeConStr("127.0.0.1", 5432, "avito_challenge", "avito", "challenge")

	if err == nil {
		fmt.Printf("Starting server on %s\n", address)
		err = server.RunServer(address, conString)
		if err != nil {
			fmt.Print("Failed to start a server\n")
		}
	}
}
