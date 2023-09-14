package verify

import (
	"log"

	"github.com/cncf/presentations/pkg/verify"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "verify",
	Short: "Verifies the presentations.yaml file against a schema.",
	Long:  "Verifies that the presentations.yaml file exists and that it's contents can be marshaled as yaml.",
	Args:  cobra.OnlyValidArgs,
	RunE:  run,
}

var args struct {
	file string
}

func init() {
	flags := Cmd.Flags()

	flags.StringVar(
		&args.file,
		"file",
		"presentations.yaml",
		"File to validate.",
	)

	Cmd.RegisterFlagCompletionFunc("output-format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"yml", "yaml"}, cobra.ShellCompDirectiveDefault
	})
}

func run(cmd *cobra.Command, argv []string) error {
	if err := verify.Verify(args.file); err != nil {
		return err
	}
	log.Println("Presentations list validated")
	return nil
}
