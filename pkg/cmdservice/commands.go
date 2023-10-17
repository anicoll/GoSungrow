package cmdservice

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/anicoll/gosungrow/pkg/cmdexec"
	"github.com/anicoll/gosungrow/pkg/cmdhelp"
	"github.com/anicoll/gosungrow/pkg/cmdlog"
	"github.com/anicoll/gosungrow/pkg/only"
	"github.com/kardianos/service"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

const (
	Group            = "Run"
	ServiceStart     = "start"
	ServiceStop      = "stop"
	ServiceRestart   = "restart"
	ServiceInstall   = "install"
	ServiceUninstall = "uninstall"
	ServiceList      = "list"
)

func (c *Service) AttachCommands(cmd *cobra.Command) *cobra.Command {
	for range only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		name := aurora.White(c.name).Bold() // c.cmd.Name()

		// ******************************************************************************** //
		c.SelfCmd = &cobra.Command{
			Use:                   "service",
			Aliases:               []string{},
			Short:                 fmt.Sprintf("Run %s as a system service.", name),
			Long:                  fmt.Sprintf("Run %s as a system service.", name),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE:                  c.CmdService,
			Args:                  cobra.ExactArgs(1),
		}
		c.SelfCmd.Example = cmdhelp.PrintExamples(c.SelfCmd, ServiceInstall, ServiceUninstall, ServiceStart, ServiceStop, ServiceRestart, ServiceList)
		c.SelfCmd.Annotations = map[string]string{"group": Group}
		cmd.AddCommand(c.SelfCmd)

		// ******************************************************************************** //
		cmdServiceLoad := &cobra.Command{
			Use:                   ServiceInstall,
			Aliases:               []string{"add", "load"},
			Short:                 fmt.Sprintf("Load %s as a system service.", name),
			Long:                  fmt.Sprintf("Load %s as a system service.", name),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return c.CmdServiceCmd(cmd, []string{ServiceInstall})
			},
			Args: cobra.ExactArgs(0),
		}
		cmdServiceLoad.Example = cmdhelp.PrintExamples(cmdServiceLoad, "")
		cmdServiceLoad.Annotations = map[string]string{"group": Group}
		c.SelfCmd.AddCommand(cmdServiceLoad)

		// ******************************************************************************** //
		cmdServiceUnload := &cobra.Command{
			Use:                   ServiceUninstall,
			Aliases:               []string{"remove", "unload"},
			Short:                 fmt.Sprintf("Unload %s as a system service.", name),
			Long:                  fmt.Sprintf("Unload %s as a system service.", name),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return c.CmdServiceCmd(cmd, []string{ServiceUninstall})
			},
			Args: cobra.ExactArgs(0),
		}
		cmdServiceUnload.Example = cmdhelp.PrintExamples(cmdServiceUnload, "")
		cmdServiceUnload.Annotations = map[string]string{"group": Group}
		c.SelfCmd.AddCommand(cmdServiceUnload)

		// ******************************************************************************** //
		cmdServiceStart := &cobra.Command{
			Use:                   ServiceStart,
			Aliases:               []string{},
			Short:                 fmt.Sprintf("Start the %s system service.", name),
			Long:                  fmt.Sprintf("Start the %s system service.", name),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return c.CmdServiceCmd(cmd, []string{ServiceStart})
			},
			Args: cobra.ExactArgs(0),
		}
		cmdServiceStart.Example = cmdhelp.PrintExamples(cmdServiceStart, "")
		cmdServiceStart.Annotations = map[string]string{"group": Group}
		c.SelfCmd.AddCommand(cmdServiceStart)

		// ******************************************************************************** //
		cmdServiceStop := &cobra.Command{
			Use:                   ServiceStop,
			Aliases:               []string{},
			Short:                 fmt.Sprintf("Stop the %s system service.", name),
			Long:                  fmt.Sprintf("Stop the %s system service.", name),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return c.CmdServiceCmd(cmd, []string{ServiceStop})
			},
			Args: cobra.ExactArgs(0),
		}
		cmdServiceStop.Example = cmdhelp.PrintExamples(cmdServiceStop, "")
		cmdServiceStop.Annotations = map[string]string{"group": Group}
		c.SelfCmd.AddCommand(cmdServiceStop)

		// ******************************************************************************** //
		cmdServiceRestart := &cobra.Command{
			Use:                   ServiceRestart,
			Aliases:               []string{},
			Short:                 fmt.Sprintf("Restart the %s system service.", name),
			Long:                  fmt.Sprintf("Restart the %s system service.", name),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return c.CmdServiceCmd(cmd, []string{ServiceRestart})
			},
			Args: cobra.ExactArgs(0),
		}
		cmdServiceRestart.Example = cmdhelp.PrintExamples(cmdServiceRestart, "")
		cmdServiceRestart.Annotations = map[string]string{"group": Group}
		c.SelfCmd.AddCommand(cmdServiceRestart)

		// ******************************************************************************** //
		cmdServiceList := &cobra.Command{
			Use:                   ServiceList,
			Aliases:               []string{"show"},
			Short:                 fmt.Sprintf("List %s system service.", name),
			Long:                  fmt.Sprintf("List %s system service.", name),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return c.CmdServiceCmd(cmd, []string{ServiceList})
			},
			Args: cobra.ExactArgs(0),
		}
		cmdServiceList.Example = cmdhelp.PrintExamples(cmdServiceList, "list")
		cmdServiceList.Annotations = map[string]string{"group": Group}
		c.SelfCmd.AddCommand(cmdServiceList)
	}

	return c.SelfCmd
}

