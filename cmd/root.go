package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version string
var Commit string

var replayCmd = &cobra.Command{
	Use:   "replay",
	Short: "Replays the list of commands defined in the input file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Replaying the behaviour...")

		inputList := args[0]
		if _, err := os.Stat(inputList); errors.Is(err, os.ErrNotExist) {
			fmt.Println("The file provided does not exist, exiting...")
			os.Exit(1)
		}
	},
}

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Records the list of commands defined in the input file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Recording the behaviour...")

		inputList := args[0]
		if _, err := os.Stat(inputList); errors.Is(err, os.ErrNotExist) {
			fmt.Println("The file provided does not exist, exiting...")
			os.Exit(1)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of the go rere tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Golang Record and Replay -- v", Version)
	},
}

var rootCmd = &cobra.Command{
	Use:   "grr",
	Short: "golang record and replay",
	Long:  `Go implementation of Tsoding's python "Record and Replay" (pyrere).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Golang Record and Replay -- v", Version)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
