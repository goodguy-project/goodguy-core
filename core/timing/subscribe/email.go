package subscribe

import (
	"sort"

	"github.com/goodguy-project/goodguy-core/client/smtp"
	"github.com/goodguy-project/goodguy-core/core/oj"
	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
)

func doEmailSend(m *model.Member, sd []*idl.RecentContest_ContestMessage) error {
	req := &smtp.SendEmailRequest{
		To:      []string{m.Email},
		Subject: "",
		Body:    "",
	}
	if m.SelfingMode {
		req.From = m.Email
		req.SmtpHost = m.SMTP.Host
		req.SmtpPort = m.SMTP.Port
		req.Pwd = m.SMTP.Pwd
		return smtp.SendEmail(req)
	}
	// TODO: Common Send
	return nil
}

func doEmailSubscribe(subscriber []*model.Member, contests map[*oj.OnlineJudge][]*idl.RecentContest_ContestMessage) {
	for _, m := range subscriber {
		sd := getSubscribeData(contests, m.EmailBit)
		sort.Slice(sd, func(i, j int) bool {
			return sd[i].Timestamp < sd[j].Timestamp
		})
		_ = doEmailSend(m, sd)
		// TODO: Log Database
	}
}
