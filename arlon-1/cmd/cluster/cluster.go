package cluster

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:               "cluster",
		Short:             "Manage clusters",
		Long:              "Manage clusters",
		DisableAutoGenTag: true,
		Run: func(c *cobra.Command, args []string) {
			c.Usage()
		},
	}
	command.AddCommand(DeployClusterCommand())
	command.AddCommand(ListClustersCommand())
	command.AddCommand(UpdateClusterCommand())
	return command
}
