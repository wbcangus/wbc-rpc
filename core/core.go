package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"wbc-rpc/constant"
	"wbc-rpc/model"
)

// RpcServer 定义一个RpcServer结构体
type RpcServer struct {
	// 保存服务提供方注册的服务对象
	Handler map[string]reflect.Value
}

// NewRpcServer 创建一个RpcServer对象
func NewRpcServer() *RpcServer {
	return &RpcServer{
		Handler: make(map[string]reflect.Value),
	}
}

// RegisterService 注册一个服务对象
func (s *RpcServer) RegisterService(serviceName string, object reflect.Value) {
	if _, ok := s.Handler[serviceName]; !ok {
		log.Println("注册服务对象", serviceName)
		s.Handler[serviceName] = object
	}
	return
}

// RegisterService 注册一个服务对象
func (s *RpcServer) Start() {
	// TODO: 启动RpcServer
	mux := http.NewServeMux()
	mux.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("收到请求")
		// TODO: 处理Rpc请求
		// 获取请求的数据
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// 解析请求中的json数据
		var requestData model.RpcRequest
		if err := json.Unmarshal(body, &requestData); err != nil {
			http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
			return
		}
		res, err := s.Call(requestData.ServiceName, requestData.MethodName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("服务调用结果==============> ", reflect.TypeOf(res), "  ===  ", reflect.ValueOf(res))
		// 构造响应数据
		responseData := model.RpcResponse{
			Result: res,
		}
		// 序列化响应数据
		response, err := json.Marshal(responseData)
		if err != nil {
			http.Error(w, "Failed to marshal JSON data", http.StatusInternalServerError)
			return
		}
		log.Println("响应数据", string(response))
		// 发送响应数据
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	})
	err := http.ListenAndServe("127.0.0.1:8088", mux)
	if err != nil {
		return
	}
	return
}

// Call 调用服务对象的方法
func (s *RpcServer) Call(serviceName string, methodName string, in ...[]reflect.Value) (result interface{}, err error) {
	if object, ok := s.Handler[serviceName]; ok {
		log.Println("获取服务对象", serviceName)
		method := object.MethodByName(methodName)

		log.Println("调用服务对象的方法", serviceName, methodName)
		callRes := method.Call(nil)
		if len(callRes) > 0 {
			result = callRes[0].Interface()
		}
		return
	}
	err = errors.New(constant.ErrMethodNotFound)
	return
}
