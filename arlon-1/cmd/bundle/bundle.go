package bundle

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:               "bundle",
		Short:             "Manage configuration bundles",
		Long:              "Manage configuration bundles",
		DisableAutoGenTag: true,
		Run: func(c *cobra.Command, args []string) {
			c.Usage()
		},
	}
	command.AddCommand(ListBundlesCommand())
	command.AddCommand(DumpBundleCommand())
	command.AddCommand(CreateBundleCommand())
	command.AddCommand(DeleteBundleCommand())
	command.AddCommand(UpdateBundleCommand())
	return command
}
