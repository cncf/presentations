package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cncf/presentations/cmd/verify"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:           "presentations",
	Long:          "Command line tool for CNCF Presentations.",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	// Add the command line flags:

	root.AddCommand(verify.Cmd)

}

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	// Execute the root command:
	// root.SetArgs(os.Args[1:])
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}
