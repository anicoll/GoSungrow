package cmd

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/anicoll/gosungrow/iSolarCloud"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/login"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/output"
	"github.com/anicoll/gosungrow/pkg/cmdconfig"
	"github.com/anicoll/gosungrow/pkg/cmdhelp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagApiUrl        = "host"
	flagApiTimeout    = "timeout"
	flagApiUsername   = "user"
	flagApiPassword   = "password"
	flagApiAppKey     = "appkey"
	flagApiLastLogin  = "token-expiry"
	flagApiOutputType = "out"
	flagApiSaveFile   = "save"
	flagApiDirectory  = "dir"
)

//goland:noinspection GoNameStartsWithPackageName
type CmdApi struct {
	CmdDefault

	// iSolarCloud api
	ApiTimeout   time.Duration
	Url          string
	Username     string
	Password     string
	AppKey       string
	LastLogin    string
	ApiToken     string
	ApiTokenFile string
	OutputType   string
	SaveFile     bool
	Directory    string

	SunGrow *iSolarCloud.SunGrow
}

func NewCmdApi() *CmdApi {
	var ret *CmdApi

	var once sync.Once
	once.Do(func() {
		ret = &CmdApi{
			CmdDefault: CmdDefault{
				Error:   nil,
				cmd:     nil,
				SelfCmd: nil,
			},
			ApiTimeout:   iSolarCloud.DefaultTimeout,
			Url:          iSolarCloud.DefaultHost,
			Username:     "",
			Password:     "",
			AppKey:       iSolarCloud.DefaultApiAppKey,
			LastLogin:    "",
			ApiToken:     "",
			ApiTokenFile: "",
			OutputType:   "",
			SunGrow:      nil,
		}
	})

	return ret
}

