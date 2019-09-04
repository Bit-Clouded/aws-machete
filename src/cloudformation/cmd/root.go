package cmd

import (
	"github.com/spf13/cobra"
)

func (cm *CommandManagement) rootCmdRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}

var rootCmdLong = `CloudFormation cli utility that can can do more sophisticated actions
awscli cannot, such as copy a stack, etc.`

func initRootCmd() *CommandManagement {
	cm := &CommandManagement{
		config: &config{},
	}
	cm.root = &cobra.Command{
		Use:   "cloudformation",
		Short: "CloudFormation cli utility that does things awscli cannot.",
		Long:  rootCmdLong,
		Run:   cm.rootCmdRun,
	}
	cm.initUpdateCmd()
	cm.initFlags()

	return cm
}

func (cm *CommandManagement) initFlags() {
	cmd := cm.root
	config := cm.config

	var mode string
	cmd.PersistentFlags().StringVarP(&mode, "mode", "m", "interactive", "Modes of command execution. Valid options are: noninteractive, changesetonly, dry, interactive.")
	config.mode = ParseMode(mode)

	var intHolder int
	cmd.PersistentFlags().IntVarP(&intHolder, "wait", "w", -1, "Time out in seconds to wait for the operation to complete. -1 means wait forever.")
}

var CommandManagerInstance CommandManager = initRootCmd()
