package getReportPsTree

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/reportService/getReportPsTree"
	Disabled     = false
	EndPointName = "WebAppService.getReportPsTree"
)

type RequestData struct {
	DeviceType valueTypes.Integer `json:"device_type" required:"true"`
	PsId       valueTypes.PsId    `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableSortOn:"PsKey"`

	PsId          valueTypes.Integer `json:"ps_id"`
	PsKey         valueTypes.String  `json:"ps_key"`
	DeviceType    valueTypes.Integer `json:"device_type"`
	PsName        valueTypes.String  `json:"ps_name"`
	Id            valueTypes.Integer `json:"id"`
	Pid           valueTypes.Integer `json:"pid"`
	Name          valueTypes.String  `json:"name"`
	DeviceName    valueTypes.String  `json:"device_name"`
	IsParent      valueTypes.Bool    `json:"isparent" PointId:"is_parent"`
	UUIDIndexCode valueTypes.String  `json:"uuid_index_code"`
	RowCount      valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range only.Once {
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
