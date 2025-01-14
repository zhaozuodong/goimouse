package goimouse

import "encoding/json"

type DataCreation struct{}

func (dc *DataCreation) SetDevice(deviceID string, msgID int, data map[string]interface{}) map[string]interface{} {
	sendData := map[string]interface{}{
		"fun":   "set_dev",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
		},
	}
	for k, v := range data {
		sendData["data"].(map[string]interface{})[k] = v
	}
	return sendData
}

func (dc *DataCreation) DelDevice(deviceID string, msgID int) map[string]interface{} {
	sendData := map[string]interface{}{
		"fun":   "del_dev",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
		},
	}
	return sendData
}

func (dc *DataCreation) SetGroup(groupID, groupName string, msgID int) map[string]interface{} {
	sendData := map[string]interface{}{
		"fun":   "set_group",
		"msgid": msgID,
		"data": map[string]interface{}{
			"gid":  groupID,
			"name": groupName,
		},
	}
	return sendData
}

func (dc *DataCreation) DelGroup(groupID, msgID int) map[string]interface{} {
	sendData := map[string]interface{}{
		"fun":   "del_group",
		"msgid": msgID,
		"data": map[string]interface{}{
			"gid": groupID,
		},
	}
	return sendData
}

func (dc *DataCreation) GetDeviceList(msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "get_device_list",
		"msgid": msgID,
		"data":  map[string]interface{}{},
	}
}

// GetGroupList generates the JSON for "get_group_list" request
func (dc *DataCreation) GetGroupList(msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "get_group_list",
		"msgid": msgID,
		"data":  map[string]interface{}{},
	}
}

// GetUsbList generates the JSON for "get_usb_list" request
func (dc *DataCreation) GetUsbList(msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "get_usb_list",
		"msgid": msgID,
		"data":  map[string]interface{}{},
	}
}

// GetDeviceModelList generates the JSON for "get_devicemodel_list" request
func (dc *DataCreation) GetDeviceModelList(msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "get_devicemodel_list",
		"msgid": msgID,
		"data":  map[string]interface{}{},
	}
}

// GetDeviceScreenshot generates the JSON for "get_device_screenshot" request
func (dc *DataCreation) GetDeviceScreenshot(deviceID string, gzip, binary, isJpg, original bool, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "get_device_screenshot",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
			"gzip":     gzip,
			"binary":   binary,
			"isjpg":    isJpg,
			"original": original,
		},
	}
}

// LoopDeviceScreenshot generates the JSON for "loop_device_screenshot" request
func (dc *DataCreation) LoopDeviceScreenshot(deviceID string, time int, stop, isJpg bool, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "loop_device_screenshot",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
			"time":     time,
			"stop":     stop,
			"isjpg":    isJpg,
		},
	}
}

// Click generates the JSON for "click" request
func (dc *DataCreation) Click(deviceID string, x, y int, button string, time, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "click",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
			"button":   button,
			"x":        x,
			"y":        y,
			"time":     time,
		},
	}
}

// Swipe generates the JSON for "swipe" request
func (dc *DataCreation) Swipe(deviceID, direction, button string, length float64, sx, sy, ex, ey, afor, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "swipe",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid":  deviceID,
			"direction": direction,
			"button":    button,
			"length":    length,
			"sx":        sx,
			"sy":        sy,
			"ex":        ex,
			"ey":        ey,
			"for":       afor,
		},
	}
}

// MouseUp generates the JSON for "mouse_up" request
func (dc *DataCreation) MouseUp(deviceID, button string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "mouse_up",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
			"button":   button,
		},
	}
}

// MouseDown generates the JSON for "mouse_down" request
func (dc *DataCreation) MouseDown(deviceID, button string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "mouse_down",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
			"button":   button,
		},
	}
}

// MouseMove generates the JSON for "mouse_move" request
func (dc *DataCreation) MouseMove(deviceID string, x, y, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "mouse_move",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
			"x":        x,
			"y":        y,
		},
	}
}

