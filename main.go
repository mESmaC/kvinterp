package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DBConnection struct {
	IP   string
	Port string
}

func main() {
	var prompt string = "KVShell> "
	reader := bufio.NewReader(os.Stdin)

	dbConnection := DBConnection{
		IP:   "localhost",
		Port: "8080",
	}

	fmt.Print("\033[H\033[2J")

	for {
		fmt.Printf("%s", prompt)
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error")
			continue
		}

		cmdString = strings.TrimSpace(cmdString)
		parts := strings.Split(cmdString, " ")

		switch parts[0] {
		case "clear":
			fmt.Print("\033[H\033[2J")
		case "connect_db":
			if len(parts) == 2 {
				address := parts[1]
				addrParts := strings.Split(address, ":")
				if len(addrParts) == 2 {
					dbConnection.IP = addrParts[0]
					dbConnection.Port = addrParts[1]
				} else {
					fmt.Println("Invalid address format. Use IP:PORT")
					continue
				}
			} else {
				dbConnection.IP = "localhost"
				dbConnection.Port = "8080"
			}
			fmt.Printf("Connecting to DB at: %s:%s\n", dbConnection.IP, dbConnection.Port)

		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  connect_db [IP:PORT]  - Connect to a database (default: localhost:8080)")
			fmt.Println("  exit                 - Exit the shell")
			fmt.Println("  help                 - Show available commands")

		case "exit":
			return

		default:
			fmt.Printf("Unknown command: '%s', try 'help'\n", parts[0])
		}
	}
}
