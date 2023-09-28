package gecko

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/verzth/go-coingecko/ping"
)

type Client struct {
	httpClient *http.Client
	baseUrl    string
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		httpClient: httpClient,
		baseUrl:    baseURL,
	}
}

func (c *Client) Close() {
	c.httpClient.CloseIdleConnections()
}

func doReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// MakeReq HTTP request helper
func (c *Client) MakeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	resp, err := doReq(req, c.httpClient)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Ping /ping endpoint
func (c *Client) Ping() (*ping.Ping, error) {
	resp, err := c.MakeReq(pingURL)
	if err != nil {
		return nil, err
	}
	var data *ping.Ping
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