func (c *Service) InitArgs(_ *cobra.Command, _ []string) error {
	var err error
	for range only.Once {
		c.service, c.Error = service.New(c.program, &c.config)
		if c.Error != nil {
			break
		}
	}
	return err
}

func (c *Service) CmdService(cmd *cobra.Command, args []string) error {
	for range only.Once {
		if len(args) == 0 {
			c.Error = cmd.Help()
			break
		}
	}

	return c.Error
}

func (c *Service) CmdServiceLoad(_ *cobra.Command, args []string) error {
	for range only.Once {
		// */1 * * * * /dir/command args args
		cronString := strings.Join(args[0:5], " ")
		cronString = strings.ReplaceAll(cronString, ".", "*")
		cmdexec.ResetArgs(args[5:]...)

		cmdlog.Printf("Created job schedule using '%s'\n", cronString)
		cmdlog.Printf("Job command '%s'\n", strings.Join(os.Args, " "))

		if c.Error != nil {
			break
		}
	}

	return c.Error
}

func (c *Service) CmdServiceCmd(_ *cobra.Command, args []string) error {
	for range only.Once {
		if len(args) == 0 {
			break
		}

		switch strings.ToLower(args[0]) {
		case ServiceStart:
			c.Error = service.Control(c.service, ServiceStart)
			if c.Error != nil {
				break
			}
			fmt.Printf("Started the %s system service.\n", c.name)

		case ServiceStop:
			c.Error = service.Control(c.service, ServiceStop)
			if c.Error != nil {
				break
			}
			fmt.Printf("Stopped the %s system service.\n", c.name)

		case ServiceRestart:
			c.Error = service.Control(c.service, ServiceRestart)
			if c.Error != nil {
				break
			}
			fmt.Printf("Restarting the %s service.\n", c.name)

		case ServiceInstall:
			c.Error = service.Control(c.service, ServiceInstall)
			if c.Error != nil {
				break
			}
			fmt.Printf("Installed %s as a system service.\n", c.name)

		case ServiceUninstall:
			c.Error = service.Control(c.service, ServiceUninstall)
			if c.Error != nil {
				break
			}
			fmt.Printf("Uninstalled %s as a system service.\n", c.name)

		case ServiceList:
			c.Error = service.Control(c.service, ServiceUninstall)
			if c.Error != nil {
				break
			}
			fmt.Printf("Uninstalled %s as a system service.\n", c.name)

		default:
			c.Error = errors.New("unknown sub-command")
		}
	}

	return c.Error
}

func (c *Service) CmdServiceList(_ *cobra.Command, _ []string) error {
	for range only.Once {
		fmt.Println("CmdServiceList() Not yet implemented.")
	}

	return c.Error
}
