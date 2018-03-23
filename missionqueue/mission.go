package missionqueue

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"monika/database"
	"monika/http"
	"monika/model"
	"monika/util"
)

type MissionQueue struct {
	Url    chan string
	Cookie *http.Cookie
	DB     *database.Database
}

func NewMission(cookie string, url string) (*MissionQueue, error) {
	c := http.NewCookie(cookie)
	db, err := database.NewDatabase(url)
	if err != nil {
		return nil, err
	}

	return &MissionQueue{
		Url:    make(chan string, 10),
		Cookie: c,
		DB:     db,
	}, nil
}

func (m *MissionQueue) setMission(url string) {
	m.Url <- url
}

func (s *MissionQueue) DoMission() {

	for {
		select {
		case url := <-s.Url:
			document, err := goquery.NewDocument(url, s.Cookie.GetCookie())
			if err != nil {
				log.Infof("missionqueue/DoMission:error %v", err)
			}
			document.Find(".c").Each(func(i int, sel *goquery.Selection) {
				// 只转发而没有评论的微博，包含class="cmt"和class="ctt"
				cmt := sel.Find("div").Find(".cmt").Text()
				if cmt == "" {
					//原创
					ctt := sel.Find("div").Find(".ctt").Text()
					//时间和客户端
					ct := sel.Find("div").Find(".ct").Text()
					if ctt != "" {
						log.Infof("原创微博: %s %s\n", ctt, ct)
						cts := util.Split(ct, "来自")
						if len(cts) == 2 {
							time, _ := util.ConvertTime(cts[0])
							weibo := model.Weibo{Content: ctt, Time: time, Client: cts[1]}
							err = s.DB.InsertWeibo(weibo)
							if err != nil {
								log.Errorf("missionqueue/domission:error %v", err)
							}
						} else {
							weibo := model.Weibo{Content: ctt, Client: ct}
							err = s.DB.InsertWeibo(weibo)
							if err != nil {
								log.Errorf("missionqueue/domission:error %v", err)
							}
						}

					}
				} else {
					//转发出处
					from := sel.Find("div").Find(".cmt").First().Text()
					//转发理由
					reason := sel.Find("div").Last().Text()
					//转发的微博正文
					ctt := sel.Find("div").Find(".ctt").Text()
					//转发微博图片地址
					image, err := sel.Find("div").First().Last().Html()
					//时间和客户端
					ct := sel.Find("div").Find(".ct").Text()
					log.Infof("转发微博: %s %s %s %s %s \n", reason, ctt, from, image, ct)
					cts := util.Split(ct, "来自")
					if len(cts) == 2 {
						time, _ := util.ConvertTime(cts[0])
						weibo := model.WeiboForward{Comment: reason, Content: ctt, Time: time, Client: cts[1]}
						err = s.DB.InsertWeiboForward(weibo)
						if err != nil {
							log.Errorf("missionqueue/domission:error %v", err)
						}
					} else {
						weibo := model.WeiboForward{Comment: reason, Content: ctt, Client: ct}
						err = s.DB.InsertWeiboForward(weibo)
						if err != nil {
							log.Errorf("missionqueue/domission:error %v", err)
						}
					}
				}
			})
		}
	}
}
