package handler

import (
	"context"
	"log"
	"testing"

	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/util"
)

func TestGetMember(t *testing.T) {
	response, err := GetMember(context.Background(), &idl.GetMemberRequest{
		PageNo:   1,
		PageSize: 50,
	})
	if err != nil {
		t.Errorf("GetMember err: %v", err)
	}
	log.Printf("size: %v\n", response.GetSize())
	for i, r := range response.GetMember() {
		log.Printf("member %d: %s", i+1, util.Json(r))
	}
}
