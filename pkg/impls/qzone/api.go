package qzone

import (
	entities "QzoneRecorder/pkg/models/qzone/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"

	goquery "github.com/PuerkitoBio/goquery"
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

	if len(json_data.Data.ModVisitCount) == 0 {
		return 0, 0, fmt.Errorf("len(json_data.Data.ModVisitCount) == 0")
	}

	return json_data.Data.ModVisitCount[0].TodayCount, json_data.Data.ModVisitCount[0].TotalCount, nil
}

func ParseEmotionFromHTML(html string) (entities.Emotion, error) {

	emotion := entities.Emotion{
		Comments: []entities.Comment{},
		Medias:   []entities.Media{},
	}

	unescaped := strings.ReplaceAll(html, "\\x3C", "<")
	unescaped = strings.ReplaceAll(unescaped, "\\/", "/")
	unescaped = strings.ReplaceAll(unescaped, "\\x22", "\"")
	unescaped = strings.ReplaceAll(unescaped, "&amp;", "&")

	// 把unescaped写入文件unescaped.html

	err := ioutil.WriteFile("unescaped.html", []byte(unescaped), 0644)

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(unescaped))
	if err != nil {
		return entities.Emotion{}, err
	}

	userCard := entities.UserCard{}
	// nickname
	// .f-nick的div下的a标签的text是nickname
	dom.Find(".f-nick").Each(func(i int, selection *goquery.Selection) {
		aTag := selection.Find(".f-name")
		userCard.Nickname = aTag.Text()
		userLink := aTag.AttrOr("link", "")
		// link="nameCard_12345"
		uin, err := strconv.ParseInt(strings.ReplaceAll(userLink, "nameCard_", ""), 10, 64)
		if err != nil {
			return
		}
		userCard.QQ = strconv.FormatInt(uin, 10)
	})

	emotion.UserCard = userCard

	// text
	dom.Find(".f-info").Each(func(i int, selection *goquery.Selection) {
		emotion.Text += selection.Text() + "\n"
	})
	if len(emotion.Text) > 0 {
		emotion.Text = emotion.Text[:len(emotion.Text)-1]
	}

	// tid
	dom.Find(".none").Each(func(i int, selection *goquery.Selection) {
		tid := selection.AttrOr("data-tid", "")
		if strings.Trim(tid, " ") != "" {
			emotion.Eid = tid
		}
	})

	if emotion.Eid == "" {
		return entities.Emotion{}, fmt.Errorf("no tid found, perhaps this is a advertisement")
	}

	// images
	dom.Find(".img-item").Each(func(i int, selection *goquery.Selection) {
		imgTag := selection.Find("img")

		// 图片的src
		src := imgTag.AttrOr("src", "")

		img := entities.Media{
			Type: "image",
			Url:  src,
		}

		emotion.Medias = append(emotion.Medias, img)
	})

	// traffic
	traffic := entities.Traffic{
		Likers: []entities.UserCard{},
	}
	// visit_count
	dom.Find(".qz_feed_plugin").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()

		num_str := strings.ReplaceAll(text, "浏览", "")
		num_str = strings.ReplaceAll(num_str, "次", "")

		tr, err := strconv.ParseInt(num_str, 10, 64)

		if err != nil {
			return
		}
		traffic.VisitAmount = int(tr)
	})

	// forward
	dom.Find(".wupfeed").Each(func(i int, s *goquery.Selection) {
		s.Find("i").Each(func(i int, s *goquery.Selection) {
			forwardCountStr := s.AttrOr("data-retweetcount", "0")

			forwardCount, err := strconv.ParseInt(forwardCountStr, 10, 64)
			if err != nil {
				return
			}

			traffic.ForwardAmount = int(forwardCount)
		})
	})

	// like amount
	dom.Find(".f-like-cnt").Each(func(i int, s *goquery.Selection) {
		likeCountStr := s.Text()

		likeCount, err := strconv.ParseInt(likeCountStr, 10, 64)
		if err != nil {
			return
		}

		traffic.LikeAmount = int(likeCount)
	})

	emotion.Traffic = traffic

	// comments

	// detail_page_url
	emotion.DetailPageUrl = fmt.Sprintf("https://user.qzone.qq.com/%s/mood/%s", emotion.UserCard.QQ, emotion.Eid)

	return emotion, nil
}

func (m *QzoneManager) FetchFeedsList(pageNum int) ([]entities.Emotion, error) {
	result := []entities.Emotion{}

	feeds_jsonp_resp, err := m.GetWithContext(
		fmt.Sprintf(feeds3_html_more, m.Uin, pageNum, strconv.FormatInt(time.Now().Unix()*1000, 10), m.GenerateGTK()),
		map[string]string{},
	)
	if err != nil {
		return nil, err
	}

	feeds_jsonp_bytes, err := ioutil.ReadAll(feeds_jsonp_resp.Body)

	if err != nil {
		return nil, err
	}

	// 字符串
	feeds_jsonp := string(feeds_jsonp_bytes)

	// 提取html:'xxx',中的xxx
	re := regexp.MustCompile(`html:'(.*?)',`)

	htmls := re.FindAllStringSubmatch(feeds_jsonp, -1)

	if len(htmls) == 0 {
		return result, nil
	} else {
		// 替换掉html中的转义字符

		for _, html := range htmls {
			emo, err := ParseEmotionFromHTML(html[1])
			if err != nil {
				continue
			}
			result = append(result, emo)
		}
	}
	return result, nil
}
