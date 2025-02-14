//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package cloudstack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type CreateIpForwardingRuleParams struct {
	p map[string]interface{}
}

func (p *CreateIpForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := p.p["openfirewall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("openfirewall", vv)
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	return u
}

func (p *CreateIpForwardingRuleParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *CreateIpForwardingRuleParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
}

func (p *CreateIpForwardingRuleParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
}

func (p *CreateIpForwardingRuleParams) SetOpenfirewall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["openfirewall"] = v
}

func (p *CreateIpForwardingRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *CreateIpForwardingRuleParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
}

// You should always use this function to get a new CreateIpForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *NATService) NewCreateIpForwardingRuleParams(ipaddressid string, protocol string, startport int) *CreateIpForwardingRuleParams {
	p := &CreateIpForwardingRuleParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddressid"] = ipaddressid
	p.p["protocol"] = protocol
	p.p["startport"] = startport
	return p
}

// Creates an IP forwarding rule
func (s *NATService) CreateIpForwardingRule(p *CreateIpForwardingRuleParams) (*CreateIpForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("createIpForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateIpForwardingRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateIpForwardingRuleResponse struct {
	Cidrlist                  string `json:"cidrlist"`
	Fordisplay                bool   `json:"fordisplay"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Ipaddressid               string `json:"ipaddressid"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Privateendport            string `json:"privateendport"`
	Privateport               string `json:"privateport"`
	Protocol                  string `json:"protocol"`
	Publicendport             string `json:"publicendport"`
	Publicport                string `json:"publicport"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vmguestip                 string `json:"vmguestip"`
}

type DeleteIpForwardingRuleParams struct {
	p map[string]interface{}
}

func (p *DeleteIpForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteIpForwardingRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteIpForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *NATService) NewDeleteIpForwardingRuleParams(id string) *DeleteIpForwardingRuleParams {
	p := &DeleteIpForwardingRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an IP forwarding rule
func (s *NATService) DeleteIpForwardingRule(p *DeleteIpForwardingRuleParams) (*DeleteIpForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteIpForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteIpForwardingRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteIpForwardingRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DisableStaticNatParams struct {
	p map[string]interface{}
}

func (p *DisableStaticNatParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	return u
}

func (p *DisableStaticNatParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
}

// You should always use this function to get a new DisableStaticNatParams instance,
// as then you are sure you have configured all required params
func (s *NATService) NewDisableStaticNatParams(ipaddressid string) *DisableStaticNatParams {
	p := &DisableStaticNatParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddressid"] = ipaddressid
	return p
}

// Disables static rule for given IP address
func (s *NATService) DisableStaticNat(p *DisableStaticNatParams) (*DisableStaticNatResponse, error) {
	resp, err := s.cs.newRequest("disableStaticNat", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableStaticNatResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DisableStaticNatResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type EnableStaticNatParams struct {
	p map[string]interface{}
}

func (p *EnableStaticNatParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["vmguestip"]; found {
		u.Set("vmguestip", v.(string))
	}
	return u
}

func (p *EnableStaticNatParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
}

func (p *EnableStaticNatParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *EnableStaticNatParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *EnableStaticNatParams) SetVmguestip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmguestip"] = v
}

// You should always use this function to get a new EnableStaticNatParams instance,
// as then you are sure you have configured all required params
func (s *NATService) NewEnableStaticNatParams(ipaddressid string, virtualmachineid string) *EnableStaticNatParams {
	p := &EnableStaticNatParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddressid"] = ipaddressid
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Enables static NAT for given IP address
func (s *NATService) EnableStaticNat(p *EnableStaticNatParams) (*EnableStaticNatResponse, error) {
	resp, err := s.cs.newRequest("enableStaticNat", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableStaticNatResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type EnableStaticNatResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *EnableStaticNatResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias EnableStaticNatResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListIpForwardingRulesParams struct {
	p map[string]interface{}
}

func (p *ListIpForwardingRulesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *ListIpForwardingRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListIpForwardingRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListIpForwardingRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListIpForwardingRulesParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
}

func (p *ListIpForwardingRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListIpForwardingRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListIpForwardingRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListIpForwardingRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListIpForwardingRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListIpForwardingRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListIpForwardingRulesParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

// You should always use this function to get a new ListIpForwardingRulesParams instance,
// as then you are sure you have configured all required params
func (s *NATService) NewListIpForwardingRulesParams() *ListIpForwardingRulesParams {
	p := &ListIpForwardingRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NATService) GetIpForwardingRuleByID(id string, opts ...OptionFunc) (*IpForwardingRule, int, error) {
	p := &ListIpForwardingRulesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListIpForwardingRules(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.IpForwardingRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for IpForwardingRule UUID: %s!", id)
}

// List the IP forwarding rules
func (s *NATService) ListIpForwardingRules(p *ListIpForwardingRulesParams) (*ListIpForwardingRulesResponse, error) {
	resp, err := s.cs.newRequest("listIpForwardingRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListIpForwardingRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListIpForwardingRulesResponse struct {
	Count             int                 `json:"count"`
	IpForwardingRules []*IpForwardingRule `json:"ipforwardingrule"`
}

type IpForwardingRule struct {
	Cidrlist                  string `json:"cidrlist"`
	Fordisplay                bool   `json:"fordisplay"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Ipaddressid               string `json:"ipaddressid"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Privateendport            string `json:"privateendport"`
	Privateport               string `json:"privateport"`
	Protocol                  string `json:"protocol"`
	Publicendport             string `json:"publicendport"`
	Publicport                string `json:"publicport"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vmguestip                 string `json:"vmguestip"`
}
