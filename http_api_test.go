package goimouse

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewApi(t *testing.T) {
	api := NewHttpApi("192.168.23.19")
	deviceId := "D8:DC:40:53:20:5D"

	// 使用 HTTP API 示例
	// Test GetDeviceList method
	deviceList, err := api.GetDeviceList(true)
	if err != nil {
		t.Errorf("GetDeviceList failed: %v", err)
	} else if deviceList == nil {
		t.Error("GetDeviceList returned nil")
	}
	fmtMapToJson(deviceList)

	// Test GetGroupList method
	groupList, err := api.GetGroupList(true)
	if err != nil {
		t.Errorf("GetGroupList failed: %v", err)
	} else if groupList == nil {
		t.Error("GetGroupList returned nil")
	}
	fmtMapToJson(groupList)

	// Test GetUsbList method
	usbList, err := api.GetUsbList(true)
	if err != nil {
		t.Errorf("GetUsbList failed: %v", err)
	} else if usbList == nil {
		t.Error("GetUsbList returned nil")
	}
	fmtMapToJson(usbList)

	// Test GetDeviceModelList method
	deviceModelList, err := api.GetDeviceModelList(true)
	if err != nil {
		t.Errorf("GetDeviceModelList failed: %v", err)
	} else if deviceModelList == nil {
		t.Error("GetDeviceModelList returned nil")
	}
	fmtMapToJson(deviceModelList)

	// Test GetDeviceScreenshot method
	screenshot, err := api.GetDeviceScreenshot(deviceId, false, false, false, false, true)
	if err != nil {
		t.Errorf("GetDeviceScreenshot failed: %v", err)
	} else if screenshot == nil {
		t.Error("GetDeviceScreenshot returned nil")
	}
	fmtMapToJson(screenshot)

	// Test LoopDeviceScreenshot method
	loopScreenshot, err := api.LoopDeviceScreenshot(deviceId, 10, false, false, true)
	if err != nil {
		t.Errorf("LoopDeviceScreenshot failed: %v", err)
	} else if loopScreenshot == nil {
		t.Error("LoopDeviceScreenshot returned nil")
	}
	fmtMapToJson(loopScreenshot)

	// Test Click method
	clickResult, err := api.Click(deviceId, 100, 200, "left", 1, true)
	if err != nil {
		t.Errorf("Click failed: %v", err)
	} else if clickResult == nil {
		t.Error("Click returned nil")
	}
	fmtMapToJson(clickResult)

	// Test Swipe method
	swipeResult, err := api.Swipe(deviceId, "left", "left", 100.0, 0, 0, 100, 100, 0, true)
	if err != nil {
		t.Errorf("Swipe failed: %v", err)
	} else if swipeResult == nil {
		t.Error("Swipe returned nil")
	}
	fmtMapToJson(swipeResult)

	// Test MouseUp method
	mouseUpResult, err := api.MouseUp(deviceId, "left", true)
	if err != nil {
		t.Errorf("MouseUp failed: %v", err)
	} else if mouseUpResult == nil {
		t.Error("MouseUp returned nil")
	}
	fmtMapToJson(mouseUpResult)

	// Test MouseDown method
	mouseDownResult, err := api.MouseDown(deviceId, "left", true)
	if err != nil {
		t.Errorf("MouseDown failed: %v", err)
	} else if mouseDownResult == nil {
		t.Error("MouseDown returned nil")
	}
	fmtMapToJson(mouseDownResult)

	// Test MouseMove method
	mouseMoveResult, err := api.MouseMove(deviceId, 100, 200, true)
	if err != nil {
		t.Errorf("MouseMove failed: %v", err)
	} else if mouseMoveResult == nil {
		t.Error("MouseMove returned nil")
	}
	fmtMapToJson(mouseMoveResult)

	// Test MouseWheel method
	mouseWheelResult, err := api.MouseWheel(deviceId, "up", 10, 1, true)
	if err != nil {
		t.Errorf("MouseWheel failed: %v", err)
	} else if mouseWheelResult == nil {
		t.Error("MouseWheel returned nil")
	}
	fmtMapToJson(mouseWheelResult)

	// Test MouseResetPos method
	mouseResetPosResult, err := api.MouseResetPos(deviceId, true)
	if err != nil {
		t.Errorf("MouseResetPos failed: %v", err)
	} else if mouseResetPosResult == nil {
		t.Error("MouseResetPos returned nil")
	}
	fmtMapToJson(mouseResetPosResult)

	// Test SendKey method
	sendKeyResult, err := api.SendKey(deviceId, "A", nil, true)
	if err != nil {
		t.Errorf("SendKey failed: %v", err)
	} else if sendKeyResult == nil {
		t.Error("SendKey returned nil")
	}
	fmtMapToJson(sendKeyResult)

	// Test Restart method
	//restartResult, err := api.Restart(true)
	//if err != nil {
	//	t.Errorf("Restart failed: %v", err)
	//} else if restartResult == nil {
	//	t.Error("Restart returned nil")
	//}
	//fmtMapToJson(restartResult)

	// Test SaveDevLocation method
	saveDevLocationResult, err := api.SaveDevLocation(deviceId, "location1", true)
	if err != nil {
		t.Errorf("SaveDevLocation failed: %v", err)
	} else if saveDevLocationResult == nil {
		t.Error("SaveDevLocation returned nil")
	}
	fmtMapToJson(saveDevLocationResult)

	// Test DelDevLocation method
	delDevLocationResult, err := api.DelDevLocation("model1", "version1", "crc1", true)
	if err != nil {
		t.Errorf("DelDevLocation failed: %v", err)
	} else if delDevLocationResult == nil {
		t.Error("DelDevLocation returned nil")
	}
	fmtMapToJson(delDevLocationResult)

	// Test GetAirplaySrvNum method
	airplaySrvNum, err := api.GetAirplaySrvNum(true)
	if err != nil {
		t.Errorf("GetAirplaySrvNum failed: %v", err)
	} else if airplaySrvNum == nil {
		t.Error("GetAirplaySrvNum returned nil")
	}
	fmtMapToJson(airplaySrvNum)

	// Test GetUsbAutoAirplay method
	usbAutoAirplay, err := api.GetUsbAutoAirplay(deviceId, true)
	if err != nil {
		t.Errorf("GetUsbAutoAirplay failed: %v", err)
	} else if usbAutoAirplay == nil {
		t.Error("GetUsbAutoAirplay returned nil")
	}
	fmtMapToJson(usbAutoAirplay)

	// Test GetAirplayMode method
	airplayMode, err := api.GetAirplayMode(deviceId, true)
	if err != nil {
		t.Errorf("GetAirplayMode failed: %v", err)
	} else if airplayMode == nil {
		t.Error("GetAirplayMode returned nil")
	}
	fmtMapToJson(airplayMode)

	// Test AutoConnectScreenAll method
	autoConnectScreenAll, err := api.AutoConnectScreenAll(true)
	if err != nil {
		t.Errorf("AutoConnectScreenAll failed: %v", err)
	} else if autoConnectScreenAll == nil {
		t.Error("AutoConnectScreenAll returned nil")
	}
	fmtMapToJson(autoConnectScreenAll)

	// Test SaveRestartPoint method
	saveRestartPointResult, err := api.SaveRestartPoint(deviceId, true)
	if err != nil {
		t.Errorf("SaveRestartPoint failed: %v", err)
	} else if saveRestartPointResult == nil {
		t.Error("SaveRestartPoint returned nil")
	}
	fmtMapToJson(saveRestartPointResult)

	// Test DisconAirplay method
	disconAirplayResult, err := api.DisconAirplay(deviceId, true)
	if err != nil {
		t.Errorf("DisconAirplay failed: %v", err)
	} else if disconAirplayResult == nil {
		t.Error("DisconAirplay returned nil")
	}
	fmtMapToJson(disconAirplayResult)

	// Test AutoConnectScreen method
	autoConnectScreenResult, err := api.AutoConnectScreen(deviceId, false, true)
	if err != nil {
		t.Errorf("AutoConnectScreen failed: %v", err)
	} else if autoConnectScreenResult == nil {
		t.Error("AutoConnectScreen returned nil")
	}
	fmtMapToJson(autoConnectScreenResult)

	// Test RestartUsb method
	restartUsbResult, err := api.RestartUsb(deviceId, true)
	if err != nil {
		t.Errorf("RestartUsb failed: %v", err)
	} else if restartUsbResult == nil {
		t.Error("RestartUsb returned nil")
	}
	fmtMapToJson(restartUsbResult)

	// Test Shortcut method
	shortcutResult, err := api.Shortcut(deviceId, 1, []string{"device2"}, map[string]interface{}{"param1": "value1"}, 30, true)
	if err != nil {
		t.Errorf("Shortcut failed: %v", err)
	} else if shortcutResult == nil {
		t.Error("Shortcut returned nil")
	}
	fmtMapToJson(shortcutResult)
}

func fmtMapToJson(data map[string]interface{}) {
	jsonScreenshot, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonScreenshot))

}
