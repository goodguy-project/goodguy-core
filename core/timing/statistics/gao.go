package statistics

import (
	"context"
	"log"

	"github.com/goodguy-project/goodguy-core/client/crawl"
	"github.com/goodguy-project/goodguy-core/core/oj"
	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
	"github.com/goodguy-project/goodguy-core/util"
	"github.com/goodguy-project/goodguy-core/util/conf"
	"github.com/goodguy-project/goodguy-core/util/jsonx"
)

func handle(sid, platform, handle string, t map[string]map[string]*idl.UserContestRecord) {
	if sid == "" || platform == "" || handle == "" {
		return
	}
	util.Debug("do statistics, platform: %s, handle: %s", platform, handle)
	var record *idl.UserContestRecord = nil
	if m1, o1 := t[platform]; o1 {
		if m2, o2 := m1[handle]; o2 {
			record = m2
		}
	}
	if record == nil || record.ProfileUrl == "" {
		log.Printf("crawl without result, request: {platform: '%s', handle: '%s'}", platform, handle)
		return
	}
	r := &model.MemberContestRecord{
		Sid:      sid,
		Platform: platform,
		Handle:   handle,
		Rating:   record.Rating,
		Length:   record.Length,
	}
	util.Debug("record: %s\n", jsonx.Json(r))
	if err := model.GetDB().Model(&model.MemberContestRecord{}).Where("Sid = ? and Platform = ? and handle = ?",
		sid, platform, handle).Assign(r).FirstOrCreate(&model.MemberContestRecord{}).Error; err != nil {
		log.Printf("update MemberContestRecord failed, err: %v", err)
	}
}

func doContestCrawl(ctx context.Context, members []*model.Member) error {
	var err error
	request := &idl.MGetUserContestRecordRequest{
		GetUserContestRecordRequest: make([]*idl.GetUserContestRecordRequest, 0),
	}
	for _, member := range members {
		if member.CodeforcesId != "" {
			request.GetUserContestRecordRequest = append(request.GetUserContestRecordRequest, &idl.GetUserContestRecordRequest{
				Platform: oj.Codeforces.Name(),
				Handle:   member.CodeforcesId,
			})
		}
		if member.AtcoderId != "" {
			request.GetUserContestRecordRequest = append(request.GetUserContestRecordRequest, &idl.GetUserContestRecordRequest{
				Platform: oj.Atcoder.Name(),
				Handle:   member.AtcoderId,
			})
		}
		if member.NowcoderId != "" {
			request.GetUserContestRecordRequest = append(request.GetUserContestRecordRequest, &idl.GetUserContestRecordRequest{
				Platform: oj.Nowcoder.Name(),
				Handle:   member.NowcoderId,
			})
		}
		if member.LeetcodeId != "" {
			request.GetUserContestRecordRequest = append(request.GetUserContestRecordRequest, &idl.GetUserContestRecordRequest{
				Platform: oj.Leetcode.Name(),
				Handle:   member.LeetcodeId,
			})
		}
	}
	response, err := crawl.Client.MGetUserContestRecord(ctx, request)
	if err != nil {
		log.Printf("doContestCrawl.MGetUserContestRecord failed, err: %v\n", err)
		return err
	}
	t := make(map[string]map[string]*idl.UserContestRecord)
	for _, r := range response.UserContestRecord {
		m, ok := t[r.Platform]
		if !ok {
			m = make(map[string]*idl.UserContestRecord)
			t[r.Platform] = m
		}
		m[r.Handle] = r
	}
	for _, member := range members {
		handle(member.Sid, oj.Codeforces.Name(), member.CodeforcesId, t)
		handle(member.Sid, oj.Atcoder.Name(), member.AtcoderId, t)
		handle(member.Sid, oj.Nowcoder.Name(), member.NowcoderId, t)
		handle(member.Sid, oj.Leetcode.Name(), member.LeetcodeId, t)
	}
	return nil
}

func gao() {
	var err error
	db := model.GetDB()
	count := int64(0)
	err = db.Model(&model.Member{}).Count(&count).Error
	if err != nil {
		log.Printf("database error, err: %v\n", err)
		return
	}
	buffer := conf.Viper().GetInt64("statistics.buffer")
	if buffer <= 0 {
		buffer = 100
	}
	times := (count + buffer - 1) / buffer
	for no := int64(0); no < times; no += 1 {
		var members []*model.Member
		err = db.Model(&model.Member{}).Offset(int(buffer * no)).Limit(int(buffer)).Find(&members).Error
		if err != nil {
			log.Printf("database error, err: %v\n", err)
		}
		_ = doContestCrawl(context.Background(), members)
	}
}
