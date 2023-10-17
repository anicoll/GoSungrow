package cmdrun

import (
	"fmt"

	"github.com/anicoll/gosungrow/pkg/cmdhelp"
	"github.com/anicoll/gosungrow/pkg/only"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

const (
	Group = "Run"
)

func (c *Run) AttachCommands(cmd *cobra.Command) *cobra.Command {
	for range only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		name := aurora.White(c.name).Bold() // c.cmd.Name()

		// ******************************************************************************** //
		c.SelfCmd = &cobra.Command{
			Use:                   "run",
			Aliases:               []string{},
			Short:                 fmt.Sprintf("Run %s as a daemon, service, cron or shell.", name),
			Long:                  fmt.Sprintf("Run %s as a daemon, service, cron or shell.", name),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  c.CmdRun,
			Args:                  cobra.MinimumNArgs(1),
		}
		c.SelfCmd.Example = cmdhelp.PrintExamples(c.SelfCmd, "daemon", "cron", "shell", "service")
		c.SelfCmd.Annotations = map[string]string{"group": Group}
		cmd.AddCommand(c.SelfCmd)

		c.CmdService.AttachCommands(c.SelfCmd)
	}

	return c.SelfCmd
}

func (c *Run) CmdRun(cmd *cobra.Command, args []string) error {
	for range only.Once {
		if len(args) == 0 {
			c.Error = cmd.Help()
			break
		}
	}

	return c.Error
}
