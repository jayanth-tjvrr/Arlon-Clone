package clusterspec

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:               "clusterspec",
		Short:             "Manage cluster specifications",
		Long:              "Manage cluster specifications",
		DisableAutoGenTag: true,
		Run: func(c *cobra.Command, args []string) {
			c.Usage()
		},
	}
	command.AddCommand(ListClusterspecsCommand())
	command.AddCommand(CreateClusterspecCommand())
	command.AddCommand(UpdateClusterspecCommand())
	command.AddCommand(DeleteClusterspecCommand())
	return command
}
