package model

type RpcRequest struct {
	ServiceName string        `json:"serviceName"`
	MethodName  string        `json:"methodName"`
	Args        []interface{} `json:"args"`
}

type RpcResponse struct {
	Result interface{} `json:"result"`
}
