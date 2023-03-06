package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/crypto/ssh"
)

func main() {
	inventory := readInventory("inventory")

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	for _, host := range inventory {
		client, err := ssh.Dial("tcp", host+":22", config)
		if err != nil {
			fmt.Printf("Could not establish SSH connection: %s\n", host)
			continue
		}
		defer client.Close()

		session, err := client.NewSession()
		if err != nil {
			fmt.Printf("Could not create SSH session: %s\n", host)
			continue
		}
		defer session.Close()

		err = session.Run("/sbin/shutdown -h now")
		if err != nil {
			fmt.Printf("Could not execute the shutdown command: %s\n", host)
			continue
		}

		fmt.Printf("Graceful shutdown initiated: %s\n", host)
	}
}

func readInventory(filename string) []string {
	inventoryFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open the inventory file: %s\n", filename)
		os.Exit(1)
	}
	defer inventoryFile.Close()

	inventoryBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not read the inventory file: %s\n", filename)
		os.Exit(1)
	}
	inventoryStr := string(inventoryBytes)
	inventory := strings.Split(inventoryStr, "\n")
	inventory = inventory[:len(inventory)-1]

	return inventory
}
