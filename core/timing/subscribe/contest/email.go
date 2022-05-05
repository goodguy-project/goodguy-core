package contest

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
	"golang.org/x/crypto/sha3"

	"github.com/goodguy-project/goodguy-core/client/smtp"
	"github.com/goodguy-project/goodguy-core/core/oj"
	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
	"github.com/goodguy-project/goodguy-core/util/jsonx"
)

func getEmailBody(platform string, contests []*model.SingleContestMsg) string {
	resp := `<!doctype html><html xmlns="http://www.w3.org/1999/xhtml"xmlns:v="urn:schemas-microsoft-com:vml"xmlns:o="urn:schemas-microsoft-com:office:office"><head><title></title><!--[if!mso]><!--><meta http-equiv="X-UA-Compatible"content="IE=edge"><!--<![endif]--><meta http-equiv="Content-Type"content="text/html; charset=UTF-8"><meta name="viewport"content="width=device-width, initial-scale=1"><style type="text/css">#outlook a{padding:0}body{margin:0;padding:0;-webkit-text-size-adjust:100%;-ms-text-size-adjust:100%}table,td{border-collapse:collapse;mso-table-lspace:0pt;mso-table-rspace:0pt}img{border:0;height:auto;line-height:100%;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic}p{display:block;margin:13px 0}</style><!--[if mso]><noscript><xml><o:OfficeDocumentSettings><o:AllowPNG/><o:PixelsPerInch>96</o:PixelsPerInch></o:OfficeDocumentSettings></xml></noscript><![endif]--><!--[if lte mso 11]><style type="text/css">.mj-outlook-group-fix{width:100%!important}</style><![endif]--><!--[if!mso]><!--><link href="https://fonts.googleapis.com/css?family=Ubuntu:300,400,500,700"rel="stylesheet"type="text/css"><style type="text/css">@import url(https:</style><!--<![endif]--><style type="text/css">@media only screen and(min-width:480px){.mj-column-per-100{width:100%!important;max-width:100%}}</style><style media="screen and (min-width:480px)">.moz-text-html.mj-column-per-100{width:100%!important;max-width:100%}</style><style type="text/css"></style><style type="text/css"></style></head><body style="word-spacing:normal;"><div style=""><!--[if mso|IE]><table align="center"border="0"cellpadding="0"cellspacing="0"class=""role="presentation"style="width:600px;"width="600"><tr><td style="line-height:0px;font-size:0px;mso-line-height-rule:exactly;"><![endif]--><div style="margin:0px auto;max-width:600px;"><table align="center"border="0"cellpadding="0"cellspacing="0"role="presentation"style="width:100%;"><tbody><tr><td style="direction:ltr;font-size:0px;padding:0;text-align:center;"><!--[if mso|IE]><table role="presentation"border="0"cellpadding="0"cellspacing="0"><tr><td class=""style="vertical-align:top;width:600px;"><![endif]--><div class="mj-column-per-100 mj-outlook-group-fix"style="font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;"><table border="0"cellpadding="0"cellspacing="0"role="presentation"style="vertical-align:top;"width="100%"><tbody><tr><td align="center"style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:13px;line-height:1;text-align:center;color:#000000;">大家好！此邮件是比赛提醒邮件，自动发送，请勿回复。</div></td></tr></tbody></table></div><!--[if mso|IE]></td></tr></table><![endif]--></td></tr></tbody></table></div><!--[if mso|IE]></td></tr></table><table align="center"border="0"cellpadding="0"cellspacing="0"class=""role="presentation"style="width:600px;"width="600"><tr><td style="line-height:0px;font-size:0px;mso-line-height-rule:exactly;"><![endif]--><div style="margin:0px auto;max-width:600px;"><table align="center"border="0"cellpadding="0"cellspacing="0"role="presentation"style="width:100%;"><tbody><tr><td style="direction:ltr;font-size:0px;padding:0;text-align:center;"><!--[if mso|IE]><table role="presentation"border="0"cellpadding="0"cellspacing="0"><tr><td class=""style="vertical-align:top;width:600px;"><![endif]--><div class="mj-column-per-100 mj-outlook-group-fix"style="font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;"><table border="0"cellpadding="0"cellspacing="0"role="presentation"style="vertical-align:top;"width="100%"><tbody><tr><td align="center"style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:13px;line-height:1;text-align:center;color:#000000;">你订阅的以下比赛将在一个小时后开始：</div></td></tr></tbody></table></div><!--[if mso|IE]></td></tr></table><![endif]--></td></tr></tbody></table></div><!--[if mso|IE]></td></tr></table><table align="center"border="0"cellpadding="0"cellspacing="0"class=""role="presentation"style="width:600px;"width="600"><tr><td style="line-height:0px;font-size:0px;mso-line-height-rule:exactly;"><![endif]-->****CONTEST****<div style="margin:0px auto;max-width:600px;"><table align="center"border="0"cellpadding="0"cellspacing="0"role="presentation"style="width:100%;"><tbody><tr><td style="direction:ltr;font-size:0px;padding:0;text-align:center;"><!--[if mso|IE]><table role="presentation"border="0"cellpadding="0"cellspacing="0"><tr><td class=""style="vertical-align:top;width:600px;"><![endif]--><div class="mj-column-per-100 mj-outlook-group-fix"style="font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;"><table border="0"cellpadding="0"cellspacing="0"role="presentation"style="vertical-align:top;"width="100%"><tbody><tr><td align="center"style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:13px;line-height:1;text-align:center;color:#000000;">****EMAILTAIL****</div></td></tr></tbody></table></div><!--[if mso|IE]></td></tr></table><![endif]--></td></tr></tbody></table></div><!--[if mso|IE]></td></tr></table><![endif]--></div></body></html>`
	inner := ""
	for _, contest := range contests {
		c := `<div style="margin:0px auto;max-width:600px;"><table align="center"border="0"cellpadding="0"cellspacing="0"role="presentation"style="width:100%;"><tbody><tr><td style="direction:ltr;font-size:0px;padding:20px 0;text-align:center;"><!--[if mso|IE]><table role="presentation"border="0"cellpadding="0"cellspacing="0"><tr><td class=""style="vertical-align:top;width:600px;"><![endif]--><div class="mj-column-per-100 mj-outlook-group-fix"style="font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;"><table border="0"cellpadding="0"cellspacing="0"role="presentation"style="border:2px solid blue;vertical-align:top;"width="100%"><tbody><tr><td align="center"style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:15px;line-height:1;text-align:center;color:#000000;">****PLATFORM****</div></td></tr><tr><td align="center"style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:24px;line-height:1;text-align:center;color:#000000;">****NAME****</div></td></tr><tr><td align="center"style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:18px;line-height:1;text-align:center;color:#000000;">****START****</div></td></tr><tr><td align="center"style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:18px;line-height:1;text-align:center;color:#000000;">****DURATION****</div></td></tr><tr><td align="center"vertical-align="middle"style="font-size:0px;padding:10px 25px;word-break:break-word;"><table border="0"cellpadding="0"cellspacing="0"role="presentation"style="border-collapse:separate;line-height:100%;"><tbody><tr><td align="center"bgcolor="#414141"role="presentation"style="border:none;border-radius:3px;cursor:auto;mso-padding-alt:10px 25px;background:#414141;"valign="middle"><a href="****URL****"style="display:inline-block;background:#414141;color:#ffffff;font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:13px;font-weight:normal;line-height:120%;margin:0;text-decoration:none;text-transform:none;padding:10px 25px;mso-padding-alt:0px;border-radius:3px;"target="_blank">进入比赛</a></td></tr></tbody></table></td></tr></tbody></table></div><!--[if mso|IE]></td></tr></table><![endif]--></td></tr></tbody></table></div><!--[if mso|IE]></td></tr></table><table align="center"border="0"cellpadding="0"cellspacing="0"class=""role="presentation"style="width:600px;"width="600"><tr><td style="line-height:0px;font-size:0px;mso-line-height-rule:exactly;"><![endif]-->`
		c = strings.ReplaceAll(c, "****PLATFORM****", platform)
		c = strings.ReplaceAll(c, "****NAME****", contest.Name)
		startTime := time.Unix(contest.Timestamp, 0)
		c = strings.ReplaceAll(c, "****START****", startTime.Format("2006年01月02日 15:04:05"))
		c = strings.ReplaceAll(c, "****DURATION****", fmt.Sprintf("时长：%v分钟", contest.Duration/60))
		c = strings.ReplaceAll(c, "****URL****", contest.Url)
		inner += c
	}
	resp = strings.ReplaceAll(resp, "****CONTEST****", inner)
	resp = strings.ReplaceAll(resp, "****EMAILTAIL****", "welcome to star my project: https://github.com/goodguy-project/goodguy-core")
	return resp
}

