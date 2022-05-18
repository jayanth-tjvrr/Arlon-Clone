package main

import (
	"flag"
	"os"

	"arlon.io/arlon/cmd/bundle"
	"arlon.io/arlon/cmd/cluster"
	"arlon.io/arlon/cmd/clusterspec"
	"arlon.io/arlon/cmd/controller"
	"arlon.io/arlon/cmd/list_clusters"
	"arlon.io/arlon/cmd/profile"
	"github.com/spf13/cobra"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	//+kubebuilder:scaffold:imports
)

func main() {
	command := &cobra.Command{
		Use:               "arlon",
		Short:             "Run the Arlon program",
		Long:              "Run the Arlon program",
		DisableAutoGenTag: true,
		Run: func(c *cobra.Command, args []string) {
			c.Println(c.UsageString())
		},
	}
	// don't display usage upon error
	command.SilenceUsage = true
	command.AddCommand(controller.NewCommand())
	command.AddCommand(list_clusters.NewCommand())
	command.AddCommand(bundle.NewCommand())
	command.AddCommand(profile.NewCommand())
	command.AddCommand(clusterspec.NewCommand())
	command.AddCommand(cluster.NewCommand())

	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	// override default log level, which is initially set to 'debug'
	flag.Set("zap-log-level", "info")
	flag.Parse()
	logger := zap.New(zap.UseFlagOptions(&opts))
	ctrl.SetLogger(logger)
	args := flag.Args()
	command.SetArgs(args)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
