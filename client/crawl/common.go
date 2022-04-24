package crawl

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/util/conf"
)

var (
	Client idl.CrawlServiceClient
)

func init() {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.Viper().GetString("client.crawler.host"), conf.Viper().GetInt("client.crawler.port")),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	Client = idl.NewCrawlServiceClient(conn)
}
