package qzone

import (
	"QzoneRecorder/pkg/utils"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getPTQRToken(qrsig string) string {
	e := 0
	for i := 1; i < len(qrsig)+1; i++ {
		e += (e << 5) + int(qrsig[i-1])
	}
	return strconv.Itoa(2147483647 & e)
}

func getResultTextArray(s string) []string {
	// 按照逗号分割字符串
	arr := strings.Split(s, ",")
	// 创建一个空的切片
	result := []string{}
	// 遍历分割后的数组
	for _, v := range arr {
		// 去掉每个元素两边的括号和引号
		v = strings.Trim(v, "()'")
		// 将元素添加到切片中
		result = append(result, v)
	}
	// 打印结果
	return result
}

func (m *QzoneManager) LoginViaQRCode(qr_got_callback func(path string)) (string, error) {
	// 下载二维码到本地
	resp, err := utils.DownloadFile(qrcode_url, "qrcode.png")
	if err != nil {
		return "", err
	}
	set_cookies := resp.Header.Get("Set-Cookie")

	qrsig := ""

	set_cookies_split := strings.Split(set_cookies, ";")
	for _, cookie := range set_cookies_split {
		if strings.HasPrefix(cookie, "qrsig") {
			qrsig = strings.Split(cookie, "=")[1]
			break
		}
	}
	if qrsig == "" {
		return "", errors.New("qrsig获取失败")
	}

	// 调用回调函数
	qr_got_callback("qrcode.png")

	ptqrtoken := getPTQRToken(qrsig)

	// 轮询
	max_retry := 10

	retry := 0

	cookies := ""
	for {
		time.Sleep(2 * time.Second)

		req, _ := http.NewRequest("GET", fmt.Sprintf(login_check_url, ptqrtoken), nil)
		req.Header.Set("Cookie", "qrsig="+qrsig)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			retry++
			if retry > max_retry {
				return "", err
			}
			continue
		}
		retry = 0
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			continue
		}

		res_text := string(bytes)

		if strings.Contains(res_text, "二维码已失效") {
			return "", errors.New("二维码已失效")
		}

		if strings.Contains(res_text, "登录成功") {
			// 登录成功
			ret_url := getResultTextArray(res_text)[2]

			// 从ret_url中提取ptsigx参数内容
			ptsigx := ""
			ptsigx_regexp := regexp.MustCompile(`ptsigx=(.*?)&`)
			ptsigx_match := ptsigx_regexp.FindStringSubmatch(ret_url)
			if len(ptsigx_match) > 1 {
				ptsigx = ptsigx_match[1]
			}

			// 从ret_url中提取uin参数内容
			uin := ""
			uin_regexp := regexp.MustCompile(`uin=(.*?)&`)
			uin_match := uin_regexp.FindStringSubmatch(ret_url)
			if len(uin_match) > 1 {
				uin = uin_match[1]
			}
			m.Uin = uin

			// set-cookie
			set_cookies_slice := res.Header.Values("Set-Cookie")

			set_cookies := ""
			for _, cookie := range set_cookies_slice {
				set_cookies += cookie
			}

			// 获取skey和p_skey
			check_sig_req, _ := http.NewRequest("GET", fmt.Sprintf(check_sig_url, uin, ptsigx), nil)

			check_sig_req.Header.Set("Cookie", set_cookies)

			check_sig_res, err := http.DefaultClient.Do(check_sig_req)
			if err != nil {
				return "", err
			}

			final_cookies_slice := check_sig_res.Header.Values("Set-Cookie")

			final_cookies := ""
			for _, cookie := range final_cookies_slice {
				final_cookies += cookie
			}

			final_cookies_dict := map[string]string{}

			final_cookies_split := strings.Split(final_cookies, ";, ")
			for _, set_cookie := range final_cookies_split {
				for _, cookie := range strings.Split(set_cookie, ";") {
					spt := strings.Split(cookie, "=")
					if len(spt) == 2 && final_cookies_dict[spt[0]] == "" {
						final_cookies_dict[spt[0]] = spt[1]
					}
				}
			}

			// 写进cookies
			cookies = ""
			for k, v := range final_cookies_dict {
				cookies += k + "=" + v + ";"
			}
			m.Cookie_map = final_cookies_dict

			break
		}
	}

	m.Cookie_str = cookies

	return cookies, nil
}

func (m *QzoneManager) LoginViaCookies(cookies string) error {
	m.Cookie_str = cookies
	// 将cookies转换成map
	cookies_map := map[string]string{}

	cookies_split := strings.Split(cookies, ";")

	for _, cookie := range cookies_split {
		spt := strings.Split(cookie, "=")
		if len(spt) == 2 && cookies_map[spt[0]] == "" {
			cookies_map[spt[0]] = spt[1]
		}
	}

	m.Cookie_map = cookies_map

	// 从cookies提取uin的值并删除最前面的o
	uin := ""

	for k, v := range cookies_map {
		if k == "uin" {
			uin = v
			break
		}
	}

	if uin == "" {
		return errors.New("cookies中没有uin")
	}

	m.Uin = uin[1:]
	if m.CheckCookiesUsability(3) {
		return nil
	}
	return errors.New("不可用的cookies")
}

func (m *QzoneManager) CheckCookiesUsability(retry int) bool {
	for i := 0; i < retry; i++ {
		_, _, err := m.GetVisitorAmount()
		if err != nil {
			continue
		}
		return true
	}
	return false
}
