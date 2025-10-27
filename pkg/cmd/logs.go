package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/cmd/logs"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

var logOpts *logs.LogsOptions

var logsCmd = &cobra.Command{
	Use:    "logs <name>",
	Short:  "Display application logs",
	Hidden: isExperimentalFlagNotSet,
	Run: func(_ *cobra.Command, args []string) {
		var appName string
		if len(args) > 0 {
			appName = args[0]
		}

		if appName == "" && appNameFromCurrentDirContext != "" {
			appName = appNameFromCurrentDirContext
		}

		reference := fmt.Sprintf("deployment/%s", appName)

		factory, streams := NewCommandFactory()
		ccmd := logs.NewCmdLogs(factory, streams)

		cmdutil.CheckErr(logOpts.Complete(factory, ccmd, []string{reference}))
		cmdutil.CheckErr(logOpts.Validate())
		cmdutil.CheckErr(logOpts.RunLogs())
	},
}

func init() {
	_, streams := NewCommandFactory()
	logOpts = logs.NewLogsOptions(streams)
	logOpts.AddFlags(logsCmd)

	configFlags.AddFlags(logsCmd.Flags())
	rootCmd.AddCommand(logsCmd)
}
