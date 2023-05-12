package qzone

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 获取访客数量
// 返回: 今天的访客数量, 总访客数量, 错误
func (m *QzoneManager) GetVisitorAmount() (int, int, error) {
	resp, err := m.GetWithContext(
		fmt.Sprintf(cgi_get_visitor_simple, m.Uin, m.GenerateGTK()),
		map[string]string{},
	)
	if err != nil {
		return 0, 0, err
	}

	json_bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}
	json_str := string(json_bytes)

	//json_text = res.text.replace("_Callback(", '')[:-3]
	json_str = json_str[10 : len(json_str)-3]

	var json_data struct {
		Code int `json:"code"`
		Data struct {
			ModVisitCount []struct {
				TodayCount int `json:"todaycount"`
				TotalCount int `json:"totalcount"`
			} `json:"modvisitcount"`
		} `json:"data"`
	}

	err = json.Unmarshal([]byte(json_str), &json_data)
	if err != nil {
		return 0, 0, err
	}

	if json_data.Code != 0 {
		return 0, 0, fmt.Errorf("code != 0 (%d)", json_data.Code)
	}

	return json_data.Data.ModVisitCount[0].TodayCount, json_data.Data.ModVisitCount[0].TotalCount, nil
}
