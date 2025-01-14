package goimouse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"sync/atomic"
)

type HttpApi struct {
	apiURL       string
	dataCreation *DataCreation
	msgID        int64
}

func NewHttpApi(host string) *HttpApi {
	return &HttpApi{
		apiURL:       fmt.Sprintf("http://%s:9912/api", host),
		dataCreation: &DataCreation{},
	}
}

func (api *HttpApi) getMsgID(sync bool) int64 {
	if !sync {
		return 0
	}
	return atomic.AddInt64(&api.msgID, 1)
}

func (api *HttpApi) post(data map[string]interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(api.apiURL, "application/json; charset=utf-8", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	utf8Body, err := convertToUTF8(body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(utf8Body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func convertToUTF8(body []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(body), simplifiedchinese.GBK.NewDecoder())
	return io.ReadAll(reader)
}

func (api *HttpApi) GetDeviceList(sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.GetDeviceList(int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) GetGroupList(sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.GetGroupList(int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) GetUsbList(sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.GetUsbList(int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) GetDeviceModelList(sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.GetDeviceModelList(int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) GetDeviceScreenshot(deviceID string, gzip, binary, isJpg, original bool, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.GetDeviceScreenshot(deviceID, gzip, binary, isJpg, original, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) LoopDeviceScreenshot(deviceID string, time int, stop, isJpg bool, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.LoopDeviceScreenshot(deviceID, time, stop, isJpg, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) Click(deviceID string, x, y int, button string, time int, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.Click(deviceID, x, y, button, time, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) Swipe(deviceID, direction, button string, length float64, sx, sy, ex, ey, afor int, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.Swipe(deviceID, direction, button, length, sx, sy, ex, ey, afor, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) MouseUp(deviceID, button string, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.MouseUp(deviceID, button, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) MouseDown(deviceID, button string, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.MouseDown(deviceID, button, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) MouseMove(deviceID string, x, y int, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.MouseMove(deviceID, x, y, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) MouseWheel(deviceID, direction string, length, number int, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.MouseWheel(deviceID, direction, length, number, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) MouseResetPos(deviceID string, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.MouseResetPos(deviceID, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) SendKey(deviceID, key string, fnKey *string, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.SendKey(deviceID, key, fnKey, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) Restart(sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.Restart(int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) SaveDevLocation(deviceID, describe string, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.SaveDevLocation(deviceID, describe, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) DelDevLocation(model, version, crc string, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.DelDevLocation(model, version, crc, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) GetAirplaySrvNum(sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.GetAirplaySrvNum(int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) GetUsbAutoAirplay(deviceID string, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.GetUsbAutoAirplay(deviceID, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) GetAirplayMode(deviceID string, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.GetAirplayMode(deviceID, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) AutoConnectScreenAll(sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.AutoConnectScreenAll(int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) SaveRestartPoint(deviceID string, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.SaveRestartPoint(deviceID, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) DisconAirplay(deviceID string, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.DisconAirplay(deviceID, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) AutoConnectScreen(deviceID string, force bool, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.AutoConnectScreen(deviceID, force, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) RestartUsb(deviceID string, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.RestartUsb(deviceID, int(api.getMsgID(sync)))
	return api.post(data)
}

func (api *HttpApi) Shortcut(deviceID string, id int, devList []string, parameter map[string]interface{}, outTime int, sync bool) (map[string]interface{}, error) {
	data := api.dataCreation.Shortcut(deviceID, id, devList, parameter, outTime, int(api.getMsgID(sync)))
	return api.post(data)
}