// MouseWheel generates the JSON for "mouse_wheel" request
func (dc *DataCreation) MouseWheel(deviceID, direction string, length, number, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "mouse_wheel",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid":  deviceID,
			"direction": direction,
			"length":    length,
			"number":    number,
		},
	}
}

// MouseResetPos generates the JSON for "mouse_reset_pos" request
func (dc *DataCreation) MouseResetPos(deviceID string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "mouse_reset_pos",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
		},
	}
}

// SendKey generates the JSON for "send_key" request
func (dc *DataCreation) SendKey(deviceID, key string, fnKey *string, msgID int) map[string]interface{} {
	data := map[string]interface{}{
		"deviceid": deviceID,
		"key":      key,
	}
	if fnKey != nil {
		data["fn_key"] = *fnKey
	}
	return map[string]interface{}{
		"fun":   "send_key",
		"msgid": msgID,
		"data":  data,
	}
}

// Restart generates the JSON for "restart" request
func (dc *DataCreation) Restart(msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "restart",
		"msgid": msgID,
		"data":  map[string]interface{}{},
	}
}

// GetAirplaySrvNum generates the JSON for "get_airplaysrvnum" request
func (dc *DataCreation) GetAirplaySrvNum(msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "get_airplaysrvnum",
		"msgid": msgID,
		"data":  map[string]interface{}{},
	}
}

// SetAirplaySrvNum generates the JSON for "set_airplaysrvnum" request
func (dc *DataCreation) SetAirplaySrvNum(airplaySrvNum, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "set_airplaysrvnum",
		"msgid": msgID,
		"data": map[string]interface{}{
			"airplaysrvnum": airplaySrvNum,
		},
	}
}

// MouseCollectionOpen generates the JSON for "mouse_collection_open" request
func (dc *DataCreation) MouseCollectionOpen(deviceID string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "mouse_collection_open",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
		},
	}
}

// MouseCollectionClose generates the JSON for "mouse_collection_close" request
func (dc *DataCreation) MouseCollectionClose(deviceID string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "mouse_collection_close",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
		},
	}
}

// MouseCollectionCfg generates the JSON for "mouse_collection_cfg" request
func (dc *DataCreation) MouseCollectionCfg(deviceID string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "mouse_collection_cfg",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
		},
	}
}

// SaveDevLocation generates the JSON for "save_dev_location" request
func (dc *DataCreation) SaveDevLocation(deviceID, describe string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "save_dev_location",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
			"describe": describe,
		},
	}
}

// DelDevLocation generates the JSON for "del_dev_location" request
func (dc *DataCreation) DelDevLocation(model, version, crc string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "del_dev_location",
		"msgid": msgID,
		"data": map[string]interface{}{
			"model":   model,
			"version": version,
			"crc":     crc,
		},
	}
}

// SendText generates the JSON for "send_text" request
func (dc *DataCreation) SendText(deviceID, key, fnKey string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "send_text",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
			"key":      key,
			"fn_key":   fnKey,
		},
	}
}

// FindImage generates the JSON for "find_image" request
func (dc *DataCreation) FindImage(deviceID, img string, rect []interface{}, original bool, similarity float64, msgID int) map[string]interface{} {
	data := map[string]interface{}{
		"deviceid":   deviceID,
		"img":        img,
		"similarity": similarity,
		"rect":       rect,
		"original":   original,
	}
	if rect == nil {
		delete(data, "rect")
	}
	return map[string]interface{}{
		"fun":   "find_image",
		"msgid": msgID,
		"data":  data,
	}
}

// FindImageEx generates the JSON for "find_image_ex" request
func (dc *DataCreation) FindImageEx(deviceID string, imgList []interface{}, rect []interface{}, original, all, repeat bool, similarity float64, msgID int) map[string]interface{} {
	data := map[string]interface{}{
		"deviceid":   deviceID,
		"img_list":   imgList,
		"similarity": similarity,
		"all":        all,
		"repeat":     repeat,
		"rect":       rect,
		"original":   original,
	}
	if rect == nil {
		delete(data, "rect")
	}
	return map[string]interface{}{
		"fun":   "find_image_ex",
		"msgid": msgID,
		"data":  data,
	}
}

