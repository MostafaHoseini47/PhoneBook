package main

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	const description = "phonebook application"
	root := &cobra.Command{Short: description}

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

	root.AddCommand(
	//
	)

	if err := root.Execute(); err != nil {
		log.Fatal("failed to excute root command: \n %v", err)
	}
}
