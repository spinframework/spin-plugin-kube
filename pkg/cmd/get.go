package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var getCmd = &cobra.Command{
	Use:    "get <name>",
	Short:  "Display detailed application information",
	Hidden: isExperimentalFlagNotSet,
	RunE: func(_ *cobra.Command, args []string) error {
		var appName string
		if len(args) > 0 {
			appName = args[0]
		}

		if appName == "" && appNameFromCurrentDirContext != "" {
			appName = appNameFromCurrentDirContext
		}

		okey := client.ObjectKey{
			Namespace: namespace,
			Name:      appName,
		}

		app, err := kubeImpl.GetSpinApp(context.TODO(), okey)
		if err != nil {
			return err
		}

		return printApps(os.Stdout, app)
	},
}

func init() {
	configFlags.AddFlags(getCmd.Flags())
	rootCmd.AddCommand(getCmd)
}
