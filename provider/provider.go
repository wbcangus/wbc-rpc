package provider

import (
	"reflect"
	"wbc-rpc/core"
)

var rpcServer = core.NewRpcServer()

func StartProvider() {
	var userService UserService
	userService = &UserServiceImpl{}
	value := reflect.ValueOf(userService)
	rpcServer.RegisterService("userService", value)
	rpcServer.Start()
}

type UserService interface {
	GetUser() string
}

type UserServiceImpl struct {
}

func (u *UserServiceImpl) GetUser() string {
	return "jack your father"
}
