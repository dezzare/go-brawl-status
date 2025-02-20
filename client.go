package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

type Client struct {
	HTTP    *http.Client
	BaseURL string
	APIKey  string
}

func newClient() *Client {
	return &Client{
		HTTP:    &http.Client{},
		BaseURL: viper.GetString("BaseURL"),
		APIKey:  viper.GetString("APIKey"),
	}
}

func (c *Client) getPlayer(playerTag string) (data []byte, err error) {
	data, err = c.doRequest("GET", c.BaseURL+"/players/"+playerTag)
	if err != nil {
		return data, fmt.Errorf("Error getting player: %v", err)
	}
	return
}

func (c *Client) getBrawlers() (data []byte, err error) {
	data, err = c.doRequest("GET", c.BaseURL+"/brawlers")
	if err != nil {
		return data, fmt.Errorf("Error getting brawlers: %v", err)
	}
	return
}

func (c *Client) doRequest(method string, path string) (data []byte, err error) {
	req, _ := http.NewRequest(method, path, nil)
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %v", err)
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Read IO error: %v", err)

	}
	if http.StatusText(res.StatusCode) != "OK" {
		return nil, fmt.Errorf("\nERROR: HTTP Status code %v\n", res.StatusCode)
	}

	fmt.Printf("\nStatus Code: %v \nHTTP Status: %v\n", res.StatusCode, http.StatusText(res.StatusCode))
	return
}