func emailJob(hashTag string) {
	subscribeLog := &model.SubscribeLog{}
	db := model.GetDB().Model(&model.SubscribeLog{})
	db.Where("hash_tag = ?", hashTag).First(subscribeLog)
	if subscribeLog.Sid == "" || subscribeLog.ContestMsg == nil {
		log.Printf("something is error, hashTag: %s, response: %s", hashTag, jsonx.Json(subscribeLog))
		return
	}
	err := doEmailSubscribe(subscribeLog.ContestMsg.Member, subscribeLog.ContestMsg.Platform, subscribeLog.ContestMsg.Contest)
	defer func() {
		db.Save(subscribeLog)
	}()
	if err != nil {
		subscribeLog.Status = model.SubscribeLogStatus_Failed
		log.Printf("doEmailSubscribe error: %v", err)
		return
	}
	subscribeLog.Status = model.SubscribeLogStatus_OK
}

func doEmailSubscribe(m *model.Member, platform string, contest []*model.SingleContestMsg) error {
	if m == nil || len(contest) == 0 {
		log.Printf("internal error, member: %s, contests: %s", jsonx.Json(m), jsonx.Json(contest))
	}
	if !m.SelfingMode {
		return nil
	}
	body := getEmailBody(platform, contest)
	req := &smtp.SendEmailRequest{
		To:      []string{m.Email},
		Subject: "最近比赛提醒",
		Body:    body,
	}
	req.From = m.Email
	req.SmtpHost = m.SMTP.Host
	req.SmtpPort = m.SMTP.Port
	req.Pwd = m.SMTP.Pwd
	return smtp.SendEmail(req)
}

