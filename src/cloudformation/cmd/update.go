package cmd

import (
	"github.com/spf13/cobra"
)

func (cm *CommandManagement) updateCmdRun(cmd *cobra.Command, args []string) {

}

var updateCmdLong = `CloudFormation cli utility that can can do more sophisticated actions
awscli cannot, such as copy a stack, etc.`

func (cm *CommandManagement) initUpdateCmd() {
	cm.root.AddCommand(&cobra.Command{
		Use:   "update",
		Short: "update",
		Long:  updateCmdLong,
		Run:   cm.rootCmdRun,
	})
}
