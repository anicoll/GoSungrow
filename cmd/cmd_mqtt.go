package cmd

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/anicoll/gosungrow/cmdHassio"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDevicePointAttrs"
	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/output"
	"github.com/anicoll/gosungrow/pkg/cmdhelp"
	"github.com/anicoll/gosungrow/pkg/cmdlog"
	"github.com/anicoll/gosungrow/pkg/only"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	DefaultServiceName = "GoSungrow"
	DefaultServiceArea = "Roof"
	DefaultVendor      = "MickMake"
	flagMqttUsername   = "mqtt-user"
	flagMqttPassword   = "mqtt-password"
	flagMqttHost       = "mqtt-host"
	flagMqttPort       = "mqtt-port"
)

//goland:noinspection GoNameStartsWithPackageName
type CmdMqtt struct {
	CmdDefault

	// HASSIO MQTT
	Username string
	Password string
	Host     string
	Port     string

	Client    *cmdHassio.Mqtt
	endpoints MqttEndPoints
	points    getDevicePointAttrs.PointsMap
	previous  map[string]*api.DataEntries

	log                 cmdlog.Log
	optionSleepDelay    time.Duration
	optionFetchSchedule time.Duration
}

func NewCmdMqtt(logLevel string) *CmdMqtt {
	var ret *CmdMqtt

	for range only.Once {
		if logLevel == "" {
			logLevel = cmdlog.LogLevelInfoStr
		}

		ret = &CmdMqtt{
			CmdDefault: CmdDefault{
				Error:   nil,
				cmd:     nil,
				SelfCmd: nil,
			},

			log:                 cmdlog.New(logLevel),
			optionSleepDelay:    time.Second * 40, // Takes up to 40 seconds for data to come in.
			optionFetchSchedule: time.Minute * 5,
			previous:            make(map[string]*api.DataEntries, 0),
		}
	}

	return ret
}

func (c *CmdMqtt) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		// ******************************************************************************** //
		cmdRoot := &cobra.Command{
			Use:                   "mqtt",
			Aliases:               []string{""},
			Annotations:           map[string]string{"group": "MQTT"},
			Short:                 "Connect to a HASSIO broker.",
			Long:                  "Connect to a HASSIO broker.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  c.CmdMqtt,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(cmdRoot)
		cmdRoot.Example = cmdhelp.PrintExamples(cmdRoot, "run", "sync")

		// ******************************************************************************** //
		cmdMqttRun := &cobra.Command{
			Use:                   "run",
			Aliases:               []string{""},
			Annotations:           map[string]string{"group": "MQTT"},
			Short:                 "One-off sync to a HASSIO broker.",
			Long:                  "One-off sync to a HASSIO broker.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE: func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.Mqtt.MqttArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE: cmds.Mqtt.CmdMqttRun,
			Args: cobra.RangeArgs(0, 1),
		}
		cmdRoot.AddCommand(cmdMqttRun)
		cmdMqttRun.Example = cmdhelp.PrintExamples(cmdMqttRun, "")

		// ******************************************************************************** //
		cmdMqttSync := &cobra.Command{
			Use:                   "sync",
			Aliases:               []string{""},
			Annotations:           map[string]string{"group": "MQTT"},
			Short:                 "Sync to a HASSIO MQTT broker.",
			Long:                  "Sync to a HASSIO MQTT broker.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE: func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.Mqtt.MqttArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE: cmds.Mqtt.CmdMqttSync,
			Args: cobra.RangeArgs(0, 1),
		}
		cmdRoot.AddCommand(cmdMqttSync)
		cmdMqttSync.Example = cmdhelp.PrintExamples(cmdMqttSync, "", "all")
	}
	return c.SelfCmd
}

