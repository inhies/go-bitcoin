package bitcoin

import (
	"encoding/json"
	"fmt"
)

// TODO(inhies): Add more calls: https://en.bitcoin.it/wiki/Original_Bitcoin_client/API_calls_list

// TODO(inhies): Deal with rounding of btc values; Decide between float or int.

// Executes a `getinfo` API call
func (conn *Client) GetInfo() (result ApiInfo, err error) {
	response, err := conn.sendRequest("getinfo", nil)
	if err != nil {
		return
	}
	var data = &apiInfo{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return
	}

	// Check for an error response from the API
	if data.Error != nil {
		err = fmt.Errorf(*data.Error)
		return
	}
	result = data.Result
	return
}

// Executes a `getblockcount` API call
func (conn *Client) GetBlockCount() (count int64, err error) {
	response, err := conn.sendRequest("getblockcount", nil)
	var data = &apiBlockCount{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return
	}

	// Check for an error response from the API
	if data.Error != nil {
		err = fmt.Errorf(*data.Error)
		return
	}
	count = data.Result
	return
}
