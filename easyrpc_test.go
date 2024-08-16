package main

import (
	"github.com/wbcangus/wbc-rpc/consumer"
	"github.com/wbcangus/wbc-rpc/model"
	"testing"
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
