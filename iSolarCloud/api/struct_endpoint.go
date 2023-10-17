package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/output"
	"github.com/anicoll/gosungrow/pkg/only"
)

type EndPointName string

func (n EndPointName) String() string {
	return string(n)
}

type EndPointStruct struct {
	ApiRoot     Web `json:"-"`
	RawResponse []byte

	Area           AreaName     `json:"area"`
	Name           EndPointName `json:"name"`
	Url            EndPointUrl  `json:"url"`
	FileNamePrefix string
	Request        interface{} `json:"-"`
	Response       interface{} `json:"-"`
	Error          error       `json:"-"`
	Debug          bool        `json:"-"`
}

func (ep *EndPointStruct) ApiIsDebug() bool {
	return ep.Debug
}

func (ep *EndPointStruct) GetArea() AreaName {
	return ep.Area
}

func (ep *EndPointStruct) GetName() EndPointName {
	return ep.Name
}

func (ep *EndPointStruct) GetUrl() EndPointUrl {
	return ep.Url
}

func (ep *EndPointStruct) Call() output.Json {
	panic("implement me")
}

func (ep *EndPointStruct) SetRequest(ref interface{}) error {
	for range only.Once {
		if ref == nil {
			ep.Error = errors.New("endpoint has a nil request structure")
			break
		}
		ep.Request = ref
	}
	return ep.Error
}

func (ep *EndPointStruct) GetRequest() output.Json {
	panic("implement me")
}

func (ep *EndPointStruct) GetResponse() output.Json {
	panic("implement me")
}

func (ep *EndPointStruct) IsValid() error {
	var err error
	for range only.Once {
		if ep == nil {
			ep.Error = errors.New("endpoint has a nil structure")
			break
		}
		if ep.Request == nil {
			ep.Error = errors.New("endpoint has a nil request structure")
			break
		}
		if ep.Response == nil {
			ep.Error = errors.New("endpoint has a nil response structure")
			break
		}
	}
	return err
}

func (ep EndPointStruct) String() string {
	var ret string
	for range only.Once {
		if ep.Name == NullAreaName {
			break
		}

		ret += fmt.Sprintf("Area:\t%s\nEndPoint:\t%s\nUrl:\t%s\n",
			ep.Area,
			ep.Name,
			ep.Url,
		)

		foo := ep.GetRequest()
		ret += fmt.Sprintf("Request JSON:\t%s\n",
			foo,
		)

		foo = ep.GetResponse()
		ret += fmt.Sprintf("Response JSON:\t%s\n",
			foo,
		)
	}
	return ret
}

func (ep EndPointStruct) ResponseAsJson(raw bool, r interface{}) output.Json {
	var ret output.Json
	for range only.Once {
		if raw {
			ret = output.GetAsPrettyJson(r)
			break
		}
		ret = output.GetAsPrettyJson(r)
	}
	return ret
}

func (ep EndPointStruct) ApiGetRequestArgNames(req interface{}) map[string]string {
	return GoStruct.GetStructFields(req)
}

func MarshalJSON(endpoint EndPoint) ([]byte, error) {
	e := endpoint.SetError("")
	j, err := json.Marshal(&struct {
		Area     string      `json:"area"`
		EndPoint string      `json:"endpoint"`
		Host     string      `json:"api_host"`
		Url      string      `json:"endpoint_url"`
		Request  interface{} `json:"request"`
		Response interface{} `json:"response"`
	}{
		Area:     string(e.GetArea()),
		EndPoint: string(e.GetName()),
		Host:     e.GetUrl().String(),
		Url:      e.GetUrl().String(),
		Request:  e.RequestRef(),
		Response: e.ResponseRef(),
	})
	return j, err
}
