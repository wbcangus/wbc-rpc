package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"wbc-rpc/config"
	"wbc-rpc/provider"
)

type Service interface {
	DoService()
}

type UserService struct {
}

func (u *UserService) DoService() {
	fmt.Println("do user service")
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	config.GetConfig()

	provider.StartProvider()
}
