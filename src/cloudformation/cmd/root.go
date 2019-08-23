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
	cm := &CommandManagement{}
	cm.root = &cobra.Command{
		Use:   "cloudformation",
		Short: "CloudFormation cli utility that does things awscli cannot.",
		Long:  rootCmdLong,
		Run:   cm.rootCmdRun,
	}
	cm.initUpdateCmd()

	return cm
}

var CommandManagerInstance CommandManager = initRootCmd()
