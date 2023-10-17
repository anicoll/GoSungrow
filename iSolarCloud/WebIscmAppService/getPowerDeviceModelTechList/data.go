package getPowerDeviceModelTechList

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/devService/getPowerDeviceModelTechList"
	Disabled     = false
	EndPointName = "WebIscmAppService.getPowerDeviceModelTechList"
)

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
	GoStructParent GoStruct.GoStructParent `json:"-" DataTable:"true"` // PointIdFrom:"CodeId" PointIdReplace:"true"`
	GoStruct       GoStruct.GoStruct       `json:"-"`                  // PointIdFrom:"CodeId" PointIdReplace:"true"`

	CodeId          valueTypes.Integer `json:"code_id"`
	CodeValue       valueTypes.String  `json:"code_value"`
	CodeName        valueTypes.String  `json:"code_name"`
	DefaultValue    interface{}        `json:"default_value"`
	TechDescription interface{}        `json:"tech_description"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
