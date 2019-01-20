package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type client struct {
}

// Client インターフェイスの定義
type Client interface {
	GetStatusCode(request *Request) (int, error)
	GetBody(request *Request) (string, error)
}

// Create はインスタンスを生成します
func Create() Client {
	return &client{}
}

// GetStatusCode はStatusCodeを取得します
func (client *client) GetStatusCode(request *Request) (int, error) {
	var req, err = http.NewRequest(request.Method, request.URL, nil)
	if err != nil {
		return -1, fmt.Errorf("Failed: create request : %s", err)
	}
	fmt.Printf("Requesting [%s] %s\n", req.Method, req.URL)

	c := new(http.Client)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Failed: request: %s\n", err)
		return -1, nil
	}
	defer resp.Body.Close()

	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
	return resp.StatusCode, nil
}

// GetBody はBodyを取得します
func (client *client) GetBody(request *Request) (string, error) {
	var req, err = http.NewRequest(request.Method, request.URL, nil)
	if err != nil {
		return "", fmt.Errorf("Failed: create request : %s", err)
	}
	fmt.Printf("Requesting [%s] %s\n", req.Method, req.URL)

	c := new(http.Client)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Failed: request: %s\n", err)
		return "", nil
	}
	defer resp.Body.Close()

	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed: read body: %s\n", err)
		return "", nil
	}

	strBody := string(body)
	fmt.Printf("Result: %v\n", strBody)
	return strBody, nil
}
