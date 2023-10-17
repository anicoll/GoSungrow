package querySysAdvancedParam

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/devService/querySysAdvancedParam"
	Disabled     = false
	EndPointName = "AppService.querySysAdvancedParam"
)

type RequestData struct {
	PsId2   valueTypes.PsId    `json:"psId" required:"true"`
	CurPage valueTypes.Integer `json:"curPage" required:"true"`
	Size    valueTypes.Integer `json:"pageSize" required:"true"`
}

func (rd RequestData) IsValid() error {
	rd.CurPage = valueTypes.SetIntegerValue(1)
	rd.Size = valueTypes.SetIntegerValue(100)
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
