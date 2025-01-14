package goimouse

import (
	"fmt"
	"testing"
	"time"
)

func TestWs(t *testing.T) {

	//使用 WebSocket API 示例
	wsApi := NewWsApi("192.168.23.19", func(data map[string]interface{}) {
		fmt.Println("Callback:", data)
	})

	if err := wsApi.Start(); err != nil {
		fmt.Println("WebSocket Error:", err)
		return
	}

	// 模拟操作
	time.Sleep(10 * time.Minute)
	wsApi.Stop()

}
