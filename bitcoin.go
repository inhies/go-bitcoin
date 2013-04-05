// Package bitcoin implements a set of functions for interfacing with bitcoind
package bitcoin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Bitcoind connection details
type Client struct {
	Username string
	Password string
	Addr     string
	UseSSL   bool
}

// Sends a HTTP JSON RPC request to bitcoin instance and returns the result
func (conn *Client) sendRequest(method string,
	params []interface{}) (response []byte, err error) {

	// Prepare the request
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     1, // This can be anything, but we don't need to use it.
		"params": params,
	})
	if err != nil {
		return nil, fmt.Errorf("Marshal: %v", err)
	}

	//fmt.Println(string(data))
	// If we are going to use SSL, append an 's' to 'http'
	var s string
	if conn.UseSSL {
		s = "s"
	}

	// Send the request
	resp, err := http.Post("http"+s+"://"+conn.Username+":"+conn.Password+"@"+conn.Addr,
		"application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, fmt.Errorf("Post: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll: %v", err)
	}
	//fmt.Println(string(response))
	return
}