func getEmailHashTag(sid string, contest *model.Contest) string {
	s := fmt.Sprintf("email:%s%s%d", sid, contest.OnlineJudge.Name(), contest.ContestMessage.GetTimestamp())
	return fmt.Sprintf("%x", sha3.Sum384([]byte(s)))
}

func newEmailSubscribeJob(m *model.Member, contest *model.Contest) {
	now := time.Now()
	t := time.Unix(contest.ContestMessage.GetTimestamp(), 0)
	t = t.Add(-time.Hour)
	if now.Before(t) && t.Before(now.Add(36*time.Hour)) {
		hashTag := getEmailHashTag(m.Sid, contest)
		subscribeLog := &model.SubscribeLog{}
		db := model.GetDB().Model(&model.SubscribeLog{}).Where("hash_tag = ?", hashTag).First(subscribeLog)
		if subscribeLog.Sid != "" {
			subscribeLog.ContestMsg.Contest = append(subscribeLog.ContestMsg.Contest, &model.SingleContestMsg{
				Name:      contest.ContestMessage.GetName(),
				Url:       contest.ContestMessage.GetUrl(),
				Timestamp: contest.ContestMessage.GetTimestamp(),
				Duration:  contest.ContestMessage.GetDuration(),
			})
			db.Save(subscribeLog)
			return
		}
		model.GetDB().Create(&model.SubscribeLog{
			Sid:           m.Sid,
			Type:          "email",
			What:          "contest",
			Bit:           m.SubscribeStatus.Email,
			SubscribeTime: time.Now().Unix(),
			Status:        model.SubscribeLogStatus_Undone,
			ContestMsg: &model.ContestMsg{
				Contest: []*model.SingleContestMsg{
					{
						Name:      contest.ContestMessage.GetName(),
						Url:       contest.ContestMessage.GetUrl(),
						Timestamp: contest.ContestMessage.GetTimestamp(),
						Duration:  contest.ContestMessage.GetDuration(),
					},
				},
				Member:   m,
				Platform: contest.OnlineJudge.Name(),
			},
		})
		shc, _ := time.LoadLocation("Asia/Shanghai")
		c := cron.New(cron.WithSeconds(), cron.WithLocation(shc))
		_, err := c.AddFunc(t.Format("05 04 15 02 01 2006"), func() {
			emailJob(hashTag)
		})
		if err != nil {
			log.Printf("cron AddFunc error: %v", err)
			return
		}
		go c.Run()
	}
}

func doEmail(subscriber []*model.Member, contests map[*oj.OnlineJudge][]*idl.RecentContest_ContestMessage) {
	for _, m := range subscriber {
		if !m.SelfingMode {
			continue
		}
		sc := getSubscribeContest(contests, m.SubscribeStatus.Email)
		for _, c := range sc {
			newEmailSubscribeJob(m, c)
		}
	}
}
