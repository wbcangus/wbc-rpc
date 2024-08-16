package main

import (
	"testing"
	"wbc-rpc/consumer"
	"wbc-rpc/model"
)

func TestRpcServer(t *testing.T) {
	err := consumer.NewClient().CallRpc(model.RpcRequest{
		ServiceName: "userService",
		MethodName:  "GetUser",
	})
	if err != nil {
		return
	}
}