func (c *CmdMqtt) AttachFlags(cmd *cobra.Command, viper *viper.Viper) {
	for range only.Once {
		cmd.PersistentFlags().StringVarP(&c.Username, flagMqttUsername, "", "", "HASSIO: mqtt username.")
		viper.SetDefault(flagMqttUsername, "")
		cmd.PersistentFlags().StringVarP(&c.Password, flagMqttPassword, "", "", "HASSIO: mqtt password.")
		viper.SetDefault(flagMqttPassword, "")
		cmd.PersistentFlags().StringVarP(&c.Host, flagMqttHost, "", "", "HASSIO: mqtt host.")
		viper.SetDefault(flagMqttHost, "")
		cmd.PersistentFlags().StringVarP(&c.Port, flagMqttPort, "", "", "HASSIO: mqtt port.")
		viper.SetDefault(flagMqttPort, "")
	}
}

func (c *CmdMqtt) MqttArgs(_ *cobra.Command, _ []string) error {
	for range only.Once {
		c.log.Info("Connecting to MQTT HASSIO Service...\n")
		c.Client = cmdHassio.New(cmdHassio.Mqtt{
			ClientId:     DefaultServiceName,
			EntityPrefix: DefaultServiceName,
			Username:     c.Username,
			Password:     c.Password,
			Host:         c.Host,
			Port:         c.Port,
		})
		c.Error = c.Client.GetError()
		if c.Error != nil {
			break
		}

		c.log.Info("Connecting to SunGrow...\n")
		c.Client.SungrowDevices, c.Error = cmds.Api.SunGrow.GetDeviceList()
		if c.Error != nil {
			break
		}

		c.log.Info("Found SunGrow %d devices\n", len(c.Client.SungrowDevices))
		c.Client.DeviceName = DefaultServiceName
		_, c.Error = c.Client.SetDeviceConfig(
			c.Client.DeviceName, c.Client.DeviceName,
			"virtual", "virtual", "", DefaultServiceName,
			DefaultServiceArea,
		)
		if c.Error != nil {
			break
		}

		_, c.Error = c.Client.SetDeviceConfig(
			c.Client.DeviceName, c.Client.DeviceName,
			"system", "system", "", DefaultServiceName,
			DefaultServiceArea,
		)
		if c.Error != nil {
			break
		}

		for _, psId := range c.Client.SungrowDevices {
			// ca.Error = ca.Mqtt.Mqtt.SetDeviceConfig(DefaultServiceName, strconv.FormatInt(id, 10), DefaultServiceName, model[0], "Sungrow", DefaultServiceArea)
			parent := psId.PsId.String()
			if parent == psId.PsKey.Value() {
				parent = c.Client.DeviceName
			}

			_, c.Error = c.Client.SetDeviceConfig(
				DefaultServiceName, DefaultServiceName,
				psId.PsId.String(), psId.FactoryName.Value(), psId.FactoryName.Value(), psId.FactoryName.Value(),
				DefaultServiceArea,
			)
			if c.Error != nil {
				break
			}

			_, c.Error = c.Client.SetDeviceConfig(
				DefaultServiceName, parent,
				psId.PsKey.Value(), psId.DeviceName.Value(), psId.DeviceModel.Value(), psId.FactoryName.Value(),
				DefaultServiceArea,
			)
			if c.Error != nil {
				break
			}

			c.Client.SungrowPsIds[psId.PsId] = true
		}

		c.Error = c.Client.Connect()
		if c.Error != nil {
			break
		}

		c.Error = c.Options()
		if c.Error != nil {
			break
		}

		c.log.Info("Caching Sungrow metadata...\n")
		c.Error = c.GetEndPoints()
		if c.Error != nil {
			break
		}

		c.points, c.Error = cmds.Api.SunGrow.DevicePointAttrsMap("")
		if c.Error != nil {
			break
		}
		c.log.Info("Cached %d Sungrow data points...\n", len(c.points))
	}

	return c.Error
}

func (c *CmdMqtt) CmdMqtt(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}

