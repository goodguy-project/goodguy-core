package statistics

import (
	"context"
	"log"

	"github.com/goodguy-project/goodguy-core/client/crawl"
	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/model"
	"github.com/goodguy-project/goodguy-core/util"
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
	util.Debug("record: %s\n", util.Json(r))
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
				Platform: util.Codeforces,
				Handle:   member.CodeforcesId,
			})
		}
		if member.AtcoderId != "" {
			request.GetUserContestRecordRequest = append(request.GetUserContestRecordRequest, &idl.GetUserContestRecordRequest{
				Platform: util.Atcoder,
				Handle:   member.AtcoderId,
			})
		}
		if member.NowcoderId != "" {
			request.GetUserContestRecordRequest = append(request.GetUserContestRecordRequest, &idl.GetUserContestRecordRequest{
				Platform: util.Nowcoder,
				Handle:   member.NowcoderId,
			})
		}
		if member.LeetcodeId != "" {
			request.GetUserContestRecordRequest = append(request.GetUserContestRecordRequest, &idl.GetUserContestRecordRequest{
				Platform: util.Leetcode,
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
		handle(member.Sid, util.Codeforces, member.CodeforcesId, t)
		handle(member.Sid, util.Atcoder, member.AtcoderId, t)
		handle(member.Sid, util.Nowcoder, member.NowcoderId, t)
		handle(member.Sid, util.Leetcode, member.LeetcodeId, t)
	}
	return nil
}
