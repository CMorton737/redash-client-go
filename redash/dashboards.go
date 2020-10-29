//
// Copyright (c) 2020 Snowplow Analytics Ltd. All rights reserved.
//
// This program is licensed to you under the Apache License Version 2.0,
// and you may not use this file except in compliance with the Apache License Version 2.0.
// You may obtain a copy of the Apache License Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the Apache License Version 2.0 is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the Apache License Version 2.0 for the specific language governing permissions and limitations there under.
//

package redash

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
)

// DashboardList struct
type DashboardList struct {
	Count    int `json:"count,omitempty"`
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
	Results  []struct {
		Tags       []interface{} `json:"tags,omitempty"`
		IsArchived bool          `json:"is_archived,omitempty"`
		UpdatedAt  time.Time     `json:"updated_at,omitempty"`
		IsFavorite bool          `json:"is_favorite,omitempty"`
		Layout                  []interface{} `json:"layout,omitempty"`
		IsDraft                 bool          `json:"is_draft,omitempty"`
		ID                      int           `json:"id,omitempty"`
		UserID                  int           `json:"user_id,omitempty"`
		Name                    string        `json:"name,omitempty"`
		CreatedAt               time.Time     `json:"created_at,omitempty"`
		Slug                    string        `json:"slug,omitempty"`
		Version                 int           `json:"version,omitempty"`
		Widgets                 interface{}   `json:"widgets,omitempty"`
		DashboardFiltersEnabled bool          `json:"dashboard_filters_enabled,omitempty"`
	} `json:"results,omitempty"`
}

// Dashboard representation
type Dashboard struct {
	Tags       []interface{} `json:"tags,omitempty"`
	IsArchived bool          `json:"is_archived,omitempty"`
	UpdatedAt  time.Time     `json:"updated_at,omitempty"`
	IsFavorite bool          `json:"is_favorite,omitempty"`
	Layout                  []interface{} `json:"layout,omitempty"`
	IsDraft                 bool          `json:"is_draft,omitempty"`
	ID                      int           `json:"id,omitempty"`
	UserID                  int           `json:"user_id,omitempty"`
	Name                    string        `json:"name,omitempty"`
	CreatedAt               time.Time     `json:"created_at,omitempty"`
	Slug                    string        `json:"slug,omitempty"`
	Version                 int           `json:"version,omitempty"`
	Widgets                 interface{}   `json:"widgets,omitempty"`
	DashboardFiltersEnabled bool          `json:"dashboard_filters_enabled,omitempty"`
}

//GetDashboards returns a paginated list of dashboards
func (c *Client) GetDashboards() (*DashboardList, error) {
	path := "/api/dashboards"

	response, err := c.get(path)

	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(response.Body)

	dashboards := DashboardList{}
	err = json.Unmarshal(body, &dashboards)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return &dashboards, nil
}

//GetDashboard gets a specific Dashboard
func (c *Client) GetDashboard(id int) (*Dashboard, error) {
	path := "/api/dashboards/" + strconv.Itoa(id)

	response, err := c.get(path)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	Dashboard := Dashboard{}

	err = json.Unmarshal(body, &Dashboard)
	if err != nil {
		return nil, err
	}

	return &Dashboard, nil
}

// // CreateDashboard creates a new Redash dashboard
// func (c *Client) CreateDashboard(dashboardCreatePayload *DashboardCreatePayload) (*Dashboard, error) {
// 	path := "/api/dashboards"

// 	payload, err := json.Marshal(dashboardCreatePayload)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response, err := c.post(path, string(payload))
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer response.Body.Close()
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	dashboard := Dashboard{}

// 	err = json.Unmarshal(body, &dashboard)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &dashboard, nil
// }

// // UpdateDashboard updates an existing Redash dashboard
// func (c *Client) UpdateDashboard(id int, dashboardUpdatePayload *DashboardUpdatePayload) (*Dashboard, error) {
// 	path := "/api/dashboards/" + strconv.Itoa(id)

// 	payload, err := json.Marshal(dashboardUpdatePayload)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response, err := c.post(path, string(payload))
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer response.Body.Close()
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	dashboard := Dashboard{}

// 	err = json.Unmarshal(body, &dashboard)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &dashboard, nil
// }

// //DisableDashboard disables an active dashboard.
// func (c *Client) DisableDashboard(id int) error {
// 	path := "/api/dashboards/" + strconv.Itoa(id) + "/disable"

// 	response, err := c.post(path, "")
// 	if err != nil {
// 		return err
// 	}

// 	defer response.Body.Close()
// 	_, err = ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// //SearchDashboards finds a list of dashboards matching a string (searches `name` and `email` fields)
// func (c *Client) SearchDashboards(term string) (*DashboardList, error) {
// 	path := "/api/dashboards?q=" + term

// 	response, err := c.get(path)

// 	if err != nil {
// 		return nil, err
// 	}
// 	body, _ := ioutil.ReadAll(response.Body)

// 	dashboards := DashboardList{}
// 	err = json.Unmarshal(body, &dashboards)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer response.Body.Close()

// 	return &dashboards, nil
// }

// // GetDashboardByEmail returns a single  dashboard from their email address
// func (c *Client) GetDashboardByEmail(email string) (*Dashboard, error) {

// 	results, err := c.SearchDashboards(email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, result := range results.Results {
// 		if result.Email != "" && result.Email == email {
// 			return c.GetDashboard(result.ID)
// 		}
// 	}

// 	return nil, fmt.Errorf("No dashboard found with email address: %s", email)
// }
