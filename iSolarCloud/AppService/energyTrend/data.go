package energyTrend

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/powerStationService/energyTrend"
	Disabled     = false
	EndPointName = "AppService.energyTrend"
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
	Echartunit valueTypes.String   `json:"echartunit" PointId:"echart_unit"`
	EndTime    valueTypes.DateTime `json:"endTime" PointId:"end_time" PointNameDateFormat:"DateTimeLayout"`
	EnergyMap  struct {
		ValStr valueTypes.String `json:"valStr" PointId:"val_str"`
	} `json:"energyMap" PointId:"energy_map"`
	Energyunit valueTypes.String `json:"energyunit" PointId:"energy_unit"`
	PowerMap   struct {
		Dates  []valueTypes.DateTime `json:"dates" PointNameDateFormat:"DateTimeLayout"`
		Units  valueTypes.String     `json:"units"`
		ValStr valueTypes.String     `json:"valStr" PointId:"val_str"`
	} `json:"powerMap" PointId:"power_map"`
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
