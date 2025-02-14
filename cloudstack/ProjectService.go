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

type ActivateProjectParams struct {
	p map[string]interface{}
}

func (p *ActivateProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ActivateProjectParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new ActivateProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewActivateProjectParams(id string) *ActivateProjectParams {
	p := &ActivateProjectParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Activates a project
func (s *ProjectService) ActivateProject(p *ActivateProjectParams) (*ActivateProjectResponse, error) {
	resp, err := s.cs.newRequest("activateProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ActivateProjectResponse
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

type ActivateProjectResponse struct {
	Account                   string  `json:"account"`
	Cpuavailable              string  `json:"cpuavailable"`
	Cpulimit                  string  `json:"cpulimit"`
	Cputotal                  int64   `json:"cputotal"`
	Displaytext               string  `json:"displaytext"`
	Domain                    string  `json:"domain"`
	Domainid                  string  `json:"domainid"`
	Id                        string  `json:"id"`
	Ipavailable               string  `json:"ipavailable"`
	Iplimit                   string  `json:"iplimit"`
	Iptotal                   int64   `json:"iptotal"`
	JobID                     string  `json:"jobid"`
	Jobstatus                 int     `json:"jobstatus"`
	Memoryavailable           string  `json:"memoryavailable"`
	Memorylimit               string  `json:"memorylimit"`
	Memorytotal               int64   `json:"memorytotal"`
	Name                      string  `json:"name"`
	Networkavailable          string  `json:"networkavailable"`
	Networklimit              string  `json:"networklimit"`
	Networktotal              int64   `json:"networktotal"`
	Primarystorageavailable   string  `json:"primarystorageavailable"`
	Primarystoragelimit       string  `json:"primarystoragelimit"`
	Primarystoragetotal       int64   `json:"primarystoragetotal"`
	Projectaccountname        string  `json:"projectaccountname"`
	Secondarystorageavailable string  `json:"secondarystorageavailable"`
	Secondarystoragelimit     string  `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64 `json:"secondarystoragetotal"`
	Snapshotavailable         string  `json:"snapshotavailable"`
	Snapshotlimit             string  `json:"snapshotlimit"`
	Snapshottotal             int64   `json:"snapshottotal"`
	State                     string  `json:"state"`
	Tags                      []Tags  `json:"tags"`
	Templateavailable         string  `json:"templateavailable"`
	Templatelimit             string  `json:"templatelimit"`
	Templatetotal             int64   `json:"templatetotal"`
	Vmavailable               string  `json:"vmavailable"`
	Vmlimit                   string  `json:"vmlimit"`
	Vmrunning                 int     `json:"vmrunning"`
	Vmstopped                 int     `json:"vmstopped"`
	Vmtotal                   int64   `json:"vmtotal"`
	Volumeavailable           string  `json:"volumeavailable"`
	Volumelimit               string  `json:"volumelimit"`
	Volumetotal               int64   `json:"volumetotal"`
	Vpcavailable              string  `json:"vpcavailable"`
	Vpclimit                  string  `json:"vpclimit"`
	Vpctotal                  int64   `json:"vpctotal"`
}

type CreateProjectParams struct {
	p map[string]interface{}
}

func (p *CreateProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *CreateProjectParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateProjectParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *CreateProjectParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateProjectParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

// You should always use this function to get a new CreateProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewCreateProjectParams(displaytext string, name string) *CreateProjectParams {
	p := &CreateProjectParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	return p
}

// Creates a project
func (s *ProjectService) CreateProject(p *CreateProjectParams) (*CreateProjectResponse, error) {
	resp, err := s.cs.newRequest("createProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateProjectResponse
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

type CreateProjectResponse struct {
	Account                   string  `json:"account"`
	Cpuavailable              string  `json:"cpuavailable"`
	Cpulimit                  string  `json:"cpulimit"`
	Cputotal                  int64   `json:"cputotal"`
	Displaytext               string  `json:"displaytext"`
	Domain                    string  `json:"domain"`
	Domainid                  string  `json:"domainid"`
	Id                        string  `json:"id"`
	Ipavailable               string  `json:"ipavailable"`
	Iplimit                   string  `json:"iplimit"`
	Iptotal                   int64   `json:"iptotal"`
	JobID                     string  `json:"jobid"`
	Jobstatus                 int     `json:"jobstatus"`
	Memoryavailable           string  `json:"memoryavailable"`
	Memorylimit               string  `json:"memorylimit"`
	Memorytotal               int64   `json:"memorytotal"`
	Name                      string  `json:"name"`
	Networkavailable          string  `json:"networkavailable"`
	Networklimit              string  `json:"networklimit"`
	Networktotal              int64   `json:"networktotal"`
	Primarystorageavailable   string  `json:"primarystorageavailable"`
	Primarystoragelimit       string  `json:"primarystoragelimit"`
	Primarystoragetotal       int64   `json:"primarystoragetotal"`
	Projectaccountname        string  `json:"projectaccountname"`
	Secondarystorageavailable string  `json:"secondarystorageavailable"`
	Secondarystoragelimit     string  `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64 `json:"secondarystoragetotal"`
	Snapshotavailable         string  `json:"snapshotavailable"`
	Snapshotlimit             string  `json:"snapshotlimit"`
	Snapshottotal             int64   `json:"snapshottotal"`
	State                     string  `json:"state"`
	Tags                      []Tags  `json:"tags"`
	Templateavailable         string  `json:"templateavailable"`
	Templatelimit             string  `json:"templatelimit"`
	Templatetotal             int64   `json:"templatetotal"`
	Vmavailable               string  `json:"vmavailable"`
	Vmlimit                   string  `json:"vmlimit"`
	Vmrunning                 int     `json:"vmrunning"`
	Vmstopped                 int     `json:"vmstopped"`
	Vmtotal                   int64   `json:"vmtotal"`
	Volumeavailable           string  `json:"volumeavailable"`
	Volumelimit               string  `json:"volumelimit"`
	Volumetotal               int64   `json:"volumetotal"`
	Vpcavailable              string  `json:"vpcavailable"`
	Vpclimit                  string  `json:"vpclimit"`
	Vpctotal                  int64   `json:"vpctotal"`
}

type DeleteProjectParams struct {
	p map[string]interface{}
}

func (p *DeleteProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteProjectParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteProjectParams(id string) *DeleteProjectParams {
	p := &DeleteProjectParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a project
func (s *ProjectService) DeleteProject(p *DeleteProjectParams) (*DeleteProjectResponse, error) {
	resp, err := s.cs.newRequest("deleteProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteProjectResponse
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

type DeleteProjectResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteProjectInvitationParams struct {
	p map[string]interface{}
}

func (p *DeleteProjectInvitationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteProjectInvitationParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteProjectInvitationParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteProjectInvitationParams(id string) *DeleteProjectInvitationParams {
	p := &DeleteProjectInvitationParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes project invitation
func (s *ProjectService) DeleteProjectInvitation(p *DeleteProjectInvitationParams) (*DeleteProjectInvitationResponse, error) {
	resp, err := s.cs.newRequest("deleteProjectInvitation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteProjectInvitationResponse
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

type DeleteProjectInvitationResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListProjectInvitationsParams struct {
	p map[string]interface{}
}

func (p *ListProjectInvitationsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["activeonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("activeonly", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *ListProjectInvitationsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListProjectInvitationsParams) SetActiveonly(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["activeonly"] = v
}

func (p *ListProjectInvitationsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListProjectInvitationsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListProjectInvitationsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListProjectInvitationsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListProjectInvitationsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListProjectInvitationsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListProjectInvitationsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListProjectInvitationsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListProjectInvitationsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

// You should always use this function to get a new ListProjectInvitationsParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewListProjectInvitationsParams() *ListProjectInvitationsParams {
	p := &ListProjectInvitationsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectInvitationByID(id string, opts ...OptionFunc) (*ProjectInvitation, int, error) {
	p := &ListProjectInvitationsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListProjectInvitations(p)
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
		return l.ProjectInvitations[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ProjectInvitation UUID: %s!", id)
}

// Lists project invitations and provides detailed information for listed invitations
func (s *ProjectService) ListProjectInvitations(p *ListProjectInvitationsParams) (*ListProjectInvitationsResponse, error) {
	resp, err := s.cs.newRequest("listProjectInvitations", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListProjectInvitationsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListProjectInvitationsResponse struct {
	Count              int                  `json:"count"`
	ProjectInvitations []*ProjectInvitation `json:"projectinvitation"`
}

type ProjectInvitation struct {
	Account   string `json:"account"`
	Domain    string `json:"domain"`
	Domainid  string `json:"domainid"`
	Email     string `json:"email"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Project   string `json:"project"`
	Projectid string `json:"projectid"`
	State     string `json:"state"`
}

type ListProjectsParams struct {
	p map[string]interface{}
}

func (p *ListProjectsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (p *ListProjectsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListProjectsParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *ListProjectsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListProjectsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListProjectsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListProjectsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListProjectsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListProjectsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListProjectsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListProjectsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListProjectsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListProjectsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

// You should always use this function to get a new ListProjectsParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewListProjectsParams() *ListProjectsParams {
	p := &ListProjectsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListProjectsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListProjects(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Projects[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Projects {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectByName(name string, opts ...OptionFunc) (*Project, int, error) {
	id, count, err := s.GetProjectID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetProjectByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectByID(id string, opts ...OptionFunc) (*Project, int, error) {
	p := &ListProjectsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListProjects(p)
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
		return l.Projects[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Project UUID: %s!", id)
}

// Lists projects and provides detailed information for listed projects
func (s *ProjectService) ListProjects(p *ListProjectsParams) (*ListProjectsResponse, error) {
	resp, err := s.cs.newRequest("listProjects", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListProjectsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListProjectsResponse struct {
	Count    int        `json:"count"`
	Projects []*Project `json:"project"`
}

type Project struct {
	Account                   string  `json:"account"`
	Cpuavailable              string  `json:"cpuavailable"`
	Cpulimit                  string  `json:"cpulimit"`
	Cputotal                  int64   `json:"cputotal"`
	Displaytext               string  `json:"displaytext"`
	Domain                    string  `json:"domain"`
	Domainid                  string  `json:"domainid"`
	Id                        string  `json:"id"`
	Ipavailable               string  `json:"ipavailable"`
	Iplimit                   string  `json:"iplimit"`
	Iptotal                   int64   `json:"iptotal"`
	JobID                     string  `json:"jobid"`
	Jobstatus                 int     `json:"jobstatus"`
	Memoryavailable           string  `json:"memoryavailable"`
	Memorylimit               string  `json:"memorylimit"`
	Memorytotal               int64   `json:"memorytotal"`
	Name                      string  `json:"name"`
	Networkavailable          string  `json:"networkavailable"`
	Networklimit              string  `json:"networklimit"`
	Networktotal              int64   `json:"networktotal"`
	Primarystorageavailable   string  `json:"primarystorageavailable"`
	Primarystoragelimit       string  `json:"primarystoragelimit"`
	Primarystoragetotal       int64   `json:"primarystoragetotal"`
	Projectaccountname        string  `json:"projectaccountname"`
	Secondarystorageavailable string  `json:"secondarystorageavailable"`
	Secondarystoragelimit     string  `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64 `json:"secondarystoragetotal"`
	Snapshotavailable         string  `json:"snapshotavailable"`
	Snapshotlimit             string  `json:"snapshotlimit"`
	Snapshottotal             int64   `json:"snapshottotal"`
	State                     string  `json:"state"`
	Tags                      []Tags  `json:"tags"`
	Templateavailable         string  `json:"templateavailable"`
	Templatelimit             string  `json:"templatelimit"`
	Templatetotal             int64   `json:"templatetotal"`
	Vmavailable               string  `json:"vmavailable"`
	Vmlimit                   string  `json:"vmlimit"`
	Vmrunning                 int     `json:"vmrunning"`
	Vmstopped                 int     `json:"vmstopped"`
	Vmtotal                   int64   `json:"vmtotal"`
	Volumeavailable           string  `json:"volumeavailable"`
	Volumelimit               string  `json:"volumelimit"`
	Volumetotal               int64   `json:"volumetotal"`
	Vpcavailable              string  `json:"vpcavailable"`
	Vpclimit                  string  `json:"vpclimit"`
	Vpctotal                  int64   `json:"vpctotal"`
}

type SuspendProjectParams struct {
	p map[string]interface{}
}

func (p *SuspendProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *SuspendProjectParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new SuspendProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewSuspendProjectParams(id string) *SuspendProjectParams {
	p := &SuspendProjectParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Suspends a project
func (s *ProjectService) SuspendProject(p *SuspendProjectParams) (*SuspendProjectResponse, error) {
	resp, err := s.cs.newRequest("suspendProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r SuspendProjectResponse
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

type SuspendProjectResponse struct {
	Account                   string  `json:"account"`
	Cpuavailable              string  `json:"cpuavailable"`
	Cpulimit                  string  `json:"cpulimit"`
	Cputotal                  int64   `json:"cputotal"`
	Displaytext               string  `json:"displaytext"`
	Domain                    string  `json:"domain"`
	Domainid                  string  `json:"domainid"`
	Id                        string  `json:"id"`
	Ipavailable               string  `json:"ipavailable"`
	Iplimit                   string  `json:"iplimit"`
	Iptotal                   int64   `json:"iptotal"`
	JobID                     string  `json:"jobid"`
	Jobstatus                 int     `json:"jobstatus"`
	Memoryavailable           string  `json:"memoryavailable"`
	Memorylimit               string  `json:"memorylimit"`
	Memorytotal               int64   `json:"memorytotal"`
	Name                      string  `json:"name"`
	Networkavailable          string  `json:"networkavailable"`
	Networklimit              string  `json:"networklimit"`
	Networktotal              int64   `json:"networktotal"`
	Primarystorageavailable   string  `json:"primarystorageavailable"`
	Primarystoragelimit       string  `json:"primarystoragelimit"`
	Primarystoragetotal       int64   `json:"primarystoragetotal"`
	Projectaccountname        string  `json:"projectaccountname"`
	Secondarystorageavailable string  `json:"secondarystorageavailable"`
	Secondarystoragelimit     string  `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64 `json:"secondarystoragetotal"`
	Snapshotavailable         string  `json:"snapshotavailable"`
	Snapshotlimit             string  `json:"snapshotlimit"`
	Snapshottotal             int64   `json:"snapshottotal"`
	State                     string  `json:"state"`
	Tags                      []Tags  `json:"tags"`
	Templateavailable         string  `json:"templateavailable"`
	Templatelimit             string  `json:"templatelimit"`
	Templatetotal             int64   `json:"templatetotal"`
	Vmavailable               string  `json:"vmavailable"`
	Vmlimit                   string  `json:"vmlimit"`
	Vmrunning                 int     `json:"vmrunning"`
	Vmstopped                 int     `json:"vmstopped"`
	Vmtotal                   int64   `json:"vmtotal"`
	Volumeavailable           string  `json:"volumeavailable"`
	Volumelimit               string  `json:"volumelimit"`
	Volumetotal               int64   `json:"volumetotal"`
	Vpcavailable              string  `json:"vpcavailable"`
	Vpclimit                  string  `json:"vpclimit"`
	Vpctotal                  int64   `json:"vpctotal"`
}

type UpdateProjectParams struct {
	p map[string]interface{}
}

func (p *UpdateProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateProjectParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *UpdateProjectParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateProjectParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new UpdateProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewUpdateProjectParams(id string) *UpdateProjectParams {
	p := &UpdateProjectParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a project
func (s *ProjectService) UpdateProject(p *UpdateProjectParams) (*UpdateProjectResponse, error) {
	resp, err := s.cs.newRequest("updateProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateProjectResponse
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

type UpdateProjectResponse struct {
	Account                   string  `json:"account"`
	Cpuavailable              string  `json:"cpuavailable"`
	Cpulimit                  string  `json:"cpulimit"`
	Cputotal                  int64   `json:"cputotal"`
	Displaytext               string  `json:"displaytext"`
	Domain                    string  `json:"domain"`
	Domainid                  string  `json:"domainid"`
	Id                        string  `json:"id"`
	Ipavailable               string  `json:"ipavailable"`
	Iplimit                   string  `json:"iplimit"`
	Iptotal                   int64   `json:"iptotal"`
	JobID                     string  `json:"jobid"`
	Jobstatus                 int     `json:"jobstatus"`
	Memoryavailable           string  `json:"memoryavailable"`
	Memorylimit               string  `json:"memorylimit"`
	Memorytotal               int64   `json:"memorytotal"`
	Name                      string  `json:"name"`
	Networkavailable          string  `json:"networkavailable"`
	Networklimit              string  `json:"networklimit"`
	Networktotal              int64   `json:"networktotal"`
	Primarystorageavailable   string  `json:"primarystorageavailable"`
	Primarystoragelimit       string  `json:"primarystoragelimit"`
	Primarystoragetotal       int64   `json:"primarystoragetotal"`
	Projectaccountname        string  `json:"projectaccountname"`
	Secondarystorageavailable string  `json:"secondarystorageavailable"`
	Secondarystoragelimit     string  `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64 `json:"secondarystoragetotal"`
	Snapshotavailable         string  `json:"snapshotavailable"`
	Snapshotlimit             string  `json:"snapshotlimit"`
	Snapshottotal             int64   `json:"snapshottotal"`
	State                     string  `json:"state"`
	Tags                      []Tags  `json:"tags"`
	Templateavailable         string  `json:"templateavailable"`
	Templatelimit             string  `json:"templatelimit"`
	Templatetotal             int64   `json:"templatetotal"`
	Vmavailable               string  `json:"vmavailable"`
	Vmlimit                   string  `json:"vmlimit"`
	Vmrunning                 int     `json:"vmrunning"`
	Vmstopped                 int     `json:"vmstopped"`
	Vmtotal                   int64   `json:"vmtotal"`
	Volumeavailable           string  `json:"volumeavailable"`
	Volumelimit               string  `json:"volumelimit"`
	Volumetotal               int64   `json:"volumetotal"`
	Vpcavailable              string  `json:"vpcavailable"`
	Vpclimit                  string  `json:"vpclimit"`
	Vpctotal                  int64   `json:"vpctotal"`
}

type UpdateProjectInvitationParams struct {
	p map[string]interface{}
}

func (p *UpdateProjectInvitationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accept"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("accept", vv)
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["token"]; found {
		u.Set("token", v.(string))
	}
	return u
}

func (p *UpdateProjectInvitationParams) SetAccept(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accept"] = v
}

func (p *UpdateProjectInvitationParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *UpdateProjectInvitationParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *UpdateProjectInvitationParams) SetToken(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["token"] = v
}

// You should always use this function to get a new UpdateProjectInvitationParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewUpdateProjectInvitationParams(projectid string) *UpdateProjectInvitationParams {
	p := &UpdateProjectInvitationParams{}
	p.p = make(map[string]interface{})
	p.p["projectid"] = projectid
	return p
}

// Accepts or declines project invitation
func (s *ProjectService) UpdateProjectInvitation(p *UpdateProjectInvitationParams) (*UpdateProjectInvitationResponse, error) {
	resp, err := s.cs.newRequest("updateProjectInvitation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateProjectInvitationResponse
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

type UpdateProjectInvitationResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
