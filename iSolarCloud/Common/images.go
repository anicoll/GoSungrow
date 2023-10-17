package Common

import (
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

type Device struct {
	GoStruct GoStruct.GoStruct `json:"GoStruct" PointIdFrom:"PsId" PointIdReplace:"true" PointDeviceFrom:"PsId"`

	PsKey      valueTypes.PsKey   `json:"ps_key" PointId:"ps_key" PointUpdateFreq:"UpdateFreqBoot"`
	PsId       valueTypes.PsId    `json:"ps_id" PointId:"ps_id" PointUpdateFreq:"UpdateFreqBoot"`
	DeviceType valueTypes.Integer `json:"device_type" PointId:"device_type" PointUpdateFreq:"UpdateFreqBoot"`
	DeviceCode valueTypes.Integer `json:"device_code" PointId:"device_code" PointUpdateFreq:"UpdateFreqBoot"`
	ChannelId  valueTypes.Integer `json:"chnnl_id" PointId:"channel_id" PointUpdateFreq:"UpdateFreqBoot"`

	PsName                    valueTypes.String    `json:"ps_name"`
	PsStatus                  valueTypes.Bool      `json:"ps_status"`
	PsIsNotInit               valueTypes.Bool      `json:"ps_is_not_init"`
	IsBankPs                  valueTypes.Bool      `json:"is_bank_ps"`
	IsTuv                     valueTypes.Bool      `json:"is_tuv"`
	ValidFlag                 valueTypes.Bool      `json:"valid_flag"`
	PsType                    valueTypes.Integer   `json:"ps_type"`
	PsCode                    valueTypes.String    `json:"ps_code"`
	PsShortName               valueTypes.String    `json:"ps_short_name"`
	PsTimezone                valueTypes.String    `json:"ps_timezone"`
	PsFaultStatus             valueTypes.Integer   `json:"ps_fault_status"`
	PsHealthStatus            valueTypes.Integer   `json:"ps_health_status"`
	PsCountryId               valueTypes.Integer   `json:"ps_country_id"`
	PsHolder                  valueTypes.String    `json:"ps_holder"`
	AlarmCount                valueTypes.Count     `json:"alarm_count" PointUpdateFreq:"UpdateFreqBoot"`
	AlarmDevCount             valueTypes.Count     `json:"alarm_dev_count" PointUpdateFreq:"UpdateFreqBoot"`
	AreaId                    valueTypes.String    `json:"area_id"`
	AreaType                  valueTypes.Integer   `json:"area_type"`
	ArrearsStatus             valueTypes.Integer   `json:"arrears_status"`
	BuildDate                 valueTypes.DateTime  `json:"build_date" PointUpdateFreq:"UpdateFreqBoot" PointNameDateFormat:"2006/01/02 15:04:05"`
	ExpectInstallDate         valueTypes.DateTime  `json:"expect_install_date" PointNameDateFormat:"2006/01/02 15:04:05"`
	InstallDate               valueTypes.DateTime  `json:"install_date" PointNameDateFormat:"2006/01/02 15:04:05"`
	BuildStatus               valueTypes.Integer   `json:"build_status" PointUpdateFreq:"UpdateFreqBoot"`
	Description               valueTypes.String    `json:"description"`
	DesignCapacity            valueTypes.Float     `json:"design_capacity" PointUnitFrom:"DesignCapacityUnit"`
	DesignCapacityUnit        valueTypes.String    `json:"design_capacity_unit" PointIgnore:"true"`
	DesignCapacityVirgin      valueTypes.Float     `json:"design_capacity_virgin" PointIgnore:"true"`
	EquivalentHour            valueTypes.UnitValue `json:"equivalent_hour" PointUpdateFreq:"UpdateFreqDay"`
	FaultAlarmOfflineDevCount valueTypes.Count     `json:"fault_alarm_offline_dev_count"`
	FaultCount                valueTypes.Count     `json:"fault_count"`
	FaultDevCount             valueTypes.Count     `json:"fault_dev_count"`
	GcjLatitude               valueTypes.Float     `json:"gcj_latitude" PointUnit:"GPS"`
	GcjLongitude              valueTypes.Float     `json:"gcj_longitude" PointUnit:"GPS"`
	GprsLatitude              valueTypes.Float     `json:"gprs_latitude" PointUnit:"GPS"`
	GprsLongitude             valueTypes.Float     `json:"gprs_longitude" PointUnit:"GPS"`
	InstalledPowerMap         valueTypes.UnitValue `json:"installed_power_map"`
	InstalledPowerVirgin      valueTypes.Float     `json:"installed_power_virgin" PointIgnore:"true"`
	InstallerAlarmCount       valueTypes.Count     `json:"installer_alarm_count"`
	InstallerFaultCount       valueTypes.Count     `json:"installer_fault_count"`
	InstallerPsFaultStatus    valueTypes.Integer   `json:"installer_ps_fault_status"`
	JoinYearInitElec          valueTypes.Float     `json:"join_year_init_elec"`
	Latitude                  valueTypes.Float     `json:"latitude" PointUnit:"GPS"`
	Location                  valueTypes.String    `json:"location"`
	Longitude                 valueTypes.Float     `json:"longitude" PointUnit:"GPS"`
	MapLatitude               valueTypes.Float     `json:"map_latitude" PointUnit:"GPS"`
	MapLongitude              valueTypes.Float     `json:"map_longitude" PointUnit:"GPS"`
	MlpeFlag                  valueTypes.Bool      `json:"mlpe_flag"`
	Nmi                       valueTypes.String    `json:"nmi"`
	OfflineDevCount           valueTypes.Count     `json:"offline_dev_count"`
	OperateYear               valueTypes.String    `json:"operate_year"`
	OperationBusName          valueTypes.String    `json:"operation_bus_name"`
	OwnerAlarmCount           valueTypes.Count     `json:"owner_alarm_count"`
	OwnerFaultCount           valueTypes.Count     `json:"owner_fault_count"`
	OwnerPsFaultStatus        valueTypes.Integer   `json:"owner_ps_fault_status"`
	Producer                  valueTypes.String    `json:"producer"`
	RecordCreateTime          valueTypes.DateTime  `json:"recore_create_time" PointId:"record_create_time" PointNameDateFormat:"2006/01/02 15:04:05"`
	SafeStartDate             valueTypes.DateTime  `json:"safe_start_date" PointNameDateFormat:"2006/01/02 15:04:05"`
	ShareType                 valueTypes.Integer   `json:"share_type"`
	ShippingAddress           valueTypes.String    `json:"shipping_address"`
	ShippingZipCode           valueTypes.String    `json:"shipping_zip_code"`
	SysScheme                 valueTypes.Integer   `json:"sys_scheme"`
	WgsLatitude               valueTypes.Float     `json:"wgs_latitude" PointUnit:"GPS"`
	WgsLongitude              valueTypes.Float     `json:"wgs_longitude" PointUnit:"GPS"`
	ZipCode                   valueTypes.String    `json:"zip_code"`
	Images                    PowerStationImages   `json:"images" PointArrayFlatten:"false"`

	P83022y valueTypes.String `json:"p83022y" PointId:"p83022" PointUpdateFreq:"UpdateFreq5Mins" PointUnit:"Wh" PointVirtual:"true"`
	P83046  valueTypes.Float  `json:"p83046" PointUpdateFreq:"UpdateFreq5Mins" PointUnit:"kW" PointVirtual:"true"`
	P83048  valueTypes.Float  `json:"p83048" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true"`
	P83049  valueTypes.Float  `json:"p83049" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true"`
	P83050  valueTypes.Float  `json:"p83050" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true"`
	P83051  valueTypes.Float  `json:"p83051" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true"`
	P83054  valueTypes.Float  `json:"p83054" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true"`
	P83055  valueTypes.Float  `json:"p83055" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true"`
	P83067  valueTypes.Float  `json:"p83067" PointUpdateFreq:"UpdateFreq5Mins" PointUnit:"kW" PointVirtual:"true"`
	P83070  valueTypes.Float  `json:"p83070" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true"`
	P83076  valueTypes.Float  `json:"p83076" PointId:"_p83076" PointName:"Pv Power" PointIgnore:"true"`                  // Dupe of PvPower
	P83077  valueTypes.Float  `json:"p83077" PointId:"_p83077" PointName:"Pv Energy" PointIgnore:"true"`                 // Dupe of PvEnergy
	P83081  valueTypes.Float  `json:"p83081" PointId:"_p83081" PointName:"Es Power" PointIgnore:"true"`                  // Dupe of EsPower
	P83089  valueTypes.Float  `json:"p83089" PointId:"_p83089" PointName:"Es Discharge Energy" PointIgnore:"true"`       // Dupe of EsDischargeEnergy
	P83095  valueTypes.Float  `json:"p83095" PointId:"_p83095" PointName:"Es Total Discharge Energy" PointIgnore:"true"` // Dupe of EsTotalDischargeEnergy
	P83118  valueTypes.Float  `json:"p83118" PointId:"_p83118" PointName:"Use Energy" PointIgnore:"true"`                // Dupe of UseEnergy
	P83120  valueTypes.Float  `json:"p83120" PointId:"_p83120" PointName:"Es Energy" PointIgnore:"true"`                 // Dupe of EsEnergy
	P83127  valueTypes.Float  `json:"p83127" PointId:"_p83127" PointName:"Es Total Energy" PointIgnore:"true"`           // Dupe of EsTotalEnergy

	Co2Reduce              valueTypes.UnitValue `json:"co2_reduce" PointVirtual:"true"`
	Co2ReduceTotal         valueTypes.UnitValue `json:"co2_reduce_total" PointUpdateFreq:"UpdateFreqTotal" PointVirtual:"true"`
	CurrPower              valueTypes.UnitValue `json:"curr_power" PointVirtual:"true"`
	DailyIrradiation       valueTypes.UnitValue `json:"daily_irradiation" PointUpdateFreq:"UpdateFreqDay" PointVirtual:"true"`
	DailyIrradiationVirgin valueTypes.Float     `json:"daily_irradiation_virgin" PointIgnore:"true"`
	PvPower                valueTypes.UnitValue `json:"pv_power" PointId:"p83076" PointName:"Pv Power" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true"`
	PvEnergy               valueTypes.UnitValue `json:"pv_energy" PointId:"p83077" PointName:"Pv Energy" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true"`
	EsPower                valueTypes.UnitValue `json:"es_power" PointId:"p83081" PointName:"ES Power" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true"`
	EsDischargeEnergy      valueTypes.UnitValue `json:"es_disenergy" PointId:"p83089" PointName:"ES Discharge Energy" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true"`
	EsTotalDischargeEnergy valueTypes.UnitValue `json:"es_total_disenergy" PointId:"p83095" PointName:"ES Total Discharge Energy" PointUpdateFreq:"UpdateFreqTotal" PointVirtual:"true"`
	UseEnergy              valueTypes.UnitValue `json:"use_energy" PointId:"p83118" PointName:"Use Energy" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true"`
	EsEnergy               valueTypes.UnitValue `json:"es_energy" PointId:"p83120" PointName:"ES Energy" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true"`
	EsTotalEnergy          valueTypes.UnitValue `json:"es_total_energy" PointId:"p83127" PointName:"ES Total Energy" PointUpdateFreq:"UpdateFreqTotal" PointVirtual:"true"`

	ParamCo2    valueTypes.Float `json:"param_co2" PointVirtual:"true"`
	ParamCoal   valueTypes.Float `json:"param_coal" PointVirtual:"true"`
	ParamIncome valueTypes.Float `json:"param_income" PointVirtual:"true"`
	ParamMeter  valueTypes.Float `json:"param_meter" PointVirtual:"true"`
	ParamNox    valueTypes.Float `json:"param_nox" PointVirtual:"true"`
	ParamPowder valueTypes.Float `json:"param_powder" PointVirtual:"true"`
	ParamSo2    valueTypes.Float `json:"param_so2" PointVirtual:"true"`
	ParamTree   valueTypes.Float `json:"param_tree" PointVirtual:"true"`
	ParamWater  valueTypes.Float `json:"param_water" PointVirtual:"true"`

	PrScale                valueTypes.String    `json:"pr_scale"`
	Radiation              valueTypes.UnitValue `json:"radiation" PointVirtual:"true"`
	RadiationVirgin        valueTypes.Float     `json:"radiation_virgin" PointIgnore:"true"`
	TodayEnergy            valueTypes.UnitValue `json:"today_energy" PointUpdateFreq:"UpdateFreqDay" PointVirtual:"true"`
	TodayIncome            valueTypes.UnitValue `json:"today_income" PointUpdateFreq:"UpdateFreqDay" PointVirtual:"true"`
	TotalCapacity          valueTypes.UnitValue `json:"total_capcity" PointId:"total_capacity" PointUpdateFreq:"UpdateFreqTotal" PointVirtual:"true"`
	TotalEnergy            valueTypes.UnitValue `json:"total_energy" PointUpdateFreq:"UpdateFreqTotal" PointVirtual:"true"`
	TotalIncome            valueTypes.UnitValue `json:"total_income" PointUpdateFreq:"UpdateFreqTotal" PointVirtual:"true"`
	TotalInitCo2Accelerate valueTypes.Float     `json:"total_init_co2_accelerate" PointUpdateFreq:"UpdateFreqTotal" PointVirtual:"true"`
	TotalInitElec          valueTypes.Float     `json:"total_init_elec" PointUpdateFreq:"UpdateFreqTotal" PointVirtual:"true"`
	TotalInitProfit        valueTypes.Float     `json:"total_init_profit" PointUpdateFreq:"UpdateFreqTotal" PointVirtual:"true"`
}

// PowerStationImage - `json:"images" PointArrayFlatten:"false"`
type PowerStationImage struct {
	FileId      valueTypes.Integer `json:"file_id"`
	Id          valueTypes.Integer `json:"id"`
	PicLanguage valueTypes.Integer `json:"pic_language"`
	PicType     valueTypes.Integer `json:"pic_type"`
	PictureName valueTypes.String  `json:"picture_name"`
	PictureURL  valueTypes.String  `json:"picture_url"`
	PsId        valueTypes.PsId    `json:"ps_id"`
	PsUnitUUID  interface{}        `json:"ps_unit_uuid"`
}
type PowerStationImages []PowerStationImage

// PsDirectOrgList - `json:"ps_direct_org_list" PointArrayFlatten:"false"`
type PsDirectOrgList []struct {
	OrgId        valueTypes.Integer `json:"org_id"`
	OrgIndexCode valueTypes.String  `json:"org_index_code"`
	OrgName      valueTypes.String  `json:"org_name"`
}

// PsOrgInfo - `json:"ps_org_info" PointArrayFlatten:"false"`
type PsOrgInfo []struct {
	DealerOrgCode   valueTypes.String  `json:"dealer_org_code"`
	Installer       valueTypes.String  `json:"installer"`
	InstallerEmail  valueTypes.String  `json:"installer_email"`
	InstallerPhone  valueTypes.String  `json:"installer_phone"`
	OrgId           valueTypes.Integer `json:"org_id"`
	OrgIndexCode    valueTypes.String  `json:"org_index_code"`
	OrgName         valueTypes.String  `json:"org_name"`
	PsDealerOrgCode valueTypes.String  `json:"ps_dealer_org_code"`
	UpOrgId         valueTypes.Integer `json:"up_org_id"`
}

// SelectedOrgList - `json:"selectedOrgList" PointId:"selected_org_list" PointArrayFlatten:"false"`
type SelectedOrgList []struct {
	OrgId        valueTypes.Integer `json:"org_id"`
	OrgIndexCode valueTypes.String  `json:"org_index_code"`
	OrgName      valueTypes.String  `json:"org_name"`
}

// SnDetailList - `json:"sn_detail_list" PointArrayFlatten:"false"`
type SnDetailList []struct {
	CommunicateDeviceType     valueTypes.Integer `json:"communicate_device_type"`
	CommunicateDeviceTypeName valueTypes.String  `json:"communicate_device_type_name"`
	Id                        valueTypes.Integer `json:"id"`
	IsEnable                  valueTypes.Bool    `json:"is_enable"`
	Sn                        valueTypes.String  `json:"sn" PointName:"Serial Number"`
}

// ReportInfo - `json:"info" PointArrayFlatten:"false"`
type ReportInfo []struct {
	GoStruct GoStruct.GoStruct `json:"-" PointIdFrom:"PsId" PointIdReplace:"true" PointDeviceFrom:"PsId"`

	PsId                   valueTypes.PsId    `json:"ps_id"`
	PsName                 valueTypes.String  `json:"ps_name"`
	PsStatus               valueTypes.Bool    `json:"ps_status"`
	PsType                 valueTypes.Integer `json:"ps_type"`
	PsTypeName             valueTypes.String  `json:"ps_type_name"`
	SysScheme              valueTypes.Integer `json:"sys_scheme"`
	SysSchemeName          valueTypes.String  `json:"sys_scheme_name"`
	DesignCapacity         valueTypes.Float   `json:"design_capacity" PointUnit:"W"`
	InstallerPsFaultStatus valueTypes.Integer `json:"installer_ps_fault_status"`
	OwnerPsFaultStatus     valueTypes.Integer `json:"owner_ps_fault_status"`
	PsFaultStatus          valueTypes.Integer `json:"ps_fault_status"`
	ValidFlag              valueTypes.Bool    `json:"valid_flag"`
}

// type foo queryBatteryBoardsList.Disabled
// type fff queryDevicePointMinuteDataList.Disabled