func (c *CmdMqtt) CmdMqttRun(_ *cobra.Command, _ []string) error {
	for range only.Once {
		c.Error = c.Cron()
		if c.Error != nil {
			break
		}

		c.log.Info("Starting ticker...\n")
		c.log.Info("Fetch Schedule: %s\n", c.GetFetchSchedule())
		c.log.Info("Sleep Delay:    %s\n", c.GetSleepDelay())
		resolution := time.Second * 10
		updateCounter := int(c.optionFetchSchedule / resolution)
		timer := time.NewTicker(resolution)
		for t := range timer.C {
			if updateCounter >= 0 {
				updateCounter--
				c.log.Debug("Sleeping: %d\n", updateCounter)
				continue
			}

			updateCounter = int(c.optionFetchSchedule / resolution)
			c.log.Debug("Update: %s\n", t.String())
			c.Error = c.Cron()
			if c.Error != nil {
				break
			}
		}
	}

	return c.Error
}

func (c *CmdMqtt) CmdMqttSync(_ *cobra.Command, args []string) error {
	for range only.Once {
		// */1 * * * * /dir/command args args
		cronString := "*/5 * * * *"
		if len(args) > 0 {
			cronString = strings.Join(args[0:5], " ")
			cronString = strings.ReplaceAll(cronString, ".", "*")
		}

		cron := gocron.NewScheduler(time.UTC)
		cron = cron.Cron(cronString)
		cron = cron.SingletonMode()

		c.Error = c.Cron()
		if c.Error != nil {
			break
		}

		var job *gocron.Job
		job, c.Error = cron.Do(c.Cron)
		if c.Error != nil {
			break
		}
		job.IsRunning()

		c.log.Info("Created job schedule using '%s'\n", cronString)
		cron.StartBlocking()
		if c.Error != nil {
			break
		}
	}

	return c.Error
}

// -------------------------------------------------------------------------------- //

func (c *CmdMqtt) Cron() error {
	for range only.Once {
		if c == nil {
			c.Error = errors.New("mqtt not available")
			break
		}

		if cmds.Api.SunGrow == nil {
			c.Error = errors.New("sungrow not available")
			break
		}

		if c.Client.IsFirstRun() {
			c.Client.UnsetFirstRun()
		} else {
			c.log.Debug("Sleeping for %s...\n", c.GetSleepDelay())
			time.Sleep(c.optionSleepDelay)
		}

		newDay := false
		if c.Client.IsNewDay() {
			newDay = true
		}

		data := cmds.Api.SunGrow.NewSunGrowData()
		data.SetCacheTimeout(c.optionFetchSchedule)

		data.SetPsIds()
		if data.Error != nil {
			c.Error = data.Error
			break
		}

		data.SetEndpoints(c.endpoints.Names()...)
		c.Error = data.GetData()
		if c.Error != nil {
			break
		}

		for _, result := range data.Results {
			c.Error = c.Update(result.EndPointName.String(), result.Response.Data, newDay)
			if c.Error != nil {
				break
			}
		}

		c.Client.LastRefresh = time.Now()
	}

	if c.Error != nil {
		c.log.Error("%s\n", c.Error)
	}
	return c.Error
}

