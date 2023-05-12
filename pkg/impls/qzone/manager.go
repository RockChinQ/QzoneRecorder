package qzone

import (
	"net/http"
	"strconv"
)

const (
	// qrcode url
	qrcode_url             = "https://ssl.ptlogin2.qq.com/ptqrshow?appid=549000912&e=2&l=M&s=3&d=72&v=4&t=0.31232733520361844&daid=5&pt_3rd_aid=0"
	login_check_url        = "https://xui.ptlogin2.qq.com/ssl/ptqrlogin?u1=https://qzs.qq.com/qzone/v5/loginsucc.html?para=izone&ptqrtoken=%s&ptredirect=0&h=1&t=1&g=1&from_ui=1&ptlang=2052&action=0-0-1656992258324&js_ver=22070111&js_type=1&login_sig=&pt_uistyle=40&aid=549000912&daid=5&has_onekey=1&&o1vId=1e61428d61cb5015701ad73d5fb59f73"
	check_sig_url          = "https://ptlogin2.qzone.qq.com/check_sig?pttype=1&uin=%s&service=ptqrlogin&nodirect=1&ptsigx=%s&s_url=https://qzs.qq.com/qzone/v5/loginsucc.html?para=izone&f_url=&ptlang=2052&ptredirect=100&aid=549000912&daid=5&j_later=0&low_login_hour=0&regmaster=0&pt_login_type=3&pt_aid=0&pt_aaid=16&pt_light=0&pt_3rd_aid=0"
	cgi_get_visitor_more   = "https://h5.qzone.qq.com/proxy/domain/g.qzone.qq.com/cgi-bin/friendshow/cgi_get_visitor_more?uin=%s&mask=7&g_tk=%s&page=1&fupdate=1&clear=1"
	cgi_get_visitor_simple = "https://user.qzone.qq.com/proxy/domain/g.qzone.qq.com/cgi-bin/friendshow/cgi_get_visitor_simple?uin=%s&mask=1&g_tk=%s"
)

type QzoneManager struct {
	Cookie_str string // 登录所需的cookies
	Cookie_map map[string]string
	Headers    map[string]string
	Uin        string // 登录的QQ号
}

func NewQzoneManager() *QzoneManager {
	return &QzoneManager{}
}

func (m *QzoneManager) GenerateGTK() string {
	skey := m.Cookie_map["skey"]
	hash_val := 5381
	for _, v := range skey {
		hash_val += (hash_val << 5) + int(v)
	}

	gtk := hash_val & 2147483647

	return strconv.Itoa(gtk)
}

func (m *QzoneManager) GetWithContext(url string, data map[string]string) (*http.Response, error) {
	// 使用m的cookie_str和headers进行请求

	if len(data) > 0 {
		data_str := ""
		for k, v := range data {
			data_str += k + "=" + v + "&"
		}

		url += "?" + data_str
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cookie", m.Cookie_str)

	for k, v := range m.Headers {
		req.Header.Set(k, v)
	}

	return http.DefaultClient.Do(req)
}

func (m *QzoneManager) PostWithContext(url string, data map[string]string) (*http.Response, error) {
	// 使用m的cookie_str和headers进行请求

	if len(data) > 0 {

		data_str := ""
		for k, v := range data {
			data_str += k + "=" + v + "&"
		}

		url = url + "?" + data_str
	}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cookie", m.Cookie_str)

	for k, v := range m.Headers {
		req.Header.Set(k, v)
	}

	return http.DefaultClient.Do(req)
}
