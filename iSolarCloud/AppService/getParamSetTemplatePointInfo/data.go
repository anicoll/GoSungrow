package getParamSetTemplatePointInfo

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/devService/getParamSetTemplatePointInfo"
	Disabled     = false
	EndPointName = "AppService.getParamSetTemplatePointInfo"
)

type RequestData struct {
	UuidList valueTypes.String `json:"uuid_list" required:"true"`
	SetType  valueTypes.String `json:"set_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	fmt.Sprintf("")
	rd.UuidList = valueTypes.SetStringValue("1179860")
	rd.SetType = valueTypes.SetStringValue("1")
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
