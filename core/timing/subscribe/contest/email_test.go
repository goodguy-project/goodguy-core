package contest

import (
	"context"
	"io/ioutil"
	"sort"
	"testing"
	"time"

	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
)

func TestGetEmailBody(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		data := doCrawl(context.Background())
		var contests []*model.Contest
		for o, c := range data {
			for _, cc := range c {
				if time.Now().Unix() < cc.GetTimestamp() {
					contests = append(contests, &model.Contest{ContestMessage: cc, OnlineJudge: o})
				}
			}
		}
		sort.Slice(contests, func(i, j int) bool {
			return contests[i].ContestMessage.GetTimestamp() < contests[j].ContestMessage.GetTimestamp()
		})
		var cm []*idl.RecentContest_ContestMessage
		for _, c := range contests {
			cm = append(cm, c.ContestMessage)
			if len(cm) > 1 {
				break
			}
		}
		body := getEmailBody("test-platform", cm)
		err := ioutil.WriteFile("email-basic.html", []byte(body), 0644)
		if err != nil {
			t.Error(err)
		}
	})
}
