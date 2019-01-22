package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "strconv"
)

// JSONFirewallObjectServiceCommon contains the General parameters for Create and Update API function
type JSONFirewallObjectServiceCommon struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Protocol string `json:"protocol"`
	Comment  string `json:"comment"`
}

// JSONFirewallObjectServiceFqdn contains the FQDN parameters for Create and Update API function
type JSONFirewallObjectServiceFqdn struct {
	Fqdn string `json:"fqdn"`
}

// JSONFirewallObjectServiceIprange contains the IP Range parameters for Create and Update API function
type JSONFirewallObjectServiceIprange struct {
	Iprange string `json:"iprange"`
}

// JSONFirewallObjectService contains the parameters for Create and Update API function
type JSONFirewallObjectService struct {
	*JSONFirewallObjectServiceCommon
	*JSONFirewallObjectServiceFqdn
	*JSONFirewallObjectServiceIprange
}

// JSONCreateFirewallObjectServiceOutput contains the output results for Create API function
type JSONCreateFirewallObjectServiceOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateFirewallObjectServiceOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateFirewallObjectServiceOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateFirewallObjectService API operation for FortiOS creates a new firewall service.
// Returns the index value of the firewall service and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewal - service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallObjectService(params *JSONFirewallObjectService) (output *JSONCreateFirewallObjectServiceOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.service/custom"
	output = &JSONCreateFirewallObjectServiceOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
		if result["status"] != nil {
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get the right response")
			return
		}
		if result["http_status"] != nil {
			output.HTTPStatus = result["http_status"].(float64)
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// UpdateFirewallObjectService API operation for FortiOS updates the specified firewall service.
// Returns the index value of the firewall service and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewal - service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallObjectService(params *JSONFirewallObjectService, mkey string) (output *JSONUpdateFirewallObjectServiceOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.service/custom"
	path += "/" + mkey
	output = &JSONUpdateFirewallObjectServiceOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
		if result["status"] != nil {
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get the right response")
			return
		}
		if result["http_status"] != nil {
			output.HTTPStatus = result["http_status"].(float64)
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// DeleteFirewallObjectService API operation for FortiOS deletes the specified firewall service.
// Returns error for service API and SDK errors.
// See the firewal - service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallObjectService(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.service/custom"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// ReadFirewallObjectService API operation for FortiOS gets the firewall service
// with the specified index value.
// Returns the requested firewall service value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewal - service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallObjectService(mkey string) (output *JSONFirewallObjectService, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.service/custom"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios reading response: %s", string(body))

	j1 := JSONFirewallObjectServiceCommon{}
	j2 := JSONFirewallObjectServiceFqdn{}
	j3 := JSONFirewallObjectServiceIprange{}

	output = &JSONFirewallObjectService{
		JSONFirewallObjectServiceCommon:  &j1,
		JSONFirewallObjectServiceFqdn:    &j2,
		JSONFirewallObjectServiceIprange: &j3,
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}

		mapTmp := (result["results"].([]interface{}))[0].(map[string]interface{})

		if mapTmp == nil {
			return
		}

		if mapTmp["name"] != nil {
			output.Name = mapTmp["name"].(string)
		}
		if mapTmp["category"] != nil {
			output.Category = mapTmp["category"].(string)
		}
		if mapTmp["protocol"] != nil {
			output.Protocol = mapTmp["protocol"].(string)
		}
		if mapTmp["fqdn"] != nil {
			output.Fqdn = mapTmp["fqdn"].(string)
		}
		if mapTmp["iprange"] != nil {
			output.Iprange = mapTmp["iprange"].(string)
		}
		if mapTmp["comment"] != nil {
			output.Comment = mapTmp["comment"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}