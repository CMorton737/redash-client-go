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
	// "log"
	// "github.com/davecgh/go-spew/spew"
)

// Visualization representation
type Visualization struct {
	Description string `json:"description,omitempty"`
	ID          int    `json:"id,omitempty"`
	QueryID     int    `json:"query_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Options     map[string]interface{} `json:"options,omitempty"` 
	Type        string `json:"type,omitempty"`
}


//GetVisualization gets a specific Visualization
func (c *Client) GetVisualization(query_id int, id int) (*Visualization, error) {
	path := "/api/queries/" + strconv.Itoa(query_id)

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

	for _, x := range query.Visualizations {
		if x.ID == id {
			visualization := Visualization {
				Description: x.Description,
				ID:          x.ID,
				QueryID:     x.QueryID,
				Name:        x.Name,
				Options:     x.Options,
				Type:        x.Type,
			}
			return &visualization, nil
		}
	}

	return nil, nil	
}

// CreateVisualization creates a new Redash visualization
func (c *Client) CreateVisualization(visualizationCreatePayload *Visualization) (*Visualization, error) {
	path := "/api/visualizations"

	payload, err := json.Marshal(visualizationCreatePayload)
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

	visualization := Visualization{}

	err = json.Unmarshal(body, &visualization)
	if err != nil {
		return nil, err
	}

	return &visualization, nil
}

// UpdateVisualization 
// updates an existing Redash visualization
func (c *Client) UpdateVisualization(id int, visualizationUpdatePayload *Visualization) (*Visualization, error) {
	path := "/api/visualizations/" + strconv.Itoa(id)

	payload, err := json.Marshal(visualizationUpdatePayload)
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

	visualization := Visualization{}

	err = json.Unmarshal(body, &visualization)
	if err != nil {
		return nil, err
	}

	return &visualization, nil
}

//DeleteVisualization deletes an active visualization.
func (c *Client) DeleteVisualization(id int) error {
	path := "/api/visualizations/" + strconv.Itoa(id)

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