//
// Copyright 2018, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cloudstack

import (
	"encoding/json"
	"net/url"
)

type GetApiLimitParams struct {
	p map[string]interface{}
}

func (p *GetApiLimitParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new GetApiLimitParams instance,
// as then you are sure you have configured all required params
func (s *APIDiscoveryService) NewGetApiLimitParams() *GetApiLimitParams {
	p := &GetApiLimitParams{}
	p.p = make(map[string]interface{})
	return p
}

// Get API limit count for the caller
func (s *APIDiscoveryService) GetApiLimit(p *GetApiLimitParams) (*GetApiLimitResponse, error) {
	resp, err := s.cs.newRequest("getApiLimit", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetApiLimitResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetApiLimitResponse struct {
	Account     string `json:"account"`
	Accountid   string `json:"accountid"`
	ApiAllowed  int    `json:"apiAllowed"`
	ApiIssued   int    `json:"apiIssued"`
	ExpireAfter int64  `json:"expireAfter"`
}

type ListApisParams struct {
	p map[string]interface{}
}

func (p *ListApisParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *ListApisParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

// You should always use this function to get a new ListApisParams instance,
// as then you are sure you have configured all required params
func (s *APIDiscoveryService) NewListApisParams() *ListApisParams {
	p := &ListApisParams{}
	p.p = make(map[string]interface{})
	return p
}

// lists all available apis on the server, provided by the Api Discovery plugin
func (s *APIDiscoveryService) ListApis(p *ListApisParams) (*ListApisResponse, error) {
	resp, err := s.cs.newRequest("listApis", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListApisResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListApisResponse struct {
	Count int    `json:"count"`
	Apis  []*Api `json:"api"`
}

type Api struct {
	Description string `json:"description"`
	Isasync     bool   `json:"isasync"`
	Name        string `json:"name"`
	Params      []struct {
		Description string `json:"description"`
		Length      int    `json:"length"`
		Name        string `json:"name"`
		Related     string `json:"related"`
		Required    bool   `json:"required"`
		Since       string `json:"since"`
		Type        string `json:"type"`
	} `json:"params"`
	Related  string `json:"related"`
	Response []struct {
		Description string        `json:"description"`
		Name        string        `json:"name"`
		Response    []interface{} `json:"response"`
		Type        string        `json:"type"`
	} `json:"response"`
	Since string `json:"since"`
	Type  string `json:"type"`
}

type ResetApiLimitParams struct {
	p map[string]interface{}
}

func (p *ResetApiLimitParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	return u
}

func (p *ResetApiLimitParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

// You should always use this function to get a new ResetApiLimitParams instance,
// as then you are sure you have configured all required params
func (s *APIDiscoveryService) NewResetApiLimitParams() *ResetApiLimitParams {
	p := &ResetApiLimitParams{}
	p.p = make(map[string]interface{})
	return p
}

// Reset api count
func (s *APIDiscoveryService) ResetApiLimit(p *ResetApiLimitParams) (*ResetApiLimitResponse, error) {
	resp, err := s.cs.newRequest("resetApiLimit", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResetApiLimitResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ResetApiLimitResponse struct {
	Account     string `json:"account"`
	Accountid   string `json:"accountid"`
	ApiAllowed  int    `json:"apiAllowed"`
	ApiIssued   int    `json:"apiIssued"`
	ExpireAfter int64  `json:"expireAfter"`
}