func (c *CmdMqtt) Update(endpoint string, data api.DataMap, newDay bool) error {
	for range only.Once {
		// Also getPowerStatistics, getHouseholdStoragePsReport, getPsList, getUpTimePoint,
		c.log.Info("Syncing %d entries with HASSIO from %s.\n", len(data.Map), endpoint)

		for o := range data.Map {
			refreshConfig := newDay

			entries := data.Map[o]
			r := entries.GetEntry(api.LastEntry) // Gets the last entry

			// if strings.Contains(r.EndPoint, "active") {
			// 	fmt.Printf("EMPTY[%s] -> %s\n", r.EndPoint, r.Value.String())
			// }

			if _, ok := c.previous[o]; ok {
				previous := c.previous[o].GetEntry(api.LastEntry)
				if r.Value.String() != previous.Value.String() {
					refreshConfig = true
				}
			}
			c.previous[o] = entries

			if !r.Point.Valid {
				// Any point that shouldn't be passed through to MQTT is ignored
				// - includes child points of an aggregate of several points.
				c.log.PlainInfo("-")
				c.log.Debug("Ignored: [%s] = '%s'\n", r.EndPoint, r.Value.String())
				continue
			}

			if !c.endpoints.IsOK(r) {
				continue
			}

			if !r.Value.Valid {
				// Point doesn't have a valid value.
				// Usually a float or int that cannot be converted or is empty.
				c.log.PlainInfo("?")
				c.log.Debug("Invalid: [%s] = '%s'\n", r.EndPoint, r.Value.String())
				continue
			}

			_ = c.UpdatePoint(r)
			r.Value.UnitValueFix() // @TODO - Fix this up properly

			id := r.EndPoint
			name := r.EndPoint
			if r.Point.GroupName != "" {
				id = r.EndPoint
				name = cmdHassio.JoinStringsForName(" - ", r.Parent.Key, r.Point.Id, r.Point.GroupName, r.Point.Description)
			}

			// if r.Point.Unit == "" {
			// 	r.Point.Unit = r.Point.ValueType
			// }
			// if r.Point.Unit == "Bool" {
			// 	r.Point.Unit = mmHa.LabelBinarySensor
			// }
			// if r.Point.ValueType == "Bool" {
			// 	r.Point.Unit = mmHa.LabelBinarySensor
			// }

			re := cmdHassio.EntityConfig{
				Name:       name, // mmHa.JoinStringsForName(" - ", id), // r.Point.Name, // PointName,
				SubName:    "",
				ParentId:   r.EndPoint,
				ParentName: r.Parent.Key,
				UniqueId:   r.Point.Id,
				// UniqueId:    r.Id,
				FullId: id, // string(r.FullId),	// WAS r.Point.FullId
				// FullName:    r.Point.Name,
				Units: r.Point.Unit,
				// ValueName:   r.Point.Description,
				// ValueName:   r.Point.Id,
				DeviceClass: "",
				StateClass:  r.Point.UpdateFreq,
				Value:       &r.Value,
				Point:       r.Point,
				Icon:        r.Current.PointIcon(),
				UpdateFreq:  r.Current.DataStructure.PointUpdateFreq,

				// LastReset:   "",
				// LastResetValueTemplate: "",
			}

			re.FixConfig()
			if re.LastResetValueTemplate != "" {
				re.LastReset = r.Point.WhenReset(r.Date)
			}

			// if strings.Contains(r.EndPoint, "active") {
			// 	fmt.Printf("EMPTY[%s] -> %s\n", r.EndPoint, r.Value.String())
			// }

			if refreshConfig {
				c.log.PlainInfo("C")
				c.log.Debug("Config: [%s]\n", r.EndPoint)
				c.Error = c.Client.BinarySensorPublishConfig(re)
				if c.Error != nil {
					break
				}

				c.Error = c.Client.SensorPublishConfig(re)
				if c.Error != nil {
					break
				}
			}

			c.log.PlainInfo("U")
			c.log.Debug("Update: [%s] = '%s' %s\n", r.EndPoint, r.Value.String(), r.Value.Unit())
			c.Error = c.Client.BinarySensorPublishValue(re)
			if c.Error != nil {
				break
			}

			c.Error = c.Client.SensorPublishValue(re)
			if c.Error != nil {
				break
			}
		}
		c.log.PlainInfo("\n")
	}
	return c.Error
}

