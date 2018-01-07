package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/codingconcepts/hf/command"
	"github.com/spf13/cobra"
)

func main() {
	addHostCmd := &cobra.Command{
		Use:   "add",
		Short: "Adds a host entry to the hosts file",
		Args:  cobra.ExactArgs(2),
		Run:   addHost,
	}
	removeHostCmd := &cobra.Command{
		Use:   "remove",
		Short: "Removes a host entry from the hosts file",
		Args:  cobra.ExactArgs(2),
		Run:   removeHost,
	}
	flushCommand := &cobra.Command{
		Use:   "flush",
		Short: "Flushes dns cache",
		Run:   command.FlushDNS,
	}

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(addHostCmd, removeHostCmd, flushCommand)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error executing command: %v", err)
	}
}

func addHost(c *cobra.Command, args []string) {
	ip := args[0]
	host := args[1]

	file, err := os.OpenFile(command.HostFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalf("opening file: %v", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Fatalf("closing file: %v", err)
		}
	}()

	if _, err := file.Write(hostLine(ip, host)); err != nil {
		log.Fatalf("writing file: %v", err)
	}
}

func removeHost(c *cobra.Command, args []string) {
	ip := args[0]
	host := args[1]
	searchLine := hostLine(ip, host)

	fileBytes, err := ioutil.ReadFile(command.HostFile)
	if err != nil {
		log.Fatalf("reading file: %v", err)
	}

	fileBytes = bytes.Replace(fileBytes, searchLine, []byte{}, 1)
	if err = ioutil.WriteFile(command.HostFile, fileBytes, 0660); err != nil {
		log.Fatalf("writing file: %v", err)
	}
}

func hostLine(ip string, host string) []byte {
	return []byte(fmt.Sprintf("\n%s %s", ip, host))
}
