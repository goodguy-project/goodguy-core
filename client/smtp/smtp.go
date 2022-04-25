package smtp

import "gopkg.in/gomail.v2"

type SendEmailRequest struct {
	From        string   // 发信人
	To          []string // 收信人
	Cc          []string // 抄送人
	Bcc         []string // 密抄
	Subject     string   // 标题
	Body        string   // 内容
	SmtpHost    string   // SMTP服务器
	SmtpPort    int      // SMTP端口
	Pwd         string   // SMTP密码
	ContentType *string  // 默认为"text/html"
	SkipVerify  *bool    // 默认为false
}

func SendEmail(req *SendEmailRequest) error {
	m := gomail.NewMessage()
	m.SetHeader("From", req.From)
	m.SetHeader("To", req.To...)
	m.SetHeader("Cc", req.Cc...)
	m.SetHeader("Bcc", req.Bcc...)
	m.SetHeader("Subject", req.Subject)
	if req.ContentType != nil {
		m.SetBody(*req.ContentType, req.Body)
	} else {
		m.SetBody("text/html", req.Body)
	}
	d := gomail.NewDialer(req.SmtpHost, req.SmtpPort, req.From, req.Pwd)
	return d.DialAndSend(m)
}
