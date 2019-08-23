package cmd

import (
	"github.com/spf13/cobra"
)

type CommandManager interface {
	Execute() error
}

type CommandManagement struct {
	root *cobra.Command
}

func (cm *CommandManagement) Execute() error {
	return cm.root.Execute()
}