func (c *CmdMqtt) GetEndPoints() error {
	for range only.Once {
		fn := filepath.Join(cmds.ConfigDir, "mqtt_endpoints.json")
		if !output.FileExists(fn) {
			c.Error = output.PlainFileWrite(fn, []byte(DefaultMqttFile), 0o644)
			if c.Error != nil {
				break
			}
		}

		c.Error = output.FileRead(fn, &c.endpoints)
		if c.Error != nil {
			break
		}

		// All := []string{ "queryDeviceList", "getPsList", "getPsDetailWithPsType", "getPsDetail", "getKpiInfo"}
		// All := []string{ "queryDeviceList", "getPsList", "getPsDetailWithPsType", "getPsDetail", "getKpiInfo"}	//, queryMutiPointDataList, getDevicePointMinuteDataList }
		// All := []string{ "WebIscmAppService.getDeviceModel" }
		for name := range c.endpoints {
			_, c.Error = c.Client.SetDeviceConfig(
				DefaultServiceName, DefaultServiceName,
				name, DefaultServiceName+"."+name, DefaultServiceName, DefaultVendor,
				DefaultServiceArea,
			)
			if c.Error != nil {
				break
			}
		}
	}
	return c.Error
}

// UpdatePoint - Set Point values to something resembling sanity based off the points metadata.
func (c *CmdMqtt) UpdatePoint(entry *api.DataEntry) error {
	for range only.Once {
		// if !c.points.Exists(entry.Point.Id) {
		// 	c.LogDebug("Point Meta: %s - Not found.\n", entry.Point.Id)
		// 	break
		// }
		p := c.points.Get(entry.Point.Id)
		if p == nil {
			c.log.Debug("Point Meta: %s - Not found.\n", entry.Point.Id)
			break
		}

		// {
		// 	fmt.Printf("entry.Point: %s - FOUND - %v\n", entry.Point.Id, p)
		// 	// fmt.Printf("\tValue   - Description:'-'\t\tUnit:'%s'\tGroupName:'-'\tValueType:'%s'\n",
		// 	// 	r.Value.UnitValue, r.Point.ValueType)
		// 	fmt.Printf("\tDescription:'%s'\tPointName:'%s' - SAME:%t\n",
		// 		entry.Point.Description, entry.Current.DataStructure.PointName, entry.Current.DataStructure.PointName == entry.Point.Description)
		// 	fmt.Printf("\tUnit:'%s'\tPointUnit:'%s' - SAME:%t\n",
		// 		entry.Point.Unit, entry.Current.DataStructure.PointUnit, entry.Current.DataStructure.PointUnit == entry.Point.Unit)
		// 	fmt.Printf("\tGroupName:'%s'\tPointGroupName:'%s' - SAME:%t\n",
		// 		entry.Point.GroupName, entry.Current.DataStructure.PointGroupName, entry.Current.DataStructure.PointGroupName == entry.Point.GroupName)
		// 	fmt.Printf("\tValueType:'%s'\tValueType:'%s' - SAME:%t\n",
		// 		entry.Point.ValueType, entry.Current.DataStructure.ValueType, entry.Current.DataStructure.ValueType == entry.Point.ValueType)
		// }

		// // If Point description matches ...
		// if p.Name.String() == entry.Point.Description {
		// 	break
		// }
		//
		// // If Point unit matches ...
		// if p.Unit.String() == entry.Point.Unit {
		// 	break
		// }

		// Unit
		if entry.Value.Unit() == "" {
			entry.Value.SetUnit(p.Unit.String())
			// fmt.Printf("[%s] -> %s\n", entry.EndPoint, entry.Value.String())
		}
		if entry.Point.Unit == "" {
			entry.Point.Unit = p.Unit.String()
			// fmt.Printf("[%s] -> %s\n", entry.EndPoint, entry.Value.String())
		}
		if entry.Point.Unit == "" {
			entry.Point.Unit = entry.Point.ValueType
		}
		// if entry.Point.Unit == "Bool" {
		// 	entry.Point.Unit = mmHa.LabelBinarySensor
		// }
		// if entry.Point.ValueType == "Bool" {
		// 	entry.Point.Unit = mmHa.LabelBinarySensor
		// }
		// if entry.Value.TypeValue == "Bool" {
		// 	entry.Value.UnitValue = mmHa.LabelBinarySensor
		// }

		// Parent
		if len(entry.Point.Parents.Map) == 0 {
		}
		if entry.Parent.Key == "" {
		}

		// GroupName
		if entry.Point.GroupName == "" {
			entry.Point.Description = p.Name.String()
			entry.Point.GroupName = p.PointGroupName
		}

		// ValueType
		if entry.Point.ValueType == "" {
			entry.Point.ValueType = p.UnitType.String()
		}
	}

	return c.Error
}

