package client

import (
	"context"
	"log"
	"testing"

	"github.com/goodguy-project/goodguy-core/client/crawl"
	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/initialize"
	"github.com/goodguy-project/goodguy-core/util/jsonx"
)

func TestGetRecentContest(t *testing.T) {
	initialize.MustInit()
	t.Run("all", func(t *testing.T) {
		resp, err := crawl.Client.MGetRecentContest(context.Background(), &idl.MGetRecentContestRequest{})
		if err != nil {
			t.Errorf("MGetRecentContest.all failed, err: %v", err)
		}
		log.Printf("resp: %s", jsonx.Json(resp))
	})
}
