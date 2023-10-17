package getFaultMsgListWithYYYYMM

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/faultService/getFaultMsgListWithYYYYMM"
	Disabled     = false
	EndPointName = "AppService.getFaultMsgListWithYYYYMM"
)

type RequestData struct{}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	CreateTime    valueTypes.Integer `json:"create_time"`
	FaultCode     valueTypes.String  `json:"fault_code"`
	FaultLevel    valueTypes.Integer `json:"fault_level"`
	FaultReason   valueTypes.String  `json:"fault_reason"`
	FaultType     valueTypes.Integer `json:"fault_type"`
	FaultTypeCode valueTypes.Integer `json:"fault_type_code"`
	Id            valueTypes.Integer `json:"id"`
	PsId          valueTypes.PsId    `json:"ps_id"`
	PsKey         valueTypes.PsKey   `json:"ps_key"`
	UUID          valueTypes.Integer `json:"uuid"`
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
