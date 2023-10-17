package api

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/reflection"
	"github.com/anicoll/gosungrow/pkg/only"
)

type Api struct{}

// This is used to trim the sub-packages imported under the API.
var thisPackagePath string

func init() {
	for range only.Once {
		val := reflect.ValueOf(Api{})
		if val.Kind() == reflect.Ptr {
			thisPackagePath = val.Elem().Type().PkgPath()
			break
		}
		thisPackagePath = strings.TrimSuffix(val.Type().PkgPath(), "api")
	}
}

//goland:noinspection GoUnusedExportedFunction
func AppendUrl(host string, endpoint string) *url.URL {
	var ret *url.URL
	for range only.Once {
		endpoint = fmt.Sprintf("%s%s", host, endpoint)
		ret, _ = url.Parse(endpoint)
	}
	return ret
}

func GetArea(v interface{}) AreaName {
	return AreaName(reflection.GetArea(thisPackagePath, v))
}

func GetName(v interface{}) EndPointName {
	return EndPointName(reflection.GetName(thisPackagePath, v))
}

func GetUrl(u string) *url.URL {
	var ret *url.URL
	for range only.Once {
		var err error
		ret, err = url.Parse(u)
		if err != nil {
			ret = nil
		}
	}
	return ret
}

// func GetStructKeys(ref interface{}, keys ...string) valueTypes.UnitValueMap {
// 	ret := make(valueTypes.UnitValueMap)
//
// 	for _, k := range GoStruct.GetStructKeys(ref, keys...) {
// 		// p := UnitValue { Value: k.Value, Unit: "" }
// 		p := valueTypes.SetUnitValueString(k.Value, "", "")
// 		if k.Type.Name() == "UnitValue" {
// 			// v = JsonToUnitValue(k.JsonValue).Value
// 			// u = JsonToUnitValue(k.JsonValue).Unit
// 			// p = JsonToUnitValue(k.JsonValue)
// 			// p.Value, p.Unit = DivideByThousandIfRequired(p.Value, p.Unit)
// 			err := p.UnmarshalJSON([]byte(k.JsonValue))
// 			if err != nil {
// 				continue
// 			}
// 		}
//
// 		k.JsonName = strings.TrimSuffix(k.JsonName, "_map")	// Bit of a hack, but hey... @TODO - Future self take note.
// 		ret[valueTypes.SetPointIdString(k.JsonName)] = p
// 	}
//
// 	return ret
// }