func (c *CmdApi) AttachCommand(cmd *cobra.Command) *cobra.Command {
	var once sync.Once
	once.Do(func() {
		if cmd == nil {
			return
		}
		c.cmd = cmd

		// ******************************************************************************** //
		cmdApi := &cobra.Command{
			Use:                   "api",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Api"},
			Short:                 fmt.Sprintf("Low-level interface to the SunGrow api."),
			Long:                  fmt.Sprintf("Low-level interface to the SunGrow api."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			Run:                   c.CmdApi,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(cmdApi)
		cmdApi.Example = cmdhelp.PrintExamples(cmdApi, "get <endpoint>", "put <endpoint>")

		// ******************************************************************************** //
		cmdApiList := &cobra.Command{
			Use:                   "ls",
			Aliases:               []string{"list"},
			Annotations:           map[string]string{"group": "Api"},
			Short:                 fmt.Sprintf("List SunGrow api endpoints/areas"),
			Long:                  fmt.Sprintf("List SunGrow api endpoints/areas"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			Run:                   c.CmdApiList,
			Args:                  cobra.RangeArgs(0, 1),
		}
		cmdApi.AddCommand(cmdApiList)
		cmdApiList.Example = cmdhelp.PrintExamples(cmdApiList, "", "areas", "endpoints", "<area name>")

		// ******************************************************************************** //
		cmdApiLogin := &cobra.Command{
			Use:                   "login",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Api"},
			Short:                 fmt.Sprintf("Login to the SunGrow api."),
			Long:                  fmt.Sprintf("Login to the SunGrow api."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				c.Error = c.ApiLogin(true)
				if c.Error != nil {
					return c.Error
				}

				c.SunGrow.Auth.Print()
				return nil
			},
			Args: cobra.MinimumNArgs(0),
		}
		cmdApi.AddCommand(cmdApiLogin)
		cmdApiLogin.Example = cmdhelp.PrintExamples(cmdApiLogin, "")

		// ******************************************************************************** //
		cmdApiGet := &cobra.Command{
			Use:                   "get",
			Aliases:               []string{output.StringTypeTable},
			Annotations:           map[string]string{"group": "Api"},
			Short:                 fmt.Sprintf("Get endpoint details from the SunGrow api."),
			Long:                  fmt.Sprintf("Get endpoint details from the SunGrow api."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				c.SunGrow.SaveAsFile = false
				c.SunGrow.OutputType.SetJson()
				return c.CmdApiGet(cmd, args)
			},
			Args: cobra.MinimumNArgs(1),
		}
		cmdApi.AddCommand(cmdApiGet)
		cmdApiGet.Example = cmdhelp.PrintExamples(cmdApiGet, "[area].<endpoint>")

		// ******************************************************************************** //
		cmdApiRaw := &cobra.Command{
			Use:                   output.StringTypeRaw,
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Api"},
			Short:                 fmt.Sprintf("Raw response from the SunGrow api."),
			Long:                  fmt.Sprintf("Raw response from the SunGrow api."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				c.SunGrow.SaveAsFile = false
				c.SunGrow.OutputType.SetRaw()
				return c.CmdApiGet(cmd, args)
			},
			Args: cobra.MinimumNArgs(1),
		}
		cmdApi.AddCommand(cmdApiRaw)
		cmdApiRaw.Example = cmdhelp.PrintExamples(cmdApiRaw, "[area].<endpoint>")

		// ******************************************************************************** //
		cmdApiSave := &cobra.Command{
			Use:                   "save",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Api"},
			Short:                 fmt.Sprintf("Save the response from the SunGrow api."),
			Long:                  fmt.Sprintf("Save the response from the SunGrow api."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				c.SunGrow.SaveAsFile = true
				c.SunGrow.OutputType.SetJson()
				return c.CmdApiGet(cmd, args)
			},
			Args: cobra.MinimumNArgs(1),
		}
		cmdApi.AddCommand(cmdApiSave)
		cmdApiSave.Example = cmdhelp.PrintExamples(cmdApiSave, "[area].<endpoint>")

		// ******************************************************************************** //
		cmdApiStruct := &cobra.Command{
			Use:                   output.StringTypeStruct,
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Api"},
			Short:                 fmt.Sprintf("Show response as Go structure (debug)"),
			Long:                  fmt.Sprintf("Show response as Go structure (debug)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				// c.SunGrow.SaveAsFile = true
				c.SunGrow.OutputType.SetStruct()
				return c.CmdApiGet(cmd, args)
			},
			Args: cobra.MinimumNArgs(1),
		}
		cmdApi.AddCommand(cmdApiStruct)
		cmdApiStruct.Example = cmdhelp.PrintExamples(cmdApiStruct, "[area].<endpoint>")

		// ******************************************************************************** //
		cmdApiPut := &cobra.Command{
			Use:                   "put",
			Aliases:               []string{"write"},
			Annotations:           map[string]string{"group": "Api"},
			Short:                 fmt.Sprintf("Put details onto the SunGrow api."),
			Long:                  fmt.Sprintf("Put details onto the SunGrow api."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			Run:                   c.CmdApiPut,
			Args:                  cobra.RangeArgs(0, 1),
		}
		cmdApi.AddCommand(cmdApiPut)
		cmdApiPut.Example = cmdhelp.PrintExamples(cmdApiPut, "[area].<endpoint> <value>")
	})
	return c.SelfCmd
}

func (c *CmdApi) AttachFlags(cmd *cobra.Command, viper *viper.Viper) {
	var once sync.Once
	once.Do(func() {
		cmd.PersistentFlags().StringVarP(&c.Username, flagApiUsername, "u", "", fmt.Sprintf("SunGrow: api username."))
		viper.SetDefault(flagApiUsername, "")
		cmd.PersistentFlags().StringVarP(&c.Password, flagApiPassword, "p", "", fmt.Sprintf("SunGrow: api password."))
		viper.SetDefault(flagApiPassword, "")
		cmd.PersistentFlags().StringVarP(&c.AppKey, flagApiAppKey, "", iSolarCloud.DefaultApiAppKey, fmt.Sprintf("SunGrow: api application key."))
		viper.SetDefault(flagApiAppKey, iSolarCloud.DefaultApiAppKey)
		cmd.PersistentFlags().StringVarP(&c.Url, flagApiUrl, "", iSolarCloud.DefaultHost, fmt.Sprintf("SunGrow: Provider API URL."))
		viper.SetDefault(flagApiUrl, iSolarCloud.DefaultHost)
		// cmd.PersistentFlags().DurationVarP(&c.ApiTimeout, flagApiTimeout, "", iSolarCloud.DefaultTimeout, fmt.Sprintf("SunGrow: API timeout."))
		// viper.SetDefault(flagApiTimeout, iSolarCloud.DefaultTimeout)
		c.ApiTimeout = iSolarCloud.DefaultTimeout
		cmd.PersistentFlags().StringVar(&c.LastLogin, flagApiLastLogin, "", "SunGrow: last login.")
		viper.SetDefault(flagApiLastLogin, "")
		// _ = cmd.PersistentFlags().MarkHidden(flagApiLastLogin)

		cmd.PersistentFlags().StringVarP(&c.OutputType, flagApiOutputType, "o", "", fmt.Sprintf("Output type: 'json', 'raw', 'file'"))
		_ = cmd.PersistentFlags().MarkHidden(flagApiOutputType)
		cmd.PersistentFlags().BoolVarP(&c.SaveFile, flagApiSaveFile, "s", false, "Save output as a file.")
		viper.SetDefault(flagApiSaveFile, false)
		cmd.PersistentFlags().StringVarP(&c.Directory, flagApiDirectory, "", "", "Save output base directory.")
		viper.SetDefault(flagApiDirectory, "")
	})
}

func (ca *Cmds) SunGrowArgs(cmd *cobra.Command, args []string) error {
	var once sync.Once
	once.Do(func() {
		ca.Error = cmds.ProcessArgs(cmd, args)
		if ca.Error != nil {
			return
		}

		ca.Api.SunGrow = iSolarCloud.NewSunGro(ca.Api.Url, ca.CacheDir)
		if ca.Api.SunGrow.Error != nil {
			ca.Error = ca.Api.SunGrow.Error
			return
		}

		ca.Error = ca.Api.SunGrow.Init()
		if ca.Error != nil {
			return
		}

		ca.Api.SunGrow.SetOutputType(ca.Api.OutputType)
		ca.Api.SunGrow.SaveAsFile = ca.Api.SaveFile

		if ca.Api.AppKey == "" {
			ca.Api.AppKey = iSolarCloud.DefaultApiAppKey
		}

		ca.Error = ca.Api.ApiLogin(false)
		if ca.Error != nil {
			return
		}

		if ca.Debug {
			ca.Api.SunGrow.Auth.Print()
		}
	})

	return ca.Error
}

func (ca *Cmds) SetOutputType(cmd *cobra.Command) error {
	var err error
	var once sync.Once
	once.Do(func() {
		foo := cmd.Parent()
		ca.Api.SunGrow.SetOutputType(foo.Use)
	})

	return err
}

func (c *CmdApi) CmdApi(cmd *cobra.Command, args []string) {
	var once sync.Once
	once.Do(func() {
		if len(args) == 0 {
			c.Error = cmd.Help()
			return
		}
	})
}

func (c *CmdApi) CmdApiList(cmd *cobra.Command, args []string) {
	var once sync.Once
	once.Do(func() {
		switch {
		case len(args) == 0:
			fmt.Println("Unknown sub-command.")
			_ = cmd.Help()

		case args[0] == "endpoints":
			c.Error = c.SunGrow.ListEndpoints("")

		case args[0] == "areas":
			c.SunGrow.ListAreas()

		default:
			c.Error = c.SunGrow.ListEndpoints(args[0])
		}
	})
}

func (c *CmdApi) CmdApiGet(_ *cobra.Command, args []string) error {
	var once sync.Once
	once.Do(func() {
		args = MinimumArraySize(2, args)
		if args[0] == "all" {
			c.Error = c.SunGrow.AllCritical()
			return
		}

		ep := c.SunGrow.GetByJson(args[0], args[1])
		if ep.IsError() {
			c.Error = ep.GetError()
			return
		}

		if c.SunGrow.Error != nil {
			c.Error = c.SunGrow.Error
			return
		}

		if c.Error != nil {
			return
		}
	})

	return c.Error
}

func (c *CmdApi) CmdApiPut(_ *cobra.Command, _ []string) {
	var once sync.Once
	once.Do(func() {
		fmt.Println("Not yet implemented.")
		// args = fillArray(1, args)
		// c.Error = SunGrow.Init()
		// if c.Error != nil {
		// 	break
		// }
	})
}

func (c *CmdApi) ApiLogin(force bool) error {
	var once sync.Once
	once.Do(func() {
		if c.SunGrow == nil {
			c.Error = errors.New("sungrow instance not configured")
			return
		}

		auth := login.SunGrowAuth{
			AppKey:       c.AppKey,
			UserAccount:  c.Username,
			UserPassword: c.Password,
			TokenFile:    c.ApiTokenFile,
			Force:        force,
		}
		c.Error = c.SunGrow.Login(auth)
		if c.Error != nil {
			return
		}

		if c.SunGrow.HasTokenChanged() {
			c.LastLogin = c.SunGrow.GetLastLogin()
			c.ApiToken = c.SunGrow.GetToken()

			sf := cmds.Api.SaveFile
			cmds.Api.SaveFile = false // We don't want to lock this in the config.
			c.Error = cmds.Unify.WriteConfig()
			cmds.Api.SaveFile = sf
		}
	})
	return c.Error
}

func MinimumArraySize(count int, args []string) []string {
	var ret []string
	var once sync.Once
	once.Do(func() {
		ret = cmdconfig.FillArray(count, args)
		for i, e := range args {
			if e == "." {
				e = ""
			}
			if e == "-" {
				e = ""
			}
			ret[i] = e
		}
	})
	return ret
}
