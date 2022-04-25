package subscribe

import (
	"sort"

	"github.com/matcornic/hermes/v2"

	"github.com/goodguy-project/goodguy-core/client/smtp"
	"github.com/goodguy-project/goodguy-core/core/oj"
	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
)

// TODO: change demo
func getEmailBody() (string, error) {
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Hermes",
			Link: "https://example-hermes.com/",
			// Optional product logo
			Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
	}
	email := hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Confirm your account",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
	// Generate an HTML email with the provided contents (for modern clients)
	body, err := h.GenerateHTML(email)
	return body, err
}

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
