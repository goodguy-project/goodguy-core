package token

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc/metadata"

	"github.com/goodguy-project/goodguy-core/util"
)

var (
	c = cache.New(4*time.Hour, 6*time.Hour)
)

func GetToken(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	t := md.Get("token")
	if len(t) > 0 {
		return t[0]
	}
	return ""
}

func Auth(ctx context.Context) (string, bool) {
	return GetMemberSid(GetToken(ctx))
}

func GetMemberSid(token string) (string, bool) {
	s, ok := c.Get(token)
	if !ok {
		return "", false
	}
	return s.(string), true
}

func SetMemberSid(sid string) string {
	token := util.Hashing(fmt.Sprintf("%v+_+%v+_+%v", time.Now().Unix(), sid, rand.Float64()))
	c.Set(token, sid, 0)
	return token
}
