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
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

// QueryList struct
type QueryList struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Results  []struct {
		IsArchived        bool        `json:"is_archived,omitempty"`
		UpdatedAt         time.Time   `json:"updated_at,omitempty"`
		IsFavorite        bool        `json:"is_favorite,omitempty"`
		Query             string      `json:"query,omitempty"`
		ID                int         `json:"id,omitempty"`
		Description       interface{} `json:"description,omitempty"`
		Tags              []string    `json:"tags,omitempty"`
		Version           int         `json:"version,omitempty"`
		QueryHash         string      `json:"query_hash,omitempty"`
		APIKey            string      `json:"api_key,omitempty"`
		DataSourceID      int         `json:"data_source_id,omitempty"`
		IsSafe            bool        `json:"is_safe,omitempty"`
		LatestQueryDataID int         `json:"latest_query_data_id,omitempty"`
		Schedule          interface{} `json:"schedule,omitempty"`
		IsDraft           bool        `json:"is_draft,omitempty"`
		CanEdit           bool        `json:"can_edit,omitempty"`
		Name              string      `json:"name,omitempty"`
		CreatedAt         time.Time   `json:"created_at,omitempty"`
		Options map[string]interface{} `json:"options,omitempty"`
		Visualizations    []struct {
			Description string    `json:"description,omitempty"`
			CreatedAt   time.Time `json:"created_at,omitempty"`
			UpdatedAt   time.Time `json:"updated_at,omitempty"`
			ID          int       `json:"id,omitempty"`
			Type        string    `json:"type,omitempty"`
			QueryID     int       `json:"query_id,omitempty"`
			Options     map[string]interface{}    `json:"options,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"visualizations,omitempty"`
	} `json:"results,omitempty"`
}

// Query representation
type Query struct {
	IsArchived        bool        `json:"is_archived,omitempty"`
	UpdatedAt         time.Time   `json:"updated_at,omitempty"`
	IsFavorite        bool        `json:"is_favorite,omitempty"`
	Query             string      `json:"query,omitempty"`
	ID                int         `json:"id,omitempty"`
	Description       interface{} `json:"description,omitempty"`
	Tags              []string    `json:"tags,omitempty"`
	Version           int         `json:"version,omitempty"`
	QueryHash         string      `json:"query_hash,omitempty"`
	APIKey            string      `json:"api_key,omitempty"`
	DataSourceID      int         `json:"data_source_id,omitempty"`
	IsSafe            bool        `json:"is_safe,omitempty"`
	LatestQueryDataID int         `json:"latest_query_data_id,omitempty"`
	Schedule          interface{} `json:"schedule,omitempty"`
	IsDraft           bool        `json:"is_draft,omitempty"`
	CanEdit           bool        `json:"can_edit,omitempty"`
	Name              string      `json:"name,omitempty"`
	CreatedAt         time.Time   `json:"created_at,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
	Visualizations    []struct {
		Description string    `json:"description,omitempty"`
		CreatedAt   time.Time `json:"created_at,omitempty"`
		UpdatedAt   time.Time `json:"updated_at,omitempty"`
		ID          int       `json:"id,omitempty"`
		Type        string    `json:"type,omitempty"`
		QueryID     int       `json:"query_id,omitempty"`
		Options     map[string]interface{}    `json:"options,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"visualizations,omitempty"`
}

//GetQueries returns a paginated list of queries
func (c *Client) GetQueries() (*QueryList, error) {
	path := "/api/queries"

	response, err := c.get(path)

	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(response.Body)

	queries := QueryList{}
	err = json.Unmarshal(body, &queries)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return &queries, nil
}

//GetQuery gets a specific Query
func (c *Client) GetQuery(id int) (*Query, error) {
	path := "/api/queries/" + strconv.Itoa(id)

	response, err := c.get(path)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	query := Query{}

	err = json.Unmarshal(body, &query)
	if err != nil {
		return nil, err
	}

	return &query, nil
}

// CreateQuery creates a new Redash query
func (c *Client) CreateQuery(queryCreatePayload *Query) (*Query, error) {
	path := "/api/queries"

	payload, err := json.Marshal(queryCreatePayload)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(payload))
	response, err := c.post(path, string(payload))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	query := Query{}

	err = json.Unmarshal(body, &query)
	if err != nil {
		return nil, err
	}

	return &query, nil
}

// UpdateQuery updates an existing Redash query
func (c *Client) UpdateQuery(id int, queryUpdatePayload *Query) (*Query, error) {
	path := "/api/queries/" + strconv.Itoa(id)

	payload, err := json.Marshal(queryUpdatePayload)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(payload))
	response, err := c.post(path, string(payload))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	query := Query{}

	err = json.Unmarshal(body, &query)
	if err != nil {
		return nil, err
	}

	return &query, nil
}

//DeleteQuery deletes an active query.
func (c *Client) DeleteQuery(id int) error {
	path := "/api/queries/" + strconv.Itoa(id)

	response, err := c.delete(path)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return nil
}

// //SearchQueries finds a list of queries matching a string (searches `name` and `email` fields)
// func (c *Client) SearchQueries(term string) (*QueryList, error) {
// 	path := "/api/queries?q=" + term

// 	response, err := c.get(path)

// 	if err != nil {
// 		return nil, err
// 	}
// 	body, _ := ioutil.ReadAll(response.Body)

// 	queries := QueryList{}
// 	err = json.Unmarshal(body, &queries)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer response.Body.Close()

// 	return &queries, nil
// }

// // GetQueryByEmail returns a single  query from their email address
// func (c *Client) GetQueryByEmail(email string) (*Query, error) {

// 	results, err := c.SearchQueries(email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, result := range results.Results {
// 		if result.Email != "" && result.Email == email {
// 			return c.GetQuery(result.ID)
// 		}
// 	}

// 	return nil, fmt.Errorf("No query found with email address: %s", email)
// }

///https://dashboard.uat.qa.mhgi.io/api/queries/32/favorite
