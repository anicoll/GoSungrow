package cmdrun

import (
	"github.com/anicoll/gosungrow/pkg/cmdservice"
	"github.com/anicoll/gosungrow/pkg/only"
	"github.com/spf13/cobra"
)

type Run struct {
	name        string
	version     string
	description string
	dir         string

	CmdService *cmdservice.Service

	Error error

	cmd     *cobra.Command
	SelfCmd *cobra.Command
}

type program struct {
	exit chan struct{}
}

func New(name string, version string, description string, configDir string) *Run {
	var ret *Run

	for range only.Once {
		ret = &Run{
			name:        name,
			version:     version,
			description: description,
			dir:         configDir,
			Error:       nil,

			cmd:     nil,
			SelfCmd: nil,
		}

		ret.CmdService = cmdservice.New(name, description, configDir)
	}

	return ret
}

func (c *Run) GetCmd() *cobra.Command {
	return c.SelfCmd
}
