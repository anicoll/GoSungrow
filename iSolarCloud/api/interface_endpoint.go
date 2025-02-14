package api

import (
	"time"

	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/output"
)

type EndPoint interface {
	GetArea() AreaName
	GetName() EndPointName
	GetUrl() EndPointUrl
	IsDisabled() bool
	Help() string
	IsDebug() bool

	Call() EndPoint
	SetError(string, ...interface{}) EndPoint
	GetError() error
	IsError() bool
	MarshalJSON() ([]byte, error)
	ReadDataFile() error
	WriteDataFile() error
	String() string
	GetJsonData(bool) output.Json
	// GetTableData(filter interface{}) Table
	// GetData(bool) Json

	SetRequest(ref interface{}) EndPoint // EndPointStruct
	SetRequestByJson(j output.Json) EndPoint
	RequestRef() interface{}
	GetRequestJson() output.Json
	IsRequestValid() error
	RequestString() string
	RequestFingerprint() string
	GetRequestArgNames() map[string]string

	SetResponse([]byte) EndPoint // EndPointStruct
	ResponseRef() interface{}
	GetResponseJson() output.Json
	IsResponseValid() error
	ResponseString() string

	CacheFilename() string
	SetCacheTimeout(duration time.Duration) EndPoint
	GetCacheTimeout() time.Duration

	// GetDataTable() output.Table
	GetEndPointData() DataMap
	GetEndPointResultTable() output.Table
	GetEndPointDataTables() output.Tables
	SetFilenamePrefix(format string, args ...interface{}) string
	// ResultDataRef() ResultData
}
