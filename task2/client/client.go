package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type InputString struct {
	Input string `json:"inputString"`
}

type OutputString struct {
	Output string `json:"outputString"`
}

type Client struct {
	client    *http.Client
	serverUrl string
}

func NewClient(serverUrl string) *Client {
	return &Client{client: &http.Client{}, serverUrl: serverUrl}
}

func (c *Client) GetVersion() error {
	req, err := http.NewRequest(
		"GET", c.serverUrl+"/version", nil,
	)

	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}

func (c *Client) DecodeMessage(message string) error {

	jsonBody, err := json.Marshal(InputString{Input: message})

	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"Post", c.serverUrl+"/decode", bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var res OutputString
	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res.Output)

	return nil
}

func (c *Client) HardOp() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequest("GET", c.serverUrl+"/hard-op", nil)
	req = req.WithContext(ctx)

	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println(false, 504)
		}
		return err
	}
	defer resp.Body.Close()

	fmt.Println(true, resp.StatusCode)

	return nil
}
