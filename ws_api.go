package goimouse

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type WsApi struct {
	HttpApi
	host      string
	isWorking bool
	ws        *websocket.Conn
	callback  func(map[string]interface{})
	msgLock   sync.Mutex
	msgMap    map[int]map[string]interface{}
}

func NewWsApi(host string, callback func(map[string]interface{})) *WsApi {
	return &WsApi{
		HttpApi:  *NewHttpApi(host),
		host:     host,
		callback: callback,
		msgMap:   make(map[int]map[string]interface{}),
	}
}

func (api *WsApi) Start() error {
	api.isWorking = true
	var err error
	api.ws, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:9911/clinet", api.host), nil)
	if err != nil {
		return err
	}

	go api.listen()
	return nil
}

func (api *WsApi) listen() {
	for api.isWorking {
		_, message, err := api.ws.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			continue
		}

		var data map[string]interface{}
		if err := json.Unmarshal(message, &data); err != nil {
			fmt.Println("Error unmarshalling message:", err)
			continue
		}

		if msgID, ok := data["msgid"].(int); ok {
			api.msgLock.Lock()
			api.msgMap[msgID] = data
			api.msgLock.Unlock()
		} else if api.callback != nil {
			api.callback(data)
		}
	}
}

func (api *WsApi) Stop() {
	api.isWorking = false
	api.ws.Close()
}

// 实现 WsApi 的方法
