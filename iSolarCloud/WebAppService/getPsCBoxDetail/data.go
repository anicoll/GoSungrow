package getPsCBoxDetail

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/devService/getPsCBoxDetail"
	Disabled     = false
	EndPointName = "WebAppService.getPsCBoxDetail"
)

type RequestData struct {
	PsId   valueTypes.PsId    `json:"ps_id" required:"true"`
	UpUUID valueTypes.Integer `json:"up_uuid" required:"true"`
}

func (rd RequestData) IsValid() error {
	rd.PsId = valueTypes.SetPsIdValue(1171348)
	rd.UpUUID = valueTypes.SetIntegerValue(1179860)

	// rd.PsId = valueTypes.SetPsIdValue(1129147)
	// rd.UpUUID = valueTypes.SetStringValue("844763")

	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct{}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
