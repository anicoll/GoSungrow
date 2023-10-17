package getSnConnectionInfo

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/commonService/getSnConnectionInfo"
	Disabled     = false
	EndPointName = "AppService.getSnConnectionInfo"
)

type RequestData struct {
	Size    valueTypes.Integer `json:"size" required:"true"`
	CurPage valueTypes.Integer `json:"curPage" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	CurPage    valueTypes.Integer `json:"curPage" PointId:"cur_page"`
	IsMore     valueTypes.Integer `json:"isMore" PointId:"is_more"`
	PageList   []interface{}      `json:"pageList" PointId:"page_list"`
	RowCount   valueTypes.Integer `json:"rowCount" PointId:"row_count"`
	Size       valueTypes.Integer `json:"size"`
	StartIndex valueTypes.Integer `json:"startIndex" PointId:"start_index"`
	TotalPage  valueTypes.Integer `json:"totalPage" PointId:"total_page"`
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
