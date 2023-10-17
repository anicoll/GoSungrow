package getPowerDeviceModelTechList

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/devService/getPowerDeviceModelTechList"
	Disabled     = false
	EndPointName = "AppService.getPowerDeviceModelTechList"
)

const (
	DeviceType1  = "1"
	DeviceType3  = "3"
	DeviceType4  = "4"
	DeviceType5  = "5"
	DeviceType7  = "7"
	DeviceType11 = "11"
	DeviceType14 = "14"
	DeviceType17 = "17"
	DeviceType22 = "22"
	DeviceType23 = "23"
	DeviceType26 = "26"
	DeviceType37 = "37"
	DeviceType41 = "41"
	DeviceType43 = "43"
	DeviceType47 = "47"
)

var DeviceTypes = []string{
	DeviceType1,
	DeviceType3,
	DeviceType4,
	DeviceType5,
	DeviceType7,
	DeviceType11,
	DeviceType14,
	DeviceType17,
	DeviceType22,
	DeviceType23,
	DeviceType26,
	DeviceType37,
	DeviceType41,
	DeviceType43,
	DeviceType47,
}

type RequestData struct {
	DeviceType valueTypes.Integer `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	CodeId          valueTypes.Integer `json:"code_id"`
	CodeName        string             `json:"code_name"`
	CodeValue       string             `json:"code_value"`
	DefaultValue    interface{}        `json:"default_value"`
	TechDescription string             `json:"tech_description"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	// 	break
	// default:
	// 	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
	return err
}

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
