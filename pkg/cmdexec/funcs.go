package cmdexec

import (
	"os"

	"github.com/anicoll/gosungrow/pkg/cmdconfig"
	"github.com/anicoll/gosungrow/pkg/only"
	"github.com/spf13/cobra"
)

func ResetArgs(args ...string) {
	for range only.Once {
		newArgs := []string{os.Args[0]}
		newArgs = append(newArgs, args...)
		os.Args = newArgs
	}
}

func PopArg(args []string) (string, []string) {
	if len(args) == 0 {
		return "", args
	}
	return (args)[0], (args)[1:]
}

func PopArgs(cull int, args []string) ([]string, []string) {
	if cull > len(args) {
		args = cmdconfig.FillArray(cull, args)
		return args, []string{}
	}
	if len(args) == 0 {
		return []string{}, args
	}
	return (args)[:cull], (args)[cull:]
}

func IsLastArg(args []string) bool {
	if len(args) == 0 {
		return true
	}
	return false
}

func FindRoot(cmd *cobra.Command) *cobra.Command {
	var ret *cobra.Command
	for range only.Once {
		if !cmd.HasParent() {
			ret = cmd
			break
		}

		ret = FindRoot(cmd.Parent())
	}

	return ret
}

func ReparseArgs(cmd *cobra.Command, args []string) (bool, error) {
	var last bool
	var err error
	for range only.Once {
		if len(args) == 0 {
			last = true
			break
		}

		ResetArgs(args...)

		rootCmd := FindRoot(cmd)
		// rootCmd.SetArgs(os.Args)
		err = rootCmd.Execute()
		if err != nil {
			break
		}
	}

	return last, err
}