const (
	OptionLogLevel      = "loglevel"
	OptionFetchSchedule = "fetchschedule"
	OptionSleepDelay    = "sleepdelay"
	OptionServiceState  = "servicestate"
)

func (c *CmdMqtt) Options() error {
	for range only.Once {
		c.Error = c.Client.CreateOption(OptionLogLevel, "Log Level",
			c.optionFuncLogLevel,
			cmdlog.LogLevelErrorStr, cmdlog.LogLevelWarningStr, cmdlog.LogLevelInfoStr, cmdlog.LogLevelDebugStr)
		if c.Error != nil {
			break
		}
		c.Error = c.Client.SetOption(OptionLogLevel, c.log.GetLogLevel())
		if c.Error != nil {
			break
		}

		c.Error = c.Client.CreateOption(OptionFetchSchedule, "Fetch Schedule",
			c.optionFuncFetchSchedule,
			"2m", "3m", "4m", "5m", "6m", "7m", "8m", "9m", "10m")
		if c.Error != nil {
			break
		}
		c.Error = c.Client.SetOption(OptionFetchSchedule, c.GetFetchSchedule())
		if c.Error != nil {
			break
		}

		c.Error = c.Client.CreateOption(OptionSleepDelay, "Sleep Delay After Schedule",
			c.optionFuncSleepDelay,
			"0s", "10s", "20s", "30s", "40s", "50s", "60s")
		if c.Error != nil {
			break
		}
		c.Error = c.Client.SetOption(OptionSleepDelay, c.GetSleepDelay())
		if c.Error != nil {
			break
		}

		c.Error = c.Client.CreateOption(OptionServiceState, "Service State",
			c.optionFuncServiceState,
			"Run", "Restart", "Stop")
		if c.Error != nil {
			break
		}
		c.Error = c.Client.SetOption(OptionServiceState, "Run")
		if c.Error != nil {
			break
		}
	}
	return c.Error
}

func (c *CmdMqtt) optionFuncLogLevel(_ mqtt.Client, msg mqtt.Message) {
	for range only.Once {
		request := strings.ToLower(string(msg.Payload()))
		c.log.Info("Option[%s] set to '%s'\n", OptionLogLevel, request)
		c.log.SetLogLevel(request)
		c.Error = c.Client.SetOption(OptionLogLevel, request)
		if c.Error != nil {
			c.log.Error("%s\n", c.Error)
			break
		}
	}
}

func (c *CmdMqtt) optionFuncFetchSchedule(_ mqtt.Client, msg mqtt.Message) {
	for range only.Once {
		request := strings.ToLower(string(msg.Payload()))
		c.log.Info("Option[%s] set to '%s'\n", OptionFetchSchedule, request)
		c.optionFetchSchedule, c.Error = time.ParseDuration(request)
		if c.Error != nil {
			c.log.Error("%s\n", c.Error)
			break
		}

		c.Error = c.Client.SetOption(OptionFetchSchedule, c.GetFetchSchedule())
		if c.Error != nil {
			c.log.Error("%s\n", c.Error)
			break
		}

		// c.optionCacheTimeout = c.optionFetchSchedule
	}
}

