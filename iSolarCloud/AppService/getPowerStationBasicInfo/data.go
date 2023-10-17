package getPowerStationBasicInfo

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/powerStationService/getPowerStationBasicInfo"
	Disabled     = false
	EndPointName = "AppService.getPowerStationBasicInfo"
)

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	AreaId            interface{}         `json:"area_id"`
	ConnectType       valueTypes.Integer  `json:"connect_type"`
	County            interface{}         `json:"county"`
	CountyCode        interface{}         `json:"county_code"`
	DesignCapacity    valueTypes.Float    `json:"design_capacity"`
	Email             valueTypes.String   `json:"email"`
	EnergyScheme      interface{}         `json:"energy_scheme"`
	ExpectInstallDate valueTypes.DateTime `json:"expect_install_date" PointNameDateFormat:"DateTimeLayout"`
	FaultSendType     interface{}         `json:"fault_send_type"`
	GcjLatitude       valueTypes.Float    `json:"gcj_latitude"`
	GcjLongitude      valueTypes.Float    `json:"gcj_longitude"`
	Latitude          valueTypes.Float    `json:"latitude"`
	Longitude         valueTypes.Float    `json:"longitude"`
	Name              interface{}         `json:"name"`
	NameCode          interface{}         `json:"name_code"`
	Nation            interface{}         `json:"nation"`
	NationCode        interface{}         `json:"nation_code"`
	ParamIncome       valueTypes.Float    `json:"param_income"`
	Prov              interface{}         `json:"prov"`
	ProvCode          interface{}         `json:"prov_code"`
	PsBuildDate       valueTypes.DateTime `json:"ps_build_date" PointNameDateFormat:"DateTimeLayout"`
	PsCountryId       valueTypes.Integer  `json:"ps_country_id"`
	PsDesc            interface{}         `json:"ps_desc"`
	PsHolder          valueTypes.String   `json:"ps_holder"`
	PsId              valueTypes.PsId     `json:"ps_id"`
	PsLocation        valueTypes.String   `json:"ps_location"`
	PsName            valueTypes.String   `json:"ps_name"`
	PsType            valueTypes.Integer  `json:"ps_type"`
	RecordCreateTime  string              `json:"recore_create_time" PointId:"record_create_time"`
	ReportType        interface{}         `json:"report_type"`
	ShippingAddress   valueTypes.String   `json:"shipping_address"`
	ShippingZipCode   valueTypes.String   `json:"shipping_zip_code"`
	TimeZoneId        valueTypes.Integer  `json:"time_zone_id"`
	ValidFlag         valueTypes.Bool     `json:"valid_flag"`
	Village           interface{}         `json:"village"`
	VillageCode       interface{}         `json:"village_code"`
	WgsLatitude       valueTypes.Float    `json:"wgs_latitude"`
	WgsLongitude      valueTypes.Float    `json:"wgs_longitude"`
	ZipCode           valueTypes.String   `json:"zip_code"`
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

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
