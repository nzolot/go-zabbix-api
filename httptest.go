package zabbix

import (
//	"fmt"
)

type Step struct {
	Name    string  `json:"name"`
	No      string  `json:"no"`
	Url     string  `json:"url"`
	Headers Headers `json:"headers,omitempty"`

	// Required status code and required string
	StatusCodes string `json:"status_codes"`
	RequiredStr string `json:"required"`
}

type Steps []Step

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Headers []Header

type HttpTest struct {
	HttpTestID string  `json:"httptestid,omitempty"`
	HostID     string  `json:"hostid,omitempty"`
	Name       string  `json:"name"`
	Delay      string  `json:"delay,omitempty"`
	Retries    string  `json:"retries,omitempty"`
	Steps      Steps   `json:"steps,omitempty"`
	Headers    Headers `json:"headers,omitempty"`

	HttpTestParent Hosts `json:"hosts,omitempty"`
}

type HttpTests []HttpTest

func (api *API) HttpTestsGet(params Params) (res HttpTests, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("httptest.get", params, &res)
	return
}

func (api *API) HttpTestGetByID(id string) (res *HttpTest, err error) {
	httptests, err := api.HttpTestsGet(Params{"httptestids": id, "selectSteps": "extend"})
	if err != nil {
		return
	}

	if len(httptests) != 1 {
		e := ExpectedOneResult(len(httptests))
		err = &e
		return
	}
	res = &httptests[0]
	return
}

func (api *API) HttpTestsCreate(httptests HttpTests) (err error) {
	response, err := api.CallWithError("httptest.create", httptests)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	httptestids := result["httptestids"].([]interface{})
	for i, id := range httptestids {
		httptests[i].HttpTestID = id.(string)
	}
	return
}

func (api *API) HttpTestsUpdate(httptests HttpTests) (err error) {
	_, err = api.CallWithError("httptest.update", httptests)
	return
}

func (api *API) HttpTestsDeleteByIds(ids []string) (err error) {
	response, err := api.CallWithError("httptest.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	httptestids := result["httptestids"].([]interface{})
	if len(ids) != len(httptestids) {
		err = &ExpectedMore{len(ids), len(httptestids)}
	}
	return
}
