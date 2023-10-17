package getPointInfoPage

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/devService/getPointInfoPage"
	Disabled     = false
	EndPointName = "WebIscmAppService.getPointInfoPage"
)

type RequestData struct{}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	DeviceTypeList []struct {
		DeviceType valueTypes.Integer `json:"device_type"`
		DeviceName valueTypes.String  `json:"device_name"`
	} `json:"deviceTypeList" PointId:"device_type_list" DataTable:"true" DataTableSortOn:"DeviceType"`
	PointTypeList []struct {
		PointType valueTypes.String `json:"point_type"`
		CodeName  valueTypes.String `json:"code_name"`
	} `json:"pointTypeList" PointId:"point_type_list" DataTable:"true" DataTableSortOn:"DeviceType"`
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
