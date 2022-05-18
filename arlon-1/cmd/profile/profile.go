package profile

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:               "profile",
		Short:             "Manage configuration profiles",
		Long:              "Manage configuration profiles",
		DisableAutoGenTag: true,
		Run: func(c *cobra.Command, args []string) {
			c.Usage()
		},
	}
	command.AddCommand(ListProfilesCommand())
	command.AddCommand(CreateProfileCommand())
	command.AddCommand(DeleteProfileCommand())
	command.AddCommand(UpdateProfileCommand())
	return command
}
