package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer)Connect() chan string {
	session := make(chan string, 0)

	//annoymous function
	go func(c chan string) {
		for {
			request := <-c
			if request == "CLOSE" {  //close the connect
				break
			}
			var req Request
			//Unmarshal函数解析json编码的数据并将结果存入req指向的值。
			err := json.Unmarshal([]byte(request), &req)   //***?
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}

			resp := server.Handle(req.Method, req.Params)
			//Unmarshal和Marshal做相反的操作，必要时申请映射、切片或指针
			//Marshal编码json数据并将结果返回
			b, err := json.Marshal(resp)   //****?
			c <- string(b)   //return result
		}
		fmt.Println("Session closed.")

	}(session)

	fmt.Println("A new session has been created successfully.")

	return session
}