func (c *CmdMqtt) GetFetchSchedule() string {
	return fmt.Sprintf("%.0fm", c.optionFetchSchedule.Minutes())
}

func (c *CmdMqtt) optionFuncSleepDelay(_ mqtt.Client, msg mqtt.Message) {
	for range only.Once {
		request := strings.ToLower(string(msg.Payload()))
		c.log.Info("Option[%s] set to '%s'\n", OptionSleepDelay, request)
		c.optionSleepDelay, c.Error = time.ParseDuration(request)
		if c.Error != nil {
			c.log.Error("%s\n", c.Error)
			break
		}

		c.Error = c.Client.SetOption(OptionSleepDelay, request)
		if c.Error != nil {
			c.log.Error("%s\n", c.Error)
			break
		}
	}
}

func (c *CmdMqtt) GetSleepDelay() string {
	return fmt.Sprintf("%.0fs", c.optionSleepDelay.Seconds())
}

func (c *CmdMqtt) optionFuncServiceState(_ mqtt.Client, msg mqtt.Message) {
	for range only.Once {
		request := strings.ToLower(string(msg.Payload()))
		c.log.Info("Option[%s] set to '%s'\n", OptionServiceState, request)
		switch request {
		case "Run":
		case "Restart":
		case "Stop":
		}

		c.Error = c.Client.SetOption(OptionServiceState, request)
		if c.Error != nil {
			c.log.Error("%s\n", c.Error)
			break
		}
	}
}

// -------------------------------------------------------------------------------- //

type (
	MqttEndPoints map[string]MqttEndPoint
	MqttEndPoint  struct {
		Include []string `json:"include"`
		Exclude []string `json:"exclude"`
	}
)

func (c *MqttEndPoints) Names() []string {
	var ret []string
	for name := range *c {
		ret = append(ret, name)
	}
	return ret
}

func (c *MqttEndPoints) IsOK(check *api.DataEntry) bool {
	var yes bool
	for range only.Once {
		field := check.Current.GetFieldPath()
		name := field.First()

		var ep MqttEndPoint
		if ep, yes = (*c)[name]; !yes {
			yes = false
			break
		}

		if len(ep.Include) == 0 {
			yes = false
			break
		}

		for _, reStr := range ep.Exclude {
			reStr = strings.ReplaceAll(reStr, `.`, `\.`)
			reStr = strings.ReplaceAll(reStr, `*`, `.*?`)
			reStr = "^" + strings.TrimPrefix(reStr, "^")
			re := regexp.MustCompile(reStr)
			if re.MatchString(check.EndPoint) {
				yes = false
				break
			}
			if re.MatchString(check.Current.FieldPath.String()) {
				yes = false
				break
			}
			if re.MatchString(check.Current.DataStructure.Endpoint.String()) {
				yes = false
				break
			}
		}

		for _, reStr := range ep.Include {
			reStr = strings.ReplaceAll(reStr, `.`, `\.`)
			reStr = strings.ReplaceAll(reStr, `*`, `.*`)
			reStr = "^" + strings.TrimPrefix(reStr, "^")
			re := regexp.MustCompile(reStr)
			if re.MatchString(check.EndPoint) {
				yes = true
				break
			}
			if re.MatchString(check.Current.FieldPath.String()) {
				yes = true
				break
			}
			if re.MatchString(check.Current.DataStructure.Endpoint.String()) {
				yes = true
				break
			}
		}
	}
	return yes
}

const DefaultMqttFile = `{
	"queryDeviceList": {
		"include": [
			"virtual.*"
		],
		"exclude": [
			"queryDeviceList.*.devices.*",
			"queryDeviceList.*.device_types.*"
		]
	},
	"getPsList": {
		"include": [
			"virtual.*"
		],
		"exclude": [
		]
	},
	"getPsDetail": {
		"include": [
			"virtual.*"
		],
		"exclude": [
		]
	}
}
`
