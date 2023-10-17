package getPsList

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/powerStationService/getPsListForWeb"
	Disabled     = false
	EndPointName = "WebAppService.getPsList"
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
	GoStructParent GoStruct.GoStructParent `json:"-" PointIdFromChild:"PsId" PointIdReplace:"true" DataTable:"true" DataTableSortOn:"PsId"`
	GoStruct       GoStruct.GoStruct       `json:"-" PointDeviceFrom:"PsId"`

	PsId               valueTypes.Integer  `json:"psid" PointId:"ps_id"`
	PsName             valueTypes.String   `json:"psname" PointId:"ps_name"`
	ArrearsStatus      valueTypes.Integer  `json:"arrears_status"`
	DesignCapacity     valueTypes.Float    `json:"design_capacity" PointUnitFrom:"DesignCapacityUnit"`
	DesignCapacityUnit valueTypes.String   `json:"design_capacity_unit" PointIgnore:"true"`
	InstallDate        valueTypes.DateTime `json:"install_date" PointNameDateFormat:"DateTimeLayout"`
	PsCode             valueTypes.String   `json:"ps_code"`
	PsCountryId        valueTypes.Integer  `json:"ps_country_id"`
	PsLocation         valueTypes.String   `json:"ps_location"`
	PsOrgName          valueTypes.String   `json:"ps_org_name"`
	PsType             valueTypes.Integer  `json:"ps_type"`
	ShareType          valueTypes.String   `json:"share_type"`

	RowCount valueTypes.Integer `json:"rowCount"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
