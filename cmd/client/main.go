package main

import (
	client2 "DHT/pkg/client"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// printHelp prints the help message.
func printHelp() {
	fmt.Println("Help:")
	fmt.Println("- help: print this help message.")
	fmt.Println("- get <key:str>: get the value for the given key.")
	fmt.Println("- put <key:str> <value:str> <ttl:int> <replication:int>: put the key value pair.")
	fmt.Println("- exit: exit the client.")
	fmt.Println()
}

// clientShell run a interactive shell.
func clientShell(address string) {
	printHelp()
	reader := bufio.NewReader(os.Stdin)
	client := client2.NewClient(address)
	if client == nil {
		log.Fatal("Client cannot connect to the address")
	}
	defer client.Close()
	for true {
		fmt.Print("Client> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		fields := strings.Fields(text)
		switch fields[0] {
		case "get":
			if len(fields) != 2 {
				fmt.Println("Syntax: get <key:str>")
				continue
			}
			if value, ok, err := client.Get([]byte(fields[1])); err != nil {
				fmt.Println("error:", err)
			} else if ok {
				fmt.Println(string(value))
			} else {
				fmt.Println("<None>")
			}
		case "put":
			if len(fields) != 5 {
				fmt.Println("Syntax: put <key:str> <value:str> <ttl:int> <replication:int>")
				continue
			}
			key := fields[1]
			value := fields[2]
			ttl, err := strconv.Atoi(fields[3])
			if err != nil {
				fmt.Println("Syntax: put <key:str> <value:str> <ttl:int> <replication:int>")
				continue
			}
			replication, err := strconv.Atoi(fields[4])
			if err != nil {
				fmt.Println("Syntax: put <key:str> <value:str> <ttl:int> <replication:int>")
				continue
			}
			client.Put([]byte(key), []byte(value), uint16(ttl), uint8(replication))

		case "help":
			printHelp()
		case "exit":
			return
		default:
			fmt.Println("unknown command")
			printHelp()
		}
		if text == "quit" {
			break
		}
	}
}

func main() {
	var address string
	flag.StringVar(&address, "addr", "", "the address of API server to connect to")
	flag.Parse()
	if len(address) == 0 {
		log.Fatal("Please specify address by -addr, e.g. 127.0.0.1:7401")
	}
	clientShell(address)
}