// Ocr generates the JSON for "ocr" request
func (dc *DataCreation) Ocr(fun, deviceID string, rect []interface{}, original bool, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   fun,
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
			"rect":     rect,
			"original": original,
		},
	}
}

// FindMultiColor generates the JSON for "find_multi_color" request
func (dc *DataCreation) FindMultiColor(deviceID string, rect []int, firstColor, offsetColor string, similarity float64, dir, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "find_multi_color",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid":     deviceID,
			"x1":           rect[0],
			"y1":           rect[1],
			"x2":           rect[2],
			"y2":           rect[3],
			"first_color":  firstColor,
			"offset_color": offsetColor,
			"similarity":   similarity,
			"dir":          dir,
		},
	}
}

// AutoConnectScreen generates the JSON for "auto_connect_screen" request
func (dc *DataCreation) AutoConnectScreen(deviceID string, force bool, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "auto_connect_screen",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
			"force":    force,
		},
	}
}

// SaveAutoScreenPoint generates the JSON for "save_autoscreen_point" request
func (dc *DataCreation) SaveAutoScreenPoint(deviceID string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "save_autoscreen_point",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
		},
	}
}

// RestartUsb generates the JSON for "restart_usb" request
func (dc *DataCreation) RestartUsb(deviceID string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "restart_usb",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
		},
	}
}

// SetUsbAutoAirplay generates the JSON for "set_usb_autoairplay" request
func (dc *DataCreation) SetUsbAutoAirplay(deviceID string, autoAirplay bool, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "set_usb_autoairplay",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid":    deviceID,
			"autoairplay": autoAirplay,
		},
	}
}

// GetUsbAutoAirplay generates the JSON for "get_usb_autoairplay" request
func (dc *DataCreation) GetUsbAutoAirplay(deviceID string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "get_usb_autoairplay",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
		},
	}
}

// SetAirplayMode generates the JSON for "set_airplay_mode" request
func (dc *DataCreation) SetAirplayMode(deviceID string, airplayMode, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "set_airplay_mode",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid":     deviceID,
			"airplay_mode": airplayMode,
		},
	}
}

// GetAirplayMode generates the JSON for "get_airplay_mode" request
func (dc *DataCreation) GetAirplayMode(deviceID string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "get_airplay_mode",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
		},
	}
}

// SaveRestartPoint generates the JSON for "save_restart_point" request
func (dc *DataCreation) SaveRestartPoint(deviceID string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "save_restart_point",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
		},
	}
}

// RestartDevice generates the JSON for "restart_device" request
func (dc *DataCreation) RestartDevice(deviceID string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "restart_device",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
		},
	}
}

// SetMdns generates the JSON for "set_mdns" request
func (dc *DataCreation) SetMdns(srvName string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "set_mdns",
		"msgid": msgID,
		"data": map[string]interface{}{
			"srvname": srvName,
		},
	}
}

// AutoConnectScreenAll generates the JSON for "auto_connect_screen_all" request
func (dc *DataCreation) AutoConnectScreenAll(msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "auto_connect_screen_all",
		"msgid": msgID,
		"data":  map[string]interface{}{},
	}
}

// DisconAirplay generates the JSON for "discon_airplay" request
func (dc *DataCreation) DisconAirplay(deviceID string, msgID int) map[string]interface{} {
	return map[string]interface{}{
		"fun":   "discon_airplay",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid": deviceID,
		},
	}
}

// Shortcut generates the JSON for "shortcut" request
func (dc *DataCreation) Shortcut(deviceID string, id int, devList []string, parameter map[string]interface{}, outTime, msgID int) map[string]interface{} {
	paramJSON, _ := json.Marshal(parameter)
	return map[string]interface{}{
		"fun":   "shortcut",
		"msgid": msgID,
		"data": map[string]interface{}{
			"deviceid":  deviceID,
			"id":        id,
			"devlist":   devList,
			"outtime":   outTime,
			"parameter": string(paramJSON),
		},
	}
}
