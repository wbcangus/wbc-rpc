package consumer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"wbc-rpc/model"
)

type Client struct {
	address string
}

func NewClient() *Client {
	return &Client{
		address: "http://127.0.0.1:8088",
	}
}

func (c *Client) CallRpc(request model.RpcRequest) error {
	url := c.address + "/rpc"
	jsonData, err := json.Marshal(request)
	if err != nil {
		return err
	}
	fmt.Println("RPC call to", url, "with request:", string(jsonData))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("RPC call failed:", err)
		return err
	}
	fmt.Println("RPC call successful, response status code:", resp.StatusCode)
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read RPC response:", err)
		return err
	}
	defer resp.Body.Close()
	var rpcResponse model.RpcResponse
	err = json.Unmarshal(res, &rpcResponse)
	if err != nil {
		log.Println(string(res))
		log.Println("Failed to unmarshal RPC response:", err)
		return err
	}
	// Handle response as needed.
	// Here we just return nil to indicate success.
	fmt.Printf("RPC call successful, response: %+v\n", rpcResponse)
	return nil
}
